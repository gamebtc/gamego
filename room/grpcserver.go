package room

import (
	"context"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"local.com/abc/game/msg"
	"local.com/abc/game/util"
)

// grpcServer is used to implement msg.GameService.
type grpcServer struct {
	msg.ServerBase
	connectCount int32
	maxConnect   int32
}

func (s *grpcServer) Call(ctx context.Context, req *msg.GameFrame) (res *msg.GameFrame, err error) {
	res = &msg.GameFrame{}
	return
}

// stream processing
// the center of game logic
func (s *grpcServer) Send(stream msg.Game_SendServer) (err error) {
	defer util.PrintPanicStack()
	defer atomic.AddInt32(&s.connectCount, -1)
	count := atomic.AddInt32(&s.connectCount, 1)
	if count >= s.maxConnect {
		return ErrorServiceBusy
	}

	user := msg.ParseUserContext(stream.Context())
	if user.AgentId == 0 || user.UserId == 0 {
		err = ErrorIncorrectFrameType
		return
	}

	// session init
	return s.newSession(&GrpcStream{ServerStream: stream}, user)
}

func (s *grpcServer) newSession(stream msg.GameStream, user *msg.UserContext) error {
	sess := &Session{
		AgentId:  user.AgentId,
		Ip:       user.Ip,
		Created:  time.Now(),
		sendChan: make(chan interface{}, 512),
		stopSend: make(chan struct{}),
	}

	// cleanup work
	defer func() {
		// 连接断开事件
		if sess.UserId != 0 {
			Call(func() { userOffline(sess) })
		}
		sess.Close()
		log.Infof("connection closed user:%v, Ip:%v", user.UserId, user.Ip)
	}()

	// 连接事件
	Call(func() { userOnline(sess, user.UserId) })
	sess.Start(stream)
	return nil
}

type GrpcStream struct {
	grpc.ServerStream
}

func (stream *GrpcStream) Send(d []byte) error {
	return stream.ServerStream.SendMsg(&msg.GameFrame{Data: d})
}

func (stream *GrpcStream) Recv() ([]byte, error) {
	ret := msg.GameFrame{}
	if err := stream.ServerStream.RecvMsg(&ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (stream *GrpcStream) Close() {
}
