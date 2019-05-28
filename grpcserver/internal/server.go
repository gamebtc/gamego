package internal

import (
	"context"
	"encoding/binary"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"time"

	"github.com/ipipdotnet/ipdb-go"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

var (
	verSignError = &VerCheckAck{Code: 100, Msg: "签名错误"}
	driver       db.GameDriver
	appConf      *AppConfig
	coder        Coder
	city  		 *ipdb.City
)

func init(){
	city, _ = ipdb.NewCity("ipipfree.ipdb")
	//city.Reload("ipipfree.ipdb")
}

// server is used to implement protocol.GameService.
type Server struct {
	ServerBase
}

func (s *Server) Init(config *AppConfig) (err error) {

	appConf = config
	coder = GetCoder(config.Codec)
	s.Coder = coder
	s.SetSlowOp(int64(config.SlowOp))

	s.RegistMsg(MsgId_ErrorInfo, &ErrorInfo{})

	s.RegistHandler(MsgId_UserOffline, nil, userOffline)
	s.RegistHandler(MsgId_AllocAgentId, nil, allocAgentId)

	s.RegistHandler(MsgId_VerCheckReq, &VerCheckReq{}, verCheck)
	s.RegistMsg(MsgId_VerCheckAck, &VerCheckAck{})

	s.RegistHandler(MsgId_UserLoginReq, &LoginReq{}, userLogin)
	s.RegistMsg(MsgId_UserLoginSuccessAck, &LoginSuccessAck{})
	s.RegistMsg(MsgId_UserLoginFailAck, &LoginFailAck{})

	s.RegistHandler(MsgId_HeartBeatReq, &HeartBeatReq{}, heartBeat)
	s.RegistMsg(MsgId_HeartBeatAck, &HeartBeatAck{})

	if driver, err = db.CreateDriver(&config.Database); err != nil {
		panic(err)
	}

	if config.Pprof != "" {
		go http.ListenAndServe(config.Pprof, nil)
	}

	return
}

func (s *Server) Call(ctx context.Context, req *GameFrame) (res *GameFrame, err error) {
	// id
	id, arg, err := s.Decode(req.Data)
	if err != nil {
		return
	}
	user := ParseUserContext(ctx)
	// 检查登录
	if id >= int32(MsgId_UserLoginMessageSplit) {
		// 登录已过期
		if (user.UserId == 0) || (!driver.CheckUserAgent(user.UserId, user.AgentId)) {
			var buf []byte
			noLogin := &ErrorInfo{
				ReqId: id,
				Code:  10091,
				Msg:   "登录已过期",
				Key:   "",
			}
			if buf, err = s.Encode(noLogin); err == nil {
				res = &GameFrame{Data: buf}
			}
			return
		}
	}

	i := s.GetHandler(id)
	if i == nil || i.Handler == nil {
		err = fmt.Errorf("message %v not registered", id)
		return
	}

	defer func() {
		if e := recover(); e != nil {
			stack := string(debug.Stack())
			log.Errorf("call err:%v, id:%v, msg:%#v, stack:%v", e, id, arg, stack)
			err, _ = e.(error)
		}
	}()

	// 调用消息
	start := time.Now()
	r := i.Handler(user, arg)
	if r != nil {
		// 对返回的结果进行编码
		switch r := r.(type) {
		case *GameFrame:
			res = r
		case GameFrame:
			res = &r
		case []byte:
			res = &GameFrame{Data: r}
		case error:
			err = r
		default:
			var buf []byte
			if buf, err = s.Encode(r); err == nil {
				res = &GameFrame{Data: buf}
			}
		}
	} else {
		res = &GameFrame{}
	}
	// 慢消息检查
	if s.SlowOpNa > 0 {
		if elapse := time.Since(start).Nanoseconds(); elapse > s.SlowOpNa {
			log.Warningf("elapse:%v(ms), id:%v, arg:%v", elapse/Ms2Na, id, arg)
		}
	}
	return
}

// 版本检查
func verCheck(ctx context.Context, in interface{}) interface{} {
	arg := in.(*VerCheckReq)
	pack := driver.GetPackConf(arg.Env.Pack)
	if pack == nil {
		return &VerCheckAck{
			Code: 101,
			Msg:  fmt.Sprintf("不支持的包[%v])", arg.Env.Pack),
		}
	}
	var conf map[string]string
	if pack != nil {
		conf = pack.Conf
		// 检查客户端版本号
		if len(pack.CanVer) > 0 {
			for _, v := range pack.CanVer {
				if v == arg.Env.Ver {
					ack := &VerCheckAck{Code: 0, Conf: conf}
					// 获取IP
					user := ParseUserContext(ctx)
					strIp := util.Uint32ToIp(uint32(user.Ip))
					if strIp == "127.0.0.1" {
						strIp = "8.8.8.8"
					}

					if c, err := city.FindInfo(strIp, "CN"); err == nil {
						ack.Country = c.CountryName
						ack.Region = c.RegionName
						ack.City = c.CityName
					}
					return ack
				}
			}
		}
	}
	return &VerCheckAck{
		Code: 101,
		Msg:  fmt.Sprintf("不支持的版本[%v]:%v)", arg.Env.Pack, arg.Env.Ver),
		Url:  pack.Url,
		Conf: conf,
	}
}

func heartBeat(ctx context.Context, in interface{}) interface{} {
	arg := in.(*HeartBeatReq)
	user := ParseUserContext(ctx)
	log.Debugf("heartBeat:uid%v, id:%v", user.UserId, arg.Id)
	return &HeartBeatAck{Id: arg.Id}
}

//
func userOffline(ctx context.Context, in interface{}) interface{} {
	raw := in.([]byte)
	log.Debugf("Control%#v", raw)
	user := ctx.(*UserContext)
	// 用户断开连接消息
	driver.UnlockUser(user.AgentId, user.UserId)
	return nil
}

// 为代理服务器分配连接ID
func allocAgentId(ctx context.Context, in interface{})interface{} {
	raw := in.([]byte)
	count := binary.BigEndian.Uint32(raw[HeadLen:])
	start := driver.NewSN("aid", int64(count))

	r := &GameFrame{Data: make([]byte, 8)}
	binary.BigEndian.PutUint64(r.Data, uint64(start))
	return r
}

// 查找房间地址
func findRoomAddr(ctx context.Context, in interface{})interface{} {
	raw := in.([]byte)
	count := binary.BigEndian.Uint32(raw[HeadLen:])
	start := driver.NewSN("aid", int64(count))

	r := &GameFrame{Data: make([]byte, 8)}
	binary.BigEndian.PutUint64(r.Data, uint64(start))
	return r
}