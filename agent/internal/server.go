package internal

import (
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"

	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

var (
	ipLocker        sync.RWMutex
	sameIpMap       map[uint32]uint32 //同一IP连接数限制
	sameIp          uint32            //同一IP连接数限制
	connectCount    int32             //当前连接数
	maxConnect      int32             //最大连接数
	coder           Coder             //编码方式
	rpcServicePool  ServicePool
	roomServicePool ServicePool
	tcpReadBuf 		int
	tcpWriteBuf 	int
)

func startServer(config *AppConfig) {
	rpcServicePool = NewServicePool(config.Consul.Addr, config.Consul.ServerPrefix, config.Consul.Services, false)

	if config.Consul.RoomPrefix != "" {
		roomServicePool = NewServicePool(config.Consul.Addr, config.Consul.RoomPrefix, config.Consul.Services, false)
	}

	maxConnect = config.MaxConnect
	sameIp = config.SameIp
	sameIpMap = make(map[uint32]uint32, maxConnect*2)

	if config.Pprof != "" {
		go http.ListenAndServe(config.Pprof, nil)
	}
	if config.Tcp.Listen != "" {
		go tcpServer(config)
	}
	if config.Kcp.Listen != "" {
		go kcpServer(config)
	}
	if config.Web.Listen != "" {
		go webServer(config)
	}
}

// 增加IP计数
func addIpLimit(ip uint32) bool {
	if sameIp > 0 {
		ipLocker.Lock()
		defer ipLocker.Unlock()
		if c := sameIpMap[ip]; c < sameIp {
			sameIpMap[ip] = c + 1
		} else {
			return false
		}
	}
	return true
}

// 减少IP计数
func decIpLimit(ip uint32) {
	if sameIp > 0 {
		ipLocker.Lock()
		defer ipLocker.Unlock()
		if c := sameIpMap[ip]; c <= 1 {
			delete(sameIpMap, ip)
		} else {
			sameIpMap[ip] = c - 1
		}
	}
}

func makeMsg(h *Handshake) []byte {
	src, _ := coder.Encode(h)
	return src
}

func tcpServer(config *AppConfig) {
	// resolve address & start listening
	l, err := net.Listen("tcp", config.Tcp.Listen)
	log.Info("listening on:", l.Addr())
	checkError(err)

	success := makeMsg(&Handshake{
		Ip: []string{"127.0.0.1"},
	})
	// TODO: IP实现服务自发现
	fail := makeMsg(&Handshake{
		Code: 1,
		Msg:  "服务器忙",
	})
	lis := l.(*net.TCPListener)
	defer lis.Close()
	// loop accepting
	for {
		select {
		case <-signal.Die():
			return
		default:
			conn, err := lis.AcceptTCP()
			if err != nil {
				log.Warning("accept failed:", err)
				continue
			}
			if atomic.LoadInt32(&connectCount) >= maxConnect {
				conn.Write(fail)
				conn.Close()
			} else {
				addr := conn.RemoteAddr().String()
				ipStr, _, _ := net.SplitHostPort(addr)
				ip := util.IpToUint32(ipStr)
				if addIpLimit(ip) {
					// set socket read buffer
					conn.SetReadBuffer(tcpReadBuf)
					// set socket write buffer
					conn.SetWriteBuffer(tcpWriteBuf)
					conn.Write(success)
					// start a goroutine for every incoming connection for reading
					go handleClient(ip, addr, conn)
				} else {
					conn.Close()
				}
			}
		}
	}
}

func kcpServer(config *AppConfig) {
	l, err := kcp.Listen(config.Kcp.Listen)
	checkError(err)
	log.Info("udp listening on:", l.Addr())

	lis := l.(*kcp.Listener)
	if err := lis.SetReadBuffer(config.Kcp.ReadBuf); err != nil {
		log.Fatalf("SetReadBuffer:", err)
	}
	if err := lis.SetWriteBuffer(config.Kcp.WriteBuf); err != nil {
		log.Fatalf("SetWriteBuffer:", err)
	}
	//if err := lis.SetDSCP(config.Udp.Dscp); err != nil {
	//	log.Fatalf("SetDSCP:", err)
	//}

	success := makeMsg(new(Handshake))
	fail := makeMsg(&Handshake{
		Code: 1,
		Msg:  "服务器忙",
	})

	defer lis.Close()
	// loop accepting
	for {
		select {
		case <-signal.Die():
			return
		default:
			conn, err := lis.AcceptKCP()
			if err != nil {
				log.Warning("accept failed:", err)
				continue
			}
			if atomic.LoadInt32(&connectCount) >= maxConnect {
				conn.Write(fail)
				conn.Close()
			} else {
				addr := conn.RemoteAddr().String()
				ipStr, _, _ := net.SplitHostPort(addr)
				ip := util.IpToUint32(ipStr)
				if addIpLimit(ip) {
					// set kcp parameters
					conn.SetWindowSize(config.Kcp.Sndwnd, config.Kcp.Rcvwnd)
					conn.SetNoDelay(config.Kcp.Nodelay, config.Kcp.Interval, config.Kcp.Resend, config.Kcp.Nc)
					conn.SetStreamMode(true)
					conn.SetMtu(config.Kcp.Mtu)
					conn.Write(success)
					// start a goroutine for every incoming connection for reading
					go handleClient(ip, addr, conn)
				} else {
					conn.Close()
				}
			}
		}
	}
}

// handleClient
func handleClient(ip uint32, addr string, conn net.Conn) {
	defer func() {
		atomic.AddInt32(&connectCount, -1)
		decIpLimit(ip)
		conn.Close()
	}()
	count := atomic.AddInt32(&connectCount, 1)
	if count <= maxConnect {
		newSession(ip, addr, conn)
	}
}

type WebConn struct {
	*websocket.Conn
	remain  []byte    // 剩余未处理
}

func(c *WebConn)Read(b []byte) (int, error) {
	var msgType int
	var message []byte
	var err error
	if len(c.remain) > 0 {
		msgType, message, err = websocket.BinaryMessage, c.remain, nil
	} else {
		msgType, message, err = c.ReadMessage()
		if err != nil {
			return 0, err
		}
	}
	if msgType == websocket.TextMessage || msgType == websocket.BinaryMessage {
		copyLen := copy(b, message)
		if copyLen < len(message) {
			c.remain = message[:copyLen]
		} else {
			c.remain = nil
		}
		return copyLen, nil
	}
	return 0, nil
}

func(c *WebConn)Write(b []byte) (int, error) {
	err := c.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		return 0, nil
	}
	return len(b), nil
}

func(c *WebConn)SetDeadline(t time.Time) error {
	return nil
}

var upgrader = websocket.Upgrader{
	// 支持跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWeb(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("web connect error:", err.Error())
		return
	}
	conn := &WebConn{Conn: c}
	if atomic.LoadInt32(&connectCount) >= maxConnect {
		// TODO: IP实现服务自发现
		fail := makeMsg(&Handshake{
			Code: 1,
			Msg:  "服务器忙",
		})
		conn.Write(fail)
		conn.Close()
	} else {
		addr := conn.RemoteAddr().String()
		log.Info("web connect:", addr)
		ipStr, _, _ := net.SplitHostPort(addr)
		ip := util.IpToUint32(ipStr)
		if addIpLimit(ip) {
			success := makeMsg(&Handshake{
				Ip: []string{"127.0.0.1"},
			})
			conn.Write(success)
			handleClient(ip, addr, conn)
		}
	}
}

func webServer(config *AppConfig ){
	http.HandleFunc("/", handleWeb)
	log.Info("web listening on:", config.Web.Listen)
	http.ListenAndServe(config.Web.Listen, nil)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
