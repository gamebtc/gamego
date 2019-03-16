package internal

import (
	"google.golang.org/grpc"
	"io"
	. "local.com/abc/game/protocol"
	"net"
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

type NetStream struct {
	conn net.Conn
	head [HeadLen]byte
}

func (this *NetStream) Send(d []byte) error {
	if _, e := this.conn.Write(d); e == nil {
		return nil
	} else {
		return e
	}
}

func (this *NetStream) Recv() ([]byte, error) {
	head := this.head[:]
	n, err := io.ReadFull(this.conn, head)
	if err != nil || n != HeadLen {
		return nil, err
	}

	size := int(GetHeadLen(head))
	payload := make([]byte, HeadLen+size)

	if size > 0 {
		n, err = io.ReadFull(this.conn, payload[HeadLen:])
		if err != nil || n != size {
			return nil, err
		}
	}
	copy(payload[:HeadLen], head)
	return payload, nil
}

func (stream *NetStream) Close() {
	stream.conn.Close()
}
