package msg

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	Ms2Na = 1000000
)

// server is used to implement GameService.
type ServerBase struct {
	HealthBase
	Coder    //编码解码器
	SlowOpNa int64
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (s *ServerBase) RegistHandler(msgId MsgId_Code, msg interface{}, msgHandler Handler) {
	id := int32(msgId)
	if msg != nil {
		t := reflect.TypeOf(msg)
		if s.SetHandler(t, id, msgHandler) == false {
			log.Fatalf("message %v is already registered", t)
		}
	} else {
		s.SetHandler(nil, id, msgHandler)
	}
}

func (s *ServerBase) RegistMsg(id MsgId_Code, msg interface{}) {
	t := reflect.TypeOf(msg)
	if _, ok := s.GetMsgId(t); ok {
		log.Fatalf("message %v is already registered", t)
	}
	s.SetMsgId(t, int32(id))
}

func (s *ServerBase) SetSlowOp(ms int64) {
	s.SlowOpNa = ms * Ms2Na
}

func (s *ServerBase) Send(Game_SendServer) error {
	return nil
}

func (s *ServerBase) Call(ctx context.Context, req *GameFrame) (res *GameFrame, err error) {
	// id
	id, arg, err := s.Decode(req.Data)
	if err != nil {
		return
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
	r := i.Handler(ctx, arg)
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

type HealthBase struct{}

var health = &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}

// Check 实现健康检查接口，直接返回健康状态.
func (*HealthBase) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return health, nil
}

func (*HealthBase) Watch(req *grpc_health_v1.HealthCheckRequest, x grpc_health_v1.Health_WatchServer) error {
	return x.Send(health)
}

func RegisterGrpcServer(s *grpc.Server) {
	reflection.Register(s)
	grpc_health_v1.RegisterHealthServer(s, &HealthBase{})
}
