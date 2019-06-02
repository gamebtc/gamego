package room

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"net"
	"reflect"
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

// 大厅接口
type Haller interface {
	// Start
	Start()
	// 玩家上线
	UserOnline(sess *Session, user *model.User)
	// 玩家下线
	UserOffline(sess *Session)
	// 玩家重新上线
	UserReline(oldSess *Session, newSess *Session)
	// 帧更新
	Update()
}

// 房间,桌子管理器
var (
	logName         string
	hall            Haller                            // 房间消息处理器
	messageChan     chan interface{}                  // 消息队列消息
	messageHandlers [math.MaxUint16]func(*NetMessage) // 消息处理器
	eventHandlers   [math.MaxUint16]func(*GameEvent)  // 事件处理器
	sessions        map[model.UserId]*Session         // 所有真实玩家
	signal          *util.AppSignal
	coder           protocol.Coder
)

func Encode(v interface{}) ([]byte, error) {
	return coder.Encode(v)
}

func RegistMsg(id int32, arg interface{}) {
	t := reflect.TypeOf(arg)
	if _, ok := coder.GetMsgId(t); ok {
		log.Fatalf("message %v is already registered", t)
	}
	coder.SetMsgId(t, id)
}

func RegistHandler(id int32, arg interface{}, f func(*NetMessage)) {
	if f != nil {
		messageHandlers[id] = f
	}
	RegistMsg(int32(id), arg)
}

func RegistEvent(id int32, f func(*GameEvent)) {
	eventHandlers[id] = f
}

func GetUser(id model.UserId) *Session {
	if s, ok := sessions[id]; ok {
		return s
	}
	return nil
}

func AddUser(sess *Session) {
	sessions[sess.UserId] = sess
}

func RemoveUser(sess *Session) bool {
	uid := sess.UserId
	if s, ok := sessions[uid]; ok && s == sess {
		delete(sessions, uid)
		sess.unlockRoom()
		return true
	}
	return false
}

func UserCount() int {
	return len(sessions)
}

func Start(configName string, r Haller) {
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

// 关闭房间
func Close() {
	if signal.Close() {
	}
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

func Init(config *AppConfig, r Haller) error {
	hall = r
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
	GameId = roomInfo.Game
	CoinKey = roomInfo.CoinKey
	logName = "play" + CoinKey + "_" + strconv.Itoa(int(GameId))

	// 加载房间
	sessions = make(map[model.UserId]*Session, roomInfo.Cap*2)
	messageChan = make(chan interface{}, 65536+roomInfo.Cap*128)

	RegistMsg(int32(protocol.MsgId_ErrorInfo), (*protocol.ErrorInfo)(nil) )
	RegistMsg(int32(protocol.MsgId_LoginGameAck), (*protocol.LoginGameAck)(nil))

	log.Infof("room:%#v", roomInfo)
	return nil
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

func exec(m interface{}) {
	switch m := m.(type) {
	case *NetMessage:
		if f := messageHandlers[m.Id]; f != nil {
			if m.locked {
				f(m)
			}
		}
	case *GameEvent:
		if f := eventHandlers[m.Id]; f != nil {
			f(m)
		}
	case *Timer:
		m.Exec()
	case func():
		m()
	}
}

func roomConfigCheck(ver int32) int32 {
	defer util.PrintPanicStack()
	newConf, err := db.Driver.GetRoom(RoomId, ver)
	if err == nil && newConf != nil && newConf.Id == RoomId {
		ver = newConf.Ver
		Send(&GameEvent{Id: EventConfigChanged, Arg: newConf})
	}
	return ver
}

func startRoomConfigCheck(ver int32) {
	t := time.Tick(30 * time.Second)
	for {
		select {
		case <-t:
			ver = roomConfigCheck(ver)
		case <-signal.Die():
			return
		}
	}
}

//
func mainLoop() {
	// 帧更新周期
	hall.Start()
	period := time.Duration(Config.Period*1000) * time.Microsecond
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	go startRoomConfigCheck(Config.Ver)

	for {
		select {
		case m, ok := <-messageChan:
			if ok {
				exec(m)
			} else {
				return
			}
		case <-ticker.C: // 帧更新
			hall.Update()
		case <-signal.Die():
			return
		}
	}
}

var startKindSn int64 //起始值,包括
var endKindSn int64   //结束值,不包括
func NewKindSn() (sn int64) {
	const roundAllot = math.MaxUint16 + 1
	if startKindSn < endKindSn {
		sn = startKindSn
		startKindSn++
	} else if newStart := db.Driver.NewSN(GameId, roundAllot); newStart > 0 {
		sn = newStart
		startKindSn = newStart + 1
		endKindSn = newStart + roundAllot
	}
	return
}

var startRoundId int64 //起始值,包括
var endRoundId int64   //结束值,不包括
func NewGameRoundId() (sn int64) {
	const roundAllot = 1024
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
		flow.Sn = NewKindSn()
		for flow.Sn == 0 {
			time.Sleep(time.Second)
			flow.Sn = NewKindSn()
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

func NewRand() *rand.Rand {
	bin := make([]byte, 8)
	crand.Read(bin)
	seed := binary.LittleEndian.Uint64(bin)
	return rand.New(rand.NewSource(int64(seed)))
}
