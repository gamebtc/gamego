package internal

import (
	"context"
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

	s.RegistHandler(MsgId_Control, nil, Control)

	s.RegistMsg(MsgId_ErrorInfo, &ErrorInfo{})

	s.RegistHandler(MsgId_VerCheckReq, &VerCheckReq{}, VerCheck)
	s.RegistMsg(MsgId_VerCheckAck, &VerCheckAck{})

	s.RegistHandler(MsgId_UserLoginReq, &LoginReq{}, Login)
	s.RegistMsg(MsgId_UserLoginSuccessAck, &LoginSuccessAck{})
	s.RegistMsg(MsgId_UserLoginFailAck, &LoginFailAck{})

	s.RegistHandler(MsgId_HeartBeatReq, &HeartBeatReq{}, HeartBeat)
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
	if id >= int32(MsgId_UserMessageHeadSplit) {
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
		if elapse := time.Now().UnixNano() - start.UnixNano(); elapse > s.SlowOpNa {
			log.Warningf("elapse:%v(ms), id:%v, arg:%v", elapse/Ms2Na, id, arg)
		}
	}
	return
}

func Control(ctx context.Context, in interface{}) interface{} {
	raw := in.([]byte)
	log.Debugf("Control%#v", raw)
	user := ctx.(*UserContext)
	// 用户断开连接消息
	driver.UnlockUser(user.AgentId, user.UserId)
	return nil
}

// 版本检查
func VerCheck(ctx context.Context, in interface{}) interface{} {
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
						strIp = "223.104.18.236"
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

func HeartBeat(ctx context.Context, in interface{}) interface{} {
	arg := in.(*HeartBeatReq)
	user := ParseUserContext(ctx)
	log.Debugf("heartBeat:uid%v, id:%v", user.UserId, arg.Id)
	return &HeartBeatAck{Id: arg.Id}
}
