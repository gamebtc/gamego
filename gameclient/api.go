package main

import (
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

var (
	handlers  [65536]func(*Session, interface{})
	signal    *util.AppSignal
	coder     Coder //编码方式
	jsonCoder Coder //JSON编码
	conf      *AppConfig
)

func Run(config *AppConfig) {
	conf = config
	// listeners
	signal = util.NewAppSignal()
	slowOpNano = int64(config.SlowOp) * Ms2Na
	rpmLimit = config.RpmLimit
	sessionId = uint32(rand.Intn(math.MaxInt32))

	coder = GetCoder(config.Codec)
	jsonCoder = GetCoder("json") //jsonIter

	registMsg(MsgId_UserLoginReq, &LoginReq{}, nil)
	registMsg(MsgId_LoginRoomReq, &LoginRoomReq{}, nil)
	registMsg(MsgId_LoginRoomAck, &LoginRoomAck{}, loginRoom)
	registMsg(MsgId_BetReq, &BetReq{}, nil)

	registMsg(MsgId_BetAck, &BetAck{}, betAck)
	registMsg(MsgId_HandshakeAck, &Handshake{}, handshake)
	registMsg(MsgId_UserLoginFailAck, &LoginFailAck{}, loginFail)
	registMsg(MsgId_UserLoginSuccessAck, &LoginSuccessAck{}, loginSuccess)
	registMsg(MsgId_ErrorInfo, &ErrorInfo{}, showErrorInfo)

	registMsg(MsgId_VerCheckReq, &VerCheckReq{}, nil)
	registMsg(MsgId_VerCheckAck, &VerCheckAck{}, verCheck)

	registMsg(MsgId_HeartBeatReq, &HeartBeatReq{}, nil)
	registMsg(MsgId_HeartBeatAck, &HeartBeatAck{}, heartBeat)


	registMsg(MsgId_FolksGameInitAck, &FolksGameInitAck{}, folksGameInit)

	//handlers[MsgId_HandshakeReq] = handshakeHandler
	//handlers[MsgId_UserLoginReq] = userLoginHandler
	//handlers[MsgId_VerCheckReq] = verCheckHandler
	//handlers[MsgId_HeartBeatReq] = heartBeatHandler
	//handlers[MsgId_LoginRoomReq] = loginRoomHandler
	//handlers[MsgId_ConnectRoomReq] = connectRoomHandler

	signal.Run(func() {
		go startServer(config)
	})
}

func folksGameInit(sess *Session, arg interface{}) {
	if arg, ok := arg.(*FolksGameInitAck); ok && arg != nil {
		log.Debugf("folksGameInitAck:%#v", arg)
	}
}

func verCheck(sess *Session, arg interface{}) {
	if arg, ok := arg.(*VerCheckAck); ok && arg != nil {
		log.Debugf("verCheckAck:%#v", arg)
	}
}
func loginRoom(sess *Session, arg interface{}) {
	if arg, ok := arg.(*LoginRoomAck); ok && arg != nil {
		sess.roomId = arg.Room
		log.Debugf("loginRoom:%#v", arg)
	}
}

func heartBeat(sess *Session, arg interface{}) {
	if arg, ok := arg.(*HeartBeatAck); ok && arg != nil {
		log.Debugf("heartBeat:uid%v, id:%v", sess.UserId, arg.Id)
	}
}

func betAck(sess *Session, arg interface{}) {
	if arg, ok := arg.(*BetAck); ok && arg != nil {
		log.Debugf("betAck:%#v", arg)
	}
}

func handshake(sess *Session, arg interface{}) {
	if arg, ok := arg.(*Handshake); ok && arg != nil {
		log.Debugf("handshake:%#v", arg)
	}
}
func showErrorInfo(sess *Session, arg interface{}) {
	if arg, ok := arg.(*ErrorInfo); ok && arg != nil {
		log.Debugf("showErrorInfo:%v", arg)
	}
}

func loginFail(sess *Session, arg interface{}) {
	if arg, ok := arg.(*LoginFailAck); ok && arg != nil {
		log.Debugf("LoginFail:%v", arg)
	}
}

func loginSuccess(sess *Session, arg interface{}) {
	if arg, ok := arg.(*LoginSuccessAck); ok && arg != nil {
		sess.UserId = arg.Id
		sess.Act = arg.Act
		//addUser(sess)
		log.Debugf("LoginSuccess:uid:%v,arg:%v", arg.Id, arg)
	} else {
		log.Debugf("LoginError:%#v", arg)
	}
}

func startServer(config *AppConfig) {
	log.Debugf("start:%v", config.Pprof)
	http.HandleFunc("/cmd", cmd)
	//http.HandleFunc("/login", login)
	if err := http.ListenAndServe(config.Pprof, nil); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(id int64)(user *Session){
	//创建客户端连接
	server := conf.Server
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Errorf("Fatal error %s: %s", server, err.Error())
		return nil
	}
	user = &Session{
		Id:       id,
		Addr:     server,
		Created:  time.Now(),
		Coder:    coder,
		conn:     conn,
		recvChan: make(chan []byte, recvChanLen),
		sendChan: make(chan interface{}, sendChanLen),
		stopSend: make(chan struct{}),
		dieChan:  make(chan struct{}),
	}
	addSession(user)
	go StartSession(user)
	return
}

func cmd(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	uid, err := strconv.Atoi(query.Get("uid"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user := getSession(int64(uid))
	if user == nil{
		user = CreateUser(int64(uid))
	}
	if user == nil {
		w.Write([]byte("user not find"))
		return
	}

	cid, err := strconv.Atoi(query.Get("msg"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if cid == -1{
		user.Close()
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//user
	req, err := jsonDecode(int32(cid), buf)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	log.Debugf("cid：%v, cmd:%#v", cid, req)
	if user.send(req) {
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte("false"))
	}
}

func jsonDecode(id int32, buf []byte) (msg interface{}, err error) {
	if i := jsonCoder.GetHandler(id); i != nil && i.Type != nil {
		msg = reflect.New(i.Type.Elem()).Interface()
		err = jsonCoder.Unmarshal(buf, msg)
	} else {
		msg = buf
		log.Fatalf("jsonCoder:%#v", i)
	}
	return
}

func registMsg(id MsgId_Code, msg interface{}, h func(*Session, interface{})) {
	t := reflect.TypeOf(msg)
	if _, ok := coder.GetMsgId(t); ok {
		log.Fatalf("message %v is already registered", t)
	}
	coder.SetHandler(t, int32(id), nil)
	jsonCoder.SetHandler(t, int32(id), nil)
	if h != nil {
		handlers[id] = h
	}
}
