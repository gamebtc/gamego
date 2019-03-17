package room

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

// 游戏用户消息
type NetMessage struct {
	Id  int32       // 消息ID
	Arg interface{} // 参数
	*Session
}

// 游戏事件
type GameEvent struct {
	Id  int32       // 事件ID
	Arg interface{} // 参数
}

// 房间接口
type Roomer interface {
	// 初始化房间
	Init(*model.RoomInfo)
	// 获取玩家信息
	GetUser(int32) *Session
	// 设置玩家信息
	AddUser(sess *Session)
	// 删除玩家信息
	RemoveUser(sess *Session) bool
	// 玩家上线
	UserOnline(sess *Session, user *model.User, coin int64)
	// 玩家下线
	UserOffline(sess *Session)
	// 玩家重新上线
	UserReline(oldSess *Session, newSess *Session)
	// 处理消息
	Exec(interface{})
	// 帧更新
	Update()
}

// 房间,桌子管理器
var (
	roomer      Roomer           // 房间消息处理器
	messageChan chan interface{} // 消息队列消息
	logName     string
)

type DefaultRoomer struct {
	MessageHandler [math.MaxUint16]func(*NetMessage) // 消息处理器
	EventHandler   [math.MaxUint16]func(*GameEvent)  // 事件处理器
	Users          map[int32]*Session                // 在线玩家
}

func (r *DefaultRoomer) RegistHandler(id protocol.MsgId_Code, arg interface{}, f func(*NetMessage)) {
	if f != nil {
		r.MessageHandler[id] = f
	}
	RegistMsg(id, arg)
}

func (r *DefaultRoomer) Init(info *model.RoomInfo) {
	r.Users = make(map[int32]*Session, info.Cap*2)

	RegistMsg(protocol.MsgId_ErrorInfo, &protocol.ErrorInfo{})
	RegistMsg(protocol.MsgId_LoginRoomAck, &protocol.LoginRoomAck{})
}

func (r *DefaultRoomer) GetUser(id int32) *Session {
	if s, ok := r.Users[id]; ok {
		return s
	}
	return nil
}

func (r *DefaultRoomer) AddUser(sess *Session) {
	r.Users[sess.UserId] = sess
}

func(r *DefaultRoomer) RemoveUser(sess *Session) bool{
	uid := sess.UserId
	if s, ok := r.Users[uid]; ok && s == sess{
		delete(r.Users, uid)
		return true
	}
	return false
}

func (r *DefaultRoomer) Exec(m interface{}) {
	switch m := m.(type) {
	case *NetMessage:
		if f := r.MessageHandler[m.Id]; f != nil {
			if m.UserId != 0 && m.Disposed == false {
				f(m)
			}
		}
	case *GameEvent:
		if f := r.EventHandler[m.Id]; f != nil {
			f(m)
		}
	case *Timer:
		m.Exec()
	case func():
		m()
	}
}

func Start(configName string, r Roomer) {
	defer util.PrintPanicStack()
	// open profiling
	config := InitConfig(configName)

	if err := Init(config, r); err != nil {
		panic(err)
	}

	signal = util.NewAppSignal()
	signal.Run(func() {
		if config.Tcp.Listen != "" {
			startServer(config)
		} else {
			startGrpc(config)
		}
		go mainLoop()
	})
}

func startGrpc(config *AppConfig) {
	lis, err := net.Listen("tcp", config.Grpc.Listen)
	if err != nil {
		panic(err)
	}
	gs := grpc.NewServer()
	s := &grpcServer{}
	protocol.RegisterGameServer(gs, s)
	protocol.RegisterGrpcServer(gs)

	err = protocol.RegistConsul(config.Consul.Addr, &config.Grpc)
	if err != nil {
		panic(err)
	}
	log.Info("starting service at:", lis.Addr())
	go gs.Serve(lis)
}

func Init(config *AppConfig, r Roomer) error {
	if d, err := db.CreateDriver(&config.Database); err != nil {
		return err
	} else {
		db.Driver = d
	}

	roomInfo, err := db.Driver.LockRoomServer(&config.Room)
	if err != nil {
		return err
	}
	if roomInfo == nil {
		return errors.New(fmt.Sprintf("room config not find:%#v", config.Room))
	}

	Config = *roomInfo
	RoomId = roomInfo.Id
	KindId = roomInfo.Kind
	CoinKey = roomInfo.CoinKey

	log.Infof("room:%#v", roomInfo)
	logName = "play" + CoinKey + "_" + strconv.Itoa(int(KindId))
	// 加载房间
	//d.Users = make(map[int32]*Session, roomInfo.Cap*2)
	messageChan = make(chan interface{}, 65536+roomInfo.Cap*128)
	roomer = r
	return nil
}

//
func regConsulRoom(config *AppConfig) {

}

func AfterCall(d time.Duration, f func()) *Timer {
	t := &Timer{f: f}
	t.t = time.AfterFunc(d, func() {
		messageChan <- t
	})
	return t
}

func Call(f func()) {
	messageChan <- f
}

func Send(m interface{}) {
	messageChan <- m
}

//
func mainLoop() {
	// 帧更新周期
	roomer.Init(&Config)
	period := time.Duration(Config.Period) * time.Millisecond
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	
	go func(ver int32){
		t := time.Tick(time.Minute)
		for {
			select {
			case <- t:
				ver = roomConfigCheck(ver)
			case <-signal.Die():
				return
			}
		}
	}(Config.Ver)

	for {
		select {
		case m, ok := <-messageChan:
			if ok {
				roomer.Exec(m)
			} else {
				return
			}
		case <-ticker.C: // 帧更新
			roomer.Update()
		case <-signal.Die():
			return
		}
	}
}

func roomConfigCheck(ver int32) int32{
	defer util.PrintPanicStack()
	newConf, err := db.Driver.GetRoom(RoomId, ver)
	if err == nil && newConf != nil && newConf.Id == RoomId {
		ver = newConf.Ver
		Send(&GameEvent{Id: EventConfigChanged, Arg: newConf})
	}
	return ver
}

// 关闭房间
func Close() {
}

// 游戏
type Gamer interface {
	Process(*Session, *protocol.GameFrame)
}

var startSn int64 //起始值
var countSn int64 //SN缓存数

func NewSn(count uint16) (sn int64) {
	allot := int64(count)
	if countSn >= allot {
		sn = startSn
		startSn += allot
		countSn -= allot
	} else if newStart := db.Driver.NewSN(KindId, math.MaxUint16); newStart > 0 {
		// 需要重新分配
		sn = newStart
		startSn = newStart + allot
		countSn = math.MaxUint16 - allot
	}
	return
}

var startRoundId int64 //起始值
var endRoundId int64   //结束值,不包括
const roundAllot = 4   //1024*math.MaxUint16
func NewGameRoundId() (sn int64) {
	if startRoundId < endRoundId {
		sn = startRoundId
		startRoundId++
	} else if newStart := db.Driver.NewSN(logName, roundAllot); newStart > 0 {
		sn = newStart
		startRoundId = newStart + 1
		endRoundId = newStart + roundAllot
	}
	return
}

// 同步写分
func WriteCoin(flow *model.CoinFlow) error {
	if flow.Sn == 0 {
		flow.Sn = NewSn(1)
		for flow.Sn == 0 {
			time.Sleep(time.Second)
			flow.Sn = NewSn(1)
		}
	}
	return db.Driver.BagDeal(CoinKey, flow)
}

func SaveLog(log interface{}) error {
	return db.Driver.SaveLog(logName, log)
}

func Now() int64 {
	return time.Now().Unix()
}
