package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"sync/atomic"

	"agent/conf"

	"local.com/abc/game/msg"
	"local.com/abc/game/util"
)

var (
	ipLocker        sync.RWMutex
	sameIpMap       map[uint32]uint32 //同一IP连接数限制
	sameIp          uint32            //同一IP连接数限制
	connectCount    int32             //当前连接数
	maxConnect      int32             //最大连接数
	coder           msg.Coder         //编码方式
	rpcServicePool  msg.ServicePool
	roomServicePool msg.ServicePool
)

func startServer(config *conf.AppConfig) {
	rpcServicePool = msg.NewServicePool(config.Consul.Addr, config.Consul.ServerPrefix, config.Consul.Services, true)

	if config.Consul.RoomPrefix != "" {
		roomServicePool = msg.NewServicePool(config.Consul.Addr, config.Consul.RoomPrefix, config.Consul.Services, false)
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
	if config.Udp.Listen != "" {
		go kcpServer(config)
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

func makeMsg(h *msg.Handshake) []byte {
	src, _ := coder.Encode(h)
	return src
}

func tcpServer(config *conf.AppConfig) {
	// resolve address & start listening
	l, err := net.Listen("tcp", config.Tcp.Listen)
	log.Info("listening on:", l.Addr())
	checkError(err)

	success := makeMsg(&msg.Handshake{
		Ip: []string{"127.0.0.1"},
	})
	// TODO: IP实现服务自发现
	fail := makeMsg(&msg.Handshake{
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
					conn.SetReadBuffer(config.Tcp.ReadBuf)
					// set socket write buffer
					conn.SetWriteBuffer(config.Tcp.WriteBuf)
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

func kcpServer(config *conf.AppConfig) {
	l, err := kcp.Listen(config.Udp.Listen)
	checkError(err)
	log.Info("udp listening on:", l.Addr())

	lis := l.(*kcp.Listener)
	if err := lis.SetReadBuffer(config.Udp.ReadBuf); err != nil {
		log.Fatalf("SetReadBuffer:", err)
	}
	if err := lis.SetWriteBuffer(config.Udp.WriteBuf); err != nil {
		log.Fatalf("SetWriteBuffer:", err)
	}
	if err := lis.SetDSCP(config.Udp.Dscp); err != nil {
		log.Fatalf("SetDSCP:", err)
	}

	success := makeMsg(new(msg.Handshake))
	fail := makeMsg(&msg.Handshake{
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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
