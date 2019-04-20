package internal

import (
	"context"
	"encoding/binary"
	"math"
	"reflect"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"agent/conf"

	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

var (
	handlers [MsgId_GameMessageHeadSplit]func(*Session, []byte) (interface{}, error)
	signal   *util.AppSignal
)

func registMsg(id MsgId_Code, msg interface{}) {
	t := reflect.TypeOf(msg)
	if _, ok := coder.GetMsgId(t); ok {
		log.Fatalf("message %v is already registered", t)
	}
	coder.SetMsgId(t, int32(id))
}

func Run(config *conf.AppConfig) {
	// listeners
	checkVerSign = config.VerSign
	tcpReadBuf = config.Tcp.ReadBuf
	tcpWriteBuf = config.Tcp.WriteBuf
	sendChanLen = config.SendChanLen
	recvChanLen = config.RecvChanLen
	readTimeout = time.Duration(config.ReadTimeout) * time.Second
	writeTimeout = time.Duration(config.WriteTimeout) * time.Second
	slowOpNano = int64(config.SlowOp) * Ms2Na
	rpmLimit = config.RpmLimit

	coder = GetCoder(config.Codec)

	registMsg(MsgId_ErrorInfo, (*ErrorInfo)(nil))
	registMsg(MsgId_HandshakeAck, (*Handshake)(nil))
	registMsg(MsgId_VerCheckAck, (*VerCheckAck)(nil))
	registMsg(MsgId_HeartBeatAck, (*HeartBeatAck)(nil))
	registMsg(MsgId_UserLoginFailAck, (*LoginFailAck)(nil))
	registMsg(MsgId_LoginRoomAck, (*LoginRoomAck)(nil))
	registMsg(MsgId_UserLoginSuccessAck, (*LoginSuccessAck)(nil))
	registMsg(MsgId_ExitRoomAck, (*ExitRoomAck)(nil))

	handlers[MsgId_HandshakeReq] = handshakeHandler
	handlers[MsgId_UserLoginReq] = userLoginHandler
	handlers[MsgId_VerCheckReq] = verCheckHandler
	//handlers[MsgId_HeartBeatReq] = heartBeatHandler
	handlers[MsgId_LoginRoomReq] = loginRoomHandler
	handlers[MsgId_ExitRoomReq] = exitRoomHandler
	//handlers[MsgId_ConnectRoomReq] = connectRoomHandler

	//user := &Handshake{}
	//
	//// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	//getValue := reflect.ValueOf(user)
	//
	//// 一定要指定参数为正确的方法名
	//// 2. 先看看带有参数的调用方法
	//methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	//args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	//methodValue.Call(args)

	signal = util.NewAppSignal()
	signal.Run(func() {
		startServer(config)
	})
}

const (
	allotAgentIdCount = math.MaxUint16 + 1
)

var (
	startAgentId int64 //起始值,包括
	endAgentId   int64 //结束值,不包括
	lockAgentId  sync.Mutex
	allocAgentId = &GameFrame{Data: make([]byte, HeadLen+4)}
)

func init() {
	SetHead(allocAgentId.Data, int32(MsgId_AllocAgentId))
	binary.BigEndian.PutUint32(allocAgentId.Data[HeadLen:], allotAgentIdCount)
}

func newAgentId() (id int64) {
	lockAgentId.Lock()
	defer lockAgentId.Unlock()
	if startAgentId < endAgentId {
		id = startAgentId
		startAgentId++
	} else if conn := rpcServicePool.GetService(0); conn != nil {
		server := NewGameClient(conn)
		ackFrame, err := server.Call(context.Background(), allocAgentId)
		if err == nil && len(ackFrame.Data) >= 8 {
			if newStart := int64(binary.BigEndian.Uint64(ackFrame.Data)); newStart > 0 {
				id = newStart
				startAgentId = newStart + 1
				endAgentId = newStart + allotAgentIdCount
			}
		}
	}
	return
}
