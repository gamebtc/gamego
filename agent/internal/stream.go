package internal

import (
	"io"
	"net"
	"sync/atomic"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	. "local.com/abc/game/protocol"
)

type GrpcStream struct {
	grpc.ClientStream
}

func (stream *GrpcStream) Send(d []byte) error {
	return stream.ClientStream.SendMsg(&GameFrame{Data: d})
}

func (stream *GrpcStream) Recv() ([]byte, error) {
	ret := GameFrame{}
	if err := stream.ClientStream.RecvMsg(&ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}

func (stream *GrpcStream) Close() {
	stream.ClientStream.CloseSend()
}

const NetBuffSize = 1024*1024
type NetStream struct {
	conn     net.Conn
	recvBuf  [HeadLen+NetBuffSize]byte // 接收缓冲
	disposed uint32                    // 会话关闭标记
}

func (stream * NetStream) Send(d []byte) error {
	if _, e := stream.conn.Write(d); e == nil {
		return nil
	} else {
		return e
	}
}

func (stream *NetStream) Recv() ([]byte, error) {
	head := stream.recvBuf[:HeadLen]
	n, err := io.ReadFull(stream.conn, head)
	if err != nil || n != HeadLen {
		return nil, err
	}
	size := int(GetHeadLen(head))
	var payload []byte
	if size > NetBuffSize {
		payload = make([]byte, HeadLen+size)
		copy(payload[:HeadLen], head)
	}else{
		payload = stream.recvBuf[:HeadLen+size]
	}
	if size > 0 {
		n, err = io.ReadFull(stream.conn, payload[HeadLen:])
		if err != nil || n != size {
			return nil, err
		}
	}
	return payload, nil
}

func (stream *NetStream) Close() {
	if atomic.CompareAndSwapUint32(&stream.disposed, 0, 1) {
		stream.conn.Close()
		log.Debugf("close net stream:%v", stream.conn.LocalAddr())
	}
}
