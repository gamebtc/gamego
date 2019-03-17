package internal

import (
	"io"
	"math"
	"net"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

const (
	SESS_INIT     = 0   // 初始化
	SESS_VERCHECK = 1   // 已通过版本检查
	SESS_LOGINING = 2   // 登录中
	SESS_LOGINED  = 3   // 登录成功
	SESS_CLOSE    = 100 // 是否断开
)

// 配置
var (
	rpmLimit     int32
	sendChanLen  int
	recvChanLen  int
	readTimeout  time.Duration
	writeTimeout time.Duration
	slowOpNano   int64
)

var (
	closeBin = &GameFrame{Data:make([]byte, HeadLen)}
)

func init(){
	SetHead(closeBin.Data, int32(MsgId_Control))
}

type Session struct {
	Id      int64     // 连接唯一ID
	Ip      uint32    // IP地址
	Seed    uint32    // 接收加密种子
	Addr    string    // 玩家IP地址
	Act     string    // 玩家账号
	UserId  UserId    // 玩家ID
	RoomId  int32     // 房间服ID
	Coder             // 编码解码
	Flag    int32     // 会话标记(0:初始化，1:已通过版本检查，2:登录中，3:登录成功, 4:已关闭)
	Created time.Time // TCP链接建立时间

	gameStream GameStream      // 后端房间服数据流
	services   [8]GameClient   // 服务连接
	conn       net.Conn        // 底层网络连接
	totalRecv  int64           // 总接收字节数
	dieChan    chan struct{}   // 会话关闭信号
	dieOnce    int32           // 会话关闭保护
	callCtx    context.Context // 调用上下文
}

// mail goroutine
func (sess *Session) mainLoop() {
	// TODO:1分钟还没登录成功的断开连接
	defer util.PrintPanicStack()
	h := [HeadLen]byte{}
	head := h[:]
	for sess.Flag < SESS_CLOSE {
		// 设置读取超时时间,读取消息长度
		sess.conn.SetReadDeadline(time.Now().Add(readTimeout))
		n, err := io.ReadFull(sess.conn, head)
		if err != nil || n != HeadLen {
			return
		}
		size, id := GetHead(head)
		if size > math.MaxUint16 || id > math.MaxUint16 {
			// 大小保护
			log.Warnf("user:%v,msg:%v,addr:%v,size:%v", sess.UserId, id, sess.Addr, size)
			return
		}
		payload := make([]byte, HeadLen+size)
		if size > 0 {
			// 10秒内读完消息内容
			sess.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
			n, err = io.ReadFull(sess.conn, payload[HeadLen:])
			if err != nil || n != int(size) {
				return
			}
		}
		sess.totalRecv += int64(size + HeadLen)
		copy(payload[:HeadLen], head)
		sess.processData(payload)

		// deliver the data to the input queue
		select {
		case <-sess.dieChan: // sess.Close
			sess.Flag = SESS_CLOSE
		case <-signal.Die(): // 服务关闭
			sess.Flag = SESS_CLOSE
		default:
		}
	}
}

func (sess *Session) send(val interface{}) bool {
	var data []byte
	switch val := val.(type) {
	case []byte:
		data = val
	case error:
		if buf, e := sess.Encode(makeError(0, 1001, "调用失败", val.Error())); e == nil {
			data = buf
		}
	default:
		if buf, e := sess.Encode(val); e == nil {
			data = buf
		}
	}
	if data == nil {
		return true
	}
	return sess.write(data)
}

func(sess *Session) write(data []byte) bool {
	//sess.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if _, err := sess.conn.Write(data); err != nil {
		return false
		//switch err := err.(type) {
		//case net.Error:
		//	if err.Timeout() == false {
		//		return false
		//	}
		//default:
		//	return false
		//}
	}
	return true
}

func (sess *Session) processData(data []byte) {
	if result, err := sess.route(data); err == nil {
		if result != nil {
			if sess.send(result) == false {
				sess.Flag = SESS_CLOSE
			}
		}
	} else {
		// 发送严重错误，然后关闭客户端
		id := GetHeadId(data)
		sess.send(&FatalInfo{ReqId: id, Code: 1001, Msg: "严重错误", Key: err.Error()})
		sess.Flag = SESS_CLOSE
		log.Warnf("user:%v,api:%v,addr:%v,err:%v", sess.UserId, id, sess.Addr, err.Error())
	}
}

//每个玩家1个线程，客户端网络写1个,主线程接收1个，房间消息流接收1个
func (sess *Session) Start() {
	sess.SetUser(0)
	sess.mainLoop()
}

// 线程安全关闭
func (sess *Session) Close() {
	if atomic.CompareAndSwapInt32(&sess.dieOnce, 0, 1) {
		close(sess.dieChan)
		// 通知服务器退出
		if sess.UserId != 0 {
			if server := sess.getServer(0); server != nil {
				server.Call(sess.callCtx, closeBin)
			}
		}
	}
}

