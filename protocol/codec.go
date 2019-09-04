package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"

	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack/v4"
)

type Handler func(context.Context, interface{}) interface{}

type HandlerInfo struct {
	reflect.Type
	Handler
}

type Coder interface {
	//google.golang.org/grpc/encoding.Codec
	//Marshal(v interface{}) ([]byte, error)
	Unmarshal(buf []byte, v interface{}) (err error)

	Encode(v interface{}) ([]byte, error)
	Decode(buf []byte) (id int32, v interface{}, err error)
	Name() string

	GetMsgId(t reflect.Type) (id int32, ok bool)
	SetMsgId(t reflect.Type, id int32) (ok bool)
	SetHandler(t reflect.Type, id int32, h Handler) (ok bool)
	GetHandler(id int32) (info *HandlerInfo)
}

type IdCoder struct {
	msgID   map[reflect.Type]int32           //通过类型查找消息号
	msgInfo [math.MaxUint16 + 1]*HandlerInfo //通过消息号查找类型和回调
}

func (c IdCoder) GetMsgId(t reflect.Type) (id int32, ok bool) {
	if c.msgID != nil {
		id, ok = c.msgID[t]
	}
	return
}

func (c *IdCoder) SetMsgId(t reflect.Type, id int32) bool {
	return c.SetHandler(t, id, nil)
}

func (c *IdCoder) SetHandler(t reflect.Type, id int32, h Handler) bool {
	if t != nil {
		if c.msgID == nil {
			c.msgID = make(map[reflect.Type]int32, 100)
		} else if _, ok := c.msgID[t]; ok {
			return false
		}
		c.msgID[t] = id
	}
	c.msgInfo[id] = &HandlerInfo{
		Type:    t,
		Handler: h,
	}
	return true
}

func (c *IdCoder) GetHandler(id int32) *HandlerInfo {
	return c.msgInfo[id]
}

type ProtoMarshal interface {
	Size() int
	MarshalTo(dAtA []byte) (int, error)
}

type protoCoder struct {
	IdCoder
}

func (c *protoCoder) Name() string {
	return "proto"
}

func (c *protoCoder) Encode(v interface{}) (buf []byte, err error) {
	t := reflect.TypeOf(v)
	if id, ok := c.GetMsgId(t); ok {
		switch m := v.(type) {
		case ProtoMarshal:
			size := m.Size()
			buf = make([]byte, size+HeadLen)
			_, err = m.MarshalTo(buf[HeadLen:])
			if err != nil {
				return nil, err
			}
			SetHead(buf, id)
			return
		case proto.Message:
			var b []byte
			if b, err = proto.Marshal(m); err != nil {
				return
			}
			buf = make([]byte, len(b)+HeadLen)
			SetHead(buf, id)
			copy(buf[HeadLen:], b)
			return
		case []byte:
			buf = m
			return
		default:
			err = fmt.Errorf("message %v not registered", t)
		}
	} else {
		err = fmt.Errorf("message %v not registered", t)
	}
	return
}

func (c *protoCoder) Decode(buf []byte) (id int32, v interface{}, err error) {
	if len(buf) < HeadLen {
		err = ErrorPacketLen
		return
	}
	id = GetHeadId(buf)
	if i := c.GetHandler(id); i != nil && i.Type != nil {
		v = reflect.New(i.Type.Elem()).Interface()
		err = c.Unmarshal(buf[HeadLen:], v)
	} else {
		v = buf
	}
	return
}

func (c *protoCoder) Unmarshal(buf []byte, v interface{}) (err error) {
	return proto.UnmarshalMerge(buf, v.(proto.Message))
}

type jsonCoder struct {
	IdCoder
}

func (c *jsonCoder) Name() string {
	return "json"
}

func (c *jsonCoder) Encode(v interface{}) (buf []byte, err error) {
	t := reflect.TypeOf(v)
	if id, ok := c.GetMsgId(t); ok {
		var b []byte
		if b, err = json.Marshal(v); err != nil {
			return
		}
		buf = make([]byte, len(b)+HeadLen)
		SetHead(buf, id)
		copy(buf[HeadLen:], b)
	} else {
		err = fmt.Errorf("message %v not registered", t)
	}
	return
}

func (c *jsonCoder) Decode(buf []byte) (id int32, v interface{}, err error) {
	if len(buf) < HeadLen {
		err = ErrorPacketLen
		return
	}
	id = GetHeadId(buf)
	if i := c.GetHandler(id); i != nil && i.Type != nil {
		v = reflect.New(i.Type.Elem()).Interface()
		err = c.Unmarshal(buf[HeadLen:], v)
	} else {
		v = buf
	}
	return
}

func (c *jsonCoder) Unmarshal(buf []byte, v interface{}) (err error) {
	return json.Unmarshal(buf, v)
}

type jsoniterCoder struct {
	IdCoder
	jsoniter.API
}

func (c *jsoniterCoder) Name() string {
	return "jsoniter"
}

