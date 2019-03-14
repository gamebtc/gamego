package internal

import (
	"math"
	"math/rand"
	"reflect"
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

func Run(config *conf.AppConfig) {
	// listeners
	tcpReadBuf = config.Tcp.ReadBuf
	tcpWriteBuf = config.Tcp.WriteBuf
	sendChanLen = config.SendChanLen
	recvChanLen = config.RecvChanLen
	readTimeout = time.Duration(config.ReadTimeout) * time.Second
	writeTimeout = time.Duration(config.WriteTimeout) * time.Second
	slowOpNano = int64(config.SlowOp) * Ms2Na
	rpmLimit = config.RpmLimit
	agentHead = int64(config.AgentId) << 32
	sessionId = uint32(rand.Intn(math.MaxInt32))

	coder = GetCoder(config.Codec)

	registMsg(MsgId_ErrorInfo, &ErrorInfo{})
	registMsg(MsgId_HandshakeAck, &Handshake{})
	registMsg(MsgId_VerCheckAck, &VerCheckAck{})
	registMsg(MsgId_HeartBeatAck, &HeartBeatAck{})
	registMsg(MsgId_UserLoginFailAck, &LoginFailAck{})
	registMsg(MsgId_LoginRoomAck, &LoginRoomAck{})
	registMsg(MsgId_UserLoginSuccessAck, &LoginSuccessAck{})

	handlers[MsgId_HandshakeReq] = handshakeHandler
	handlers[MsgId_UserLoginReq] = userLoginHandler
	handlers[MsgId_VerCheckReq] = verCheckHandler
	//handlers[MsgId_HeartBeatReq] = heartBeatHandler
	handlers[MsgId_LoginRoomReq] = loginRoomHandler
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

func registMsg(id MsgId_Code, msg interface{}) {
	t := reflect.TypeOf(msg)
	if _, ok := coder.GetMsgId(t); ok {
		log.Fatalf("message %v is already registered", t)
	}
	coder.SetMsgId(t, int32(id))
}
