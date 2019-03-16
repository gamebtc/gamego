package protocol

import (
	"context"
	"encoding/binary"
	"google.golang.org/grpc/metadata"
)

const HeadLen = 6

func GetHead(b []byte) (int32, int32) {
	return int32(b[2]) | (int32(b[1]) << 8) | (int32(b[0]) << 16),
		int32(b[5]) | (int32(b[4]) << 8) | (int32(b[3]) << 16)
}

func GetHeadId(b []byte) int32 {
	return int32(b[5]) | (int32(b[4]) << 8) | (int32(b[3]) << 16)
}

func GetHeadLen(b []byte) int32 {
	return int32(b[2]) | (int32(b[1]) << 8) | (int32(b[0]) << 16)
}

func SetHead(b []byte, id int32) {
	v := len(b) - HeadLen
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
	b[3] = byte(id >> 16)
	b[4] = byte(id >> 8)
	b[5] = byte(id)
}

// 用户上下文
type UserContext struct {
	context.Context
	AgentId int64
	UserId  UserId
	Ip      IP
}

func ParseUserContext(ctx context.Context) *UserContext {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if args, ok := md["a-bin"]; ok && len(args) > 0 {
			if data := []byte(args[0]); len(data) >= 16 {
				return &UserContext{
					Context: ctx,
					AgentId: int64(binary.BigEndian.Uint64(data)),
					UserId:  int32(binary.BigEndian.Uint32(data[8:])),
					Ip:      int64(binary.BigEndian.Uint32(data[12:])),
				}
			}
		}
	}
	return &UserContext{
		Context: ctx,
	}
}

func NewServerHead(agent int64, uid int32, ip uint32) context.Context {
	md := metadata.MD{}
	callCtx := [16]byte{}
	binary.BigEndian.PutUint64(callCtx[0:8], uint64(agent))
	binary.BigEndian.PutUint32(callCtx[8:12], uint32(uid))
	binary.BigEndian.PutUint32(callCtx[12:16], uint32(ip))
	md.Set("a-bin", string(callCtx[:]))
	return metadata.NewOutgoingContext(context.Background(), md)
}

func GetUserHead(head []byte) (int64, int32, int64) {
	return int64(binary.BigEndian.Uint64(head[:])),
		int32(binary.BigEndian.Uint32(head[8:])),
		int64(binary.BigEndian.Uint32(head[12:]))
}

func NewUserHead(agent int64, uid int32, ip uint32) []byte {
	head := make([]byte, 16)
	binary.BigEndian.PutUint64(head[:], uint64(agent))
	binary.BigEndian.PutUint32(head[8:], uint32(uid))
	binary.BigEndian.PutUint32(head[12:], uint32(ip))
	return head
}

// 游戏流
type GameStream interface {
	Send([]byte) error
	Recv() ([]byte, error)
	Close()
}