func (c *jsoniterCoder) Encode(v interface{}) (buf []byte, err error) {
	t := reflect.TypeOf(v)
	if id, ok := c.GetMsgId(t); ok {
		var b []byte
		if b, err = c.API.Marshal(v); err != nil {
			return
		}
		buf = make([]byte, len(b)+HeadLen)
		SetHead(buf, id)
		copy(buf[HeadLen:], b)
	} else {
		err = fmt.Errorf("message %v not registered", t)
	}
	return
}

func (c *jsoniterCoder) Decode(buf []byte) (id int32, v interface{}, err error) {
	if len(buf) < HeadLen {
		err = ErrorPacketLen
		return
	}
	id = GetHeadId(buf)
	if i := c.GetHandler(id); i != nil && i.Type != nil {
		v = reflect.New(i.Type.Elem()).Interface()
		err = c.Unmarshal(buf[HeadLen:], v)
	} else {
		v = buf
	}
	return
}

func (c *jsoniterCoder) Unmarshal(buf []byte, v interface{}) (err error) {
	return c.API.Unmarshal(buf, v)
}

type msgpackCoder struct {
	IdCoder
}

func (c *msgpackCoder) Name() string {
	return "msgpack"
}

func (c *msgpackCoder) Encode(v interface{}) (buf []byte, err error) {
	t := reflect.TypeOf(v)
	if id, ok := c.GetMsgId(t); ok {
		var b []byte
		if b, err = c.Marshal(v); err != nil {
			return
		}
		buf = make([]byte, len(b)+HeadLen)
		SetHead(buf, id)
		copy(buf[HeadLen:], b)
	} else {
		err = fmt.Errorf("message %v not registered", t)
	}
	return
}

func (c *msgpackCoder) Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	e := msgpack.NewEncoder(&buf)
	e.UseJSONTag(true)
	err := e.Encode(v)
	return buf.Bytes(), err
}

func (c *msgpackCoder) Decode(buf []byte) (id int32, v interface{}, err error) {
	if len(buf) < HeadLen {
		err = ErrorPacketLen
		return
	}
	id = GetHeadId(buf)
	if i := c.GetHandler(id); i != nil && i.Type != nil {
		v = reflect.New(i.Type.Elem()).Interface()
		err = c.Unmarshal(buf[HeadLen:], v)
	} else {
		v = buf
	}
	return
}

func (c *msgpackCoder) Unmarshal(buf []byte, v interface{}) (err error) {
	d := msgpack.NewDecoder(bytes.NewReader(buf))
	d.UseJSONTag(true)
	return d.Decode(v)
}

type MsgMarshal interface {
	Msgsize() (s int)
	MarshalMsg([]byte) ([]byte, error)
}

type msgCoder struct {
	msgpackCoder
}

func (c *msgCoder) Name() string {
	return "msg"
}

func (c *msgCoder) Encode(v interface{}) (buf []byte, err error) {
	t := reflect.TypeOf(v)
	if id, ok := c.GetMsgId(t); ok {
		var b []byte
		switch m := v.(type) {
		case MsgMarshal:
			size := m.Msgsize()
			buf = make([]byte, size+HeadLen)
			b, err = m.MarshalMsg(buf[HeadLen:HeadLen])
			if err != nil {
				return nil, err
			}
			buf = buf[:len(b)+HeadLen]
			SetHead(buf, id)
			return
		case msgp.Marshaler:
			b, err = m.MarshalMsg(b)
		default:
			b, err = c.Marshal(v)
		}
		if err != nil {
			return
		}
		buf = make([]byte, len(b)+HeadLen)
		SetHead(buf, id)
		copy(buf[HeadLen:], b)
	} else {
		err = fmt.Errorf("message %v not registered", t)
	}
	return
}

func (c *msgCoder) Decode(buf []byte) (id int32, v interface{}, err error) {
	if len(buf) < HeadLen {
		err = ErrorPacketLen
		return
	}
	id = GetHeadId(buf)
	if i := c.GetHandler(id); i != nil && i.Type != nil {
		v = reflect.New(i.Type.Elem()).Interface()
		err = c.Unmarshal(buf[HeadLen:], v)
	} else {
		v = buf
	}
	return
}

func (c *msgCoder) Unmarshal(buf []byte, v interface{}) (err error) {
	if mg, ok := v.(msgp.Unmarshaler); ok {
		_, err = mg.UnmarshalMsg(buf)
	} else {
		err = c.msgpackCoder.Unmarshal(buf, v)
	}
	return
}

var (
	protoCode    = new(protoCoder)
	jsonCode     = new(jsonCoder)
	jsoniterCode = newJsoniterCoder()
	msgpackCode  = new(msgpackCoder)
	msgCode      = new(msgCoder)
)

func newJsoniterCoder() *jsoniterCoder {
	j := new(jsoniterCoder)
	//j.API = jsoniter.ConfigCompatibleWithStandardLibrary
	j.API = jsoniter.ConfigFastest
	return j
}

func GetCoder(name string) Coder {
	switch name {
	case "json":
		return jsonCode
	case "proto":
		return protoCode
	case "msgpack":
		return msgpackCode
	case "msg":
		return msgCode
	case "jsoniter":
		return jsoniterCode
	}
	return jsonCode
}