func (sess *Session) closeRoom() {
	if c := sess.gameStream; c != nil {
		c.Close()
		sess.RoomId = 0
		sess.gameStream = nil
	}
}

func (sess *Session) destroy() {
	//sess.once.Do(func(){
	sess.closeRoom()
	sess.Close()
	sess.Flag = SESS_CLOSE
	//})
}

func (sess *Session) SetUser(uid int32) {
	sess.UserId = uid
	sess.callCtx = NewServerHead(sess.Id, uid, sess.Ip)
}

// route client protocol
func (sess *Session) route(data []byte) (ret interface{}, err error) {
	start := time.Now()
	id := GetHeadId(data)
	// 需要登录才能调用的协议
	if id >= int32(MsgId_UserMessageHeadSplit) && sess.Flag != SESS_LOGINED {
		err = ErrorUnauthorized
		return
	}
	// 协议号的划分采用分割协议区间，转发到不同的后端服务
	if id >= int32(MsgId_GameMessageHeadSplit) {
		// 转发到游戏房间
		ret, err = sess.forwardToGame(id, data)
	} else {
		// 代理服务器优先处理的协议
		if h := handlers[id]; h != nil {
			ret, err = h(sess, data)
		} else {
			// 转发到其他服务器
			ret, err = sess.callServer(id, data)
		}
	}
	// 记录慢操作
	if slowOpNano > 0 {
		if elapsed := time.Now().UnixNano() - start.UnixNano(); elapsed > slowOpNano {
			log.Warnf("user:%v, api:%v,cost:%v ms", sess.UserId, id, elapsed/Ms2Na)
		}
	}
	return
}

func (sess *Session) readFromGame(roomId int32, stream GameStream) {
	defer stream.Close()
	for sess.Flag < SESS_CLOSE {
		if bin, err := stream.Recv(); err == nil {
			select {
			case <-sess.dieChan:
				return
			default:
				if sess.write(bin) == false {
					return
				}
			}
		} else {
			sess.SendError(0, 2002, "房间连接已断开", err.Error())
			return
		}
	}
}

func makeError(id int32, code int32, m string, k string) *ErrorInfo {
	return &ErrorInfo{ReqId: id, Code: code, Msg: m, Key: k}
}

func (sess *Session) SendError(id int32, code int32, m string, k string) {
	sess.send(&ErrorInfo{ReqId: id, Code: code, Msg: m, Key: k})
}

// forward messages to game server
func (sess *Session) forwardToGame(id int32, data []byte) (interface{}, error) {
	// check stream
	if sess.gameStream != nil {
		// forward the frame to game
		if e := sess.gameStream.Send(data); e != nil {
			sess.closeRoom()
			return makeError(id, 2002, "房间连接已断开", e.Error()), nil
		}
		log.Debugf("forwardToGame success %v, %v", id, data)
	} else {
		return makeError(id, 2001, "房间未连接", ""), nil
	}
	return nil, nil
}

func (sess *Session) getServer(sid byte) GameClient {
	server := sess.services[sid]
	if server == nil {
		if conn := rpcServicePool.GetService(int32(sid)); conn != nil {
			server = NewGameClient(conn)
		}
	}
	return server
}

// call other server
func (sess *Session) callServer(id int32, data []byte) (interface{}, error) {
	sid := byte((id & 8191) >> 10)
	server := sess.getServer(sid)
	if server != nil {
		if f, e := server.Call(sess.callCtx, &GameFrame{Data: data}); e == nil {
			return f.Data, nil
		} else {
			return makeError(id, 1001, "调用失败，请您稍后重试！", e.Error()), nil
		}
	} else {
		return makeError(id, 1001, "服务器繁忙，请您稍后重试！", ""), nil
	}
	return nil, nil
}

func (sess *Session) getContext() context.Context {
	return sess.callCtx
}

// 登录游戏房间
func (sess *Session) loginRoom(roomId int32, req []byte, stream GameStream) (interface{}, error) {
	if e := stream.Send(req); e != nil {
		return makeError(int32(MsgId_LoginRoomReq), 2001, "房间连接失败，请您稍后重试！", e.Error()), nil
	}
	data, e := stream.Recv()
	if e == nil && len(data) >= HeadLen {
		if id := GetHeadId(data); id == int32(MsgId_LoginRoomAck) {
			ack := LoginRoomAck{}
			if e = sess.Unmarshal(data[HeadLen:], &ack); e == nil && ack.Code == 0 {
				if oldStream := sess.gameStream; oldStream != nil {
					oldStream.Close()
				}
				sess.gameStream = stream
				sess.RoomId = roomId
				go sess.readFromGame(roomId, stream)
			}
		}
		return data, e
	} else {
		return makeError(int32(MsgId_LoginRoomReq), 2001, "房间连接失败，请您稍后重试！", e.Error()), nil
	}
}
