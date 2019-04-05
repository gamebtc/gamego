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
	userOffline = &GameFrame{Data: make([]byte, HeadLen)}
)

func init() {
	SetHead(userOffline.Data, int32(MsgId_UserOffline))
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
	stopRecv   chan struct{}   // 读取停止信号
	recvTag    int32           // 接收到消息信号
	dieChan    chan struct{}   // 会话关闭信号
	disposed   uint32          // 会话关闭标记
	callCtx    context.Context // 调用上下文
}

// 接收线程同步处理所有请求消息
func (sess *Session) recvLoop() {
	defer util.PrintPanicStack()
	defer close(sess.stopRecv)
	h := [HeadLen + math.MaxUint16]byte{}
	head := h[:HeadLen]
	for sess.Flag < SESS_CLOSE {
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
		payload := h[:HeadLen+size]
		if size > 0 {
			n, err = io.ReadFull(sess.conn, payload[HeadLen:])
			if err != nil || n != int(size) {
				return
			}
		}
		sess.totalRecv += int64(size + HeadLen)
		sess.processData(payload)
		atomic.StoreInt32(&sess.recvTag, 0)
	}
}

// mail goroutine
func (sess *Session) mainLoop() {
	// TODO:1分钟还没登录成功的断开连接
	defer util.PrintPanicStack()
	defer sess.Close()
	t := time.NewTicker(30 * time.Second)
	for sess.Flag < SESS_CLOSE {
		// deliver the data to the input queue
		select {
		case <-sess.dieChan: // sess.Close
			sess.Flag = SESS_CLOSE
		case <-signal.Die(): // 服务关闭
			sess.Flag = SESS_CLOSE
		case <-sess.stopRecv:
			sess.Flag = SESS_CLOSE
		case <-t.C:
			if atomic.AddInt32(&sess.recvTag, 1) >= 3 {
				sess.Flag = SESS_CLOSE
			}
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

func (sess *Session) write(data []byte) bool {
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
	result, err := sess.route(data)
	if err != nil {
		// 发送严重错误，然后关闭客户端
		id := GetHeadId(data)
		sess.send(&FatalInfo{ReqId: id, Code: 1001, Msg: "严重错误", Key: err.Error()})
		sess.Flag = SESS_CLOSE
		log.Warnf("user:%v,api:%v,addr:%v,err:%v", sess.UserId, id, sess.Addr, err.Error())
		return
	}
	if result != nil {
		if sess.send(result) == false {
			sess.Flag = SESS_CLOSE
		}
	}
}

//每个玩家1个线程,所有消息同步处理
func (sess *Session) Start() {
	sess.SetUser(0)
	go sess.recvLoop()
	sess.mainLoop()
}

// 线程安全关闭
func (sess *Session) Close() {
	if atomic.CompareAndSwapUint32(&sess.disposed, 0, 1) {
		sess.Flag = SESS_CLOSE
		sess.closeRoom()
		close(sess.dieChan)
		// 通知服务器退出
		if sess.UserId != 0 {
			if server := sess.getServer(0); server != nil {
				server.Call(sess.callCtx, userOffline)
			}
		}
	}
}

func (sess *Session) closeRoom() {
	if c := sess.gameStream; c != nil {
		log.Debugf("close room:%v, %v", sess.Id, sess.RoomId)
		c.Close()
		sess.RoomId = 0
		sess.gameStream = nil
	}
}

func (sess *Session) SetUser(uid UserId) {
	sess.UserId = uid
	sess.callCtx = NewServerHead(sess.Id, uid, sess.Ip)
}

// route client protocol
func (sess *Session) route(data []byte) (ret interface{}, err error) {
	start := time.Now()
	id := GetHeadId(data)
	if id < int32(MsgId_UserMessageHeadSplit) {
		// 内部协议，用户不能调用
		err = ErrorUndefined
		return
	}

	// 需要登录才能调用的协议
	if id >= int32(MsgId_UserLoginMessageSplit) && sess.Flag != SESS_LOGINED {
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
	log.Debugf("read room:%v, %v", sess.Id, roomId)
	defer stream.Close()
loop:
	for sess.Flag < SESS_CLOSE {
		if bin, err := stream.Recv(); err == nil {
			select {
			case <-sess.dieChan:
				break loop
			default:
				if sess.write(bin) == false {
					break loop
				}
			}
		} else {
			sess.SendError(0, 2002, "房间连接已断开", err.Error())
			break loop
		}
	}
	log.Debugf("close read:%v, %v", sess.Id, roomId)
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
	if server == nil {
		return makeError(id, 1001, "服务器繁忙，请您稍后重试！", ""), nil
	}

	f, e := server.Call(sess.callCtx, &GameFrame{Data: data})
	if e != nil {
		return makeError(id, 1001, "调用失败，请您稍后重试！", e.Error()), nil
	}
	return f.Data, nil
}

func (sess *Session) getContext() context.Context {
	return sess.callCtx
}

// 登录游戏房间
func (sess *Session) loginRoom(roomId int32, req []byte, stream GameStream) (interface{}, error) {
	log.Debugf("start login Room:%v,a:%v;uid:%v", roomId, sess.Id, sess.UserId)
	if e := stream.Send(req); e != nil {
		return makeError(int32(MsgId_LoginRoomReq), 2001, "房间连接失败，请您稍后重试！", e.Error()), nil
	}
	data, e := stream.Recv()
	if e != nil {
		return makeError(int32(MsgId_LoginRoomReq), 2001, "房间连接失败，请您稍后重试！", e.Error()), nil
	}
	if len(data) < HeadLen {
		return makeError(int32(MsgId_LoginRoomReq), 2001, "房间连接失败，请您稍后重试！", "接收数据长度错误"), nil
	}

	id := GetHeadId(data)
	if id == int32(MsgId_LoginRoomAck) {
		ack := LoginRoomAck{}
		e = sess.Unmarshal(data[HeadLen:], &ack)
		if e == nil && ack.Code == 0 {
			sess.closeRoom()
			sess.gameStream = stream
			sess.RoomId = roomId
			go sess.readFromGame(roomId, stream)
		}
		log.Debugf("loginRoom:%v,%v", e, ack)
	} else {
		log.Debugf("loginRoom error2:%v", id)
	}
	return data, e
}
