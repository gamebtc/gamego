package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/vmihailenco/msgpack"
)

// go test -v -run="指定函数名"
// https://blog.csdn.net/hjmnasdkl/article/details/81304329
// go test -v -run="none" -bench=.    不允许单元测试，运行所有的基准测试
// -benchmem 表示分配内存的次数和字节数，-benchtime="3s" 表示持续3秒

var (
	jsonTest     = new(jsonCoder)
	jsoniterTest = newJsoniterCoder()
	protoTest    = new(protoCoder)
	msgpackTest  = new(msgpackCoder)
	msgTest      = new(msgCoder)
)

var req = new(LoginReq)

func init() {

	x := "{\"type\":0,\"userId\":0,\"dev\":{\"id\":\"7ce66f652d0a4656aec512af691fa724\",\"vend\":\"apple\",\"name\":\"iphoneX\",\"mac\":\"99-33-fd-34-34\",\"osLang\":\"zh-cn\",\"osVer\":\"sx4.34.34\",\"other\":\"nukoowe\",\"imei\":\"0f5791616e574e62\",\"emid\":\"4e99eee10891449789a58a82d1c611c4\",\"sn\":\"94c2d3c23a6b42c88a9cc9d965352a2a\"},\"env\":{\"id\":1,\"pack\":1101,\"ver\":\"1.0.0\",\"chan\":\"bb_ios_1\",\"refer\":\"100909\",\"other\":\"otherapplication\"},\"name\":\"D2F92919\",\"pwd\":\"58c5ca9d017e488887ad5214756c29be\",\"udid\":\"173648c2090b440b83269f67fd6434d0\",\"time\":1543997640}"
	json.Unmarshal([]byte(x), req)

	fmt.Printf("%#v\r\n,dev%#v,env:%#v", req, req.Dev, req.Env)

	tp := reflect.TypeOf(req)
	id := int32(199)
	jsonTest.SetHandler(tp, id, nil)
	jsoniterTest.SetHandler(tp, id, nil)
	protoTest.SetHandler(tp, id, nil)
	msgpackTest.SetHandler(tp, id, nil)
	msgTest.SetHandler(tp, id, nil)

	id = int32(23)
	tp = reflect.TypeOf(&Handshake{})
	jsonTest.SetHandler(tp, id, nil)
	jsoniterTest.SetHandler(tp, id, nil)
	protoTest.SetHandler(tp, id, nil)
	msgpackTest.SetHandler(tp, id, nil)
	msgTest.SetHandler(tp, id, nil)
}

func BenchmarkProto(b *testing.B) {
	var outMsg interface{}
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = protoTest.Encode(req)
		_, outMsg, _ = protoTest.Decode(buf)
	}
	x := outMsg.(*LoginReq)
	if x.Dev.Id != req.Dev.Id {
		b.Fatalf("[%v]", x)
	}
	//fmt.Printf("Proto:%v\r\n", len(buf))
}

func BenchmarkJson(b *testing.B) {
	var outMsg interface{}
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = jsonTest.Encode(req)
		_, outMsg, _ = jsonTest.Decode(buf)
	}
	x := outMsg.(*LoginReq)
	if x.Dev.Id != req.Dev.Id {
		b.Fatalf("[%v]", x)
	}
	//fmt.Printf("Json:%v\r\n", len(buf))
}

func BenchmarkJsoniter(b *testing.B) {
	var outMsg interface{}
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = jsoniterTest.Encode(req)
		_, outMsg, _ = jsoniterTest.Decode(buf)
	}
	x := outMsg.(*LoginReq)
	if x.Dev.Id != req.Dev.Id {
		b.Fatalf("[%v]", x)
	}
	//fmt.Printf("Jsoniter:%v\r\n", len(buf))
}

func BenchmarkMsgpack(b *testing.B) {
	var outMsg interface{}
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = msgpackTest.Encode(req)
		_, outMsg, _ = msgpackTest.Decode(buf)
	}
	x := outMsg.(*LoginReq)
	if x.Dev.Id != req.Dev.Id {
		b.Fatalf("[%v]", x)
	}
	//fmt.Printf("Msgpack:%v\r\n", len(buf))
}

func BenchmarkMsg(b *testing.B) {
	var outMsg interface{}
	var buf []byte
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf, _ = msgTest.Encode(req)
		_, outMsg, _ = msgTest.Decode(buf)
	}
	x := outMsg.(*LoginReq)
	if x.Dev.Id != req.Dev.Id {
		b.Fatalf("[%v]", x)
	}
	//fmt.Printf("Msg:%v\r\n", len(buf))
}

func TestMsg(t *testing.T) {
	buf := []byte{0XDF, 0X00, 0X00, 0X00, 0X08, 0XA4, 0X74, 0X79, 0X70, 0X65, 0XCD, 0X04, 0X4D, 0XA6, 0X75, 0X73, 0X65, 0X72, 0X49, 0X64, 0X00, 0XA3, 0X64, 0X65, 0X76, 0XDF, 0X00, 0X00, 0X00, 0X0A, 0XA2, 0X69, 0X64, 0XD9, 0X20, 0X66, 0X38, 0X37, 0X36, 0X38, 0X31, 0X65, 0X62, 0X63, 0X63, 0X66, 0X62, 0X34, 0X63, 0X39, 0X36, 0X61, 0X30, 0X61, 0X39, 0X37, 0X63, 0X61, 0X35, 0X66, 0X37, 0X35, 0X32, 0X63, 0X37, 0X38, 0X65, 0XA4, 0X76, 0X65, 0X6E, 0X64, 0XA5, 0X61, 0X70, 0X70, 0X6C, 0X65, 0XA4, 0X6E, 0X61, 0X6D, 0X65, 0XA7, 0X69, 0X70, 0X68, 0X6F, 0X6E, 0X65, 0X58, 0XA3, 0X6D, 0X61, 0X63, 0XAE, 0X39, 0X39, 0X2D, 0X33, 0X33, 0X2D, 0X66, 0X64, 0X2D, 0X33, 0X34, 0X2D, 0X33, 0X34, 0XA6, 0X6F, 0X73, 0X4C, 0X61, 0X6E, 0X67, 0XA5, 0X7A, 0X68, 0X2D, 0X63, 0X6E, 0XA5, 0X6F, 0X73, 0X56, 0X65, 0X72, 0XA9, 0X73, 0X78, 0X34, 0X2E, 0X33, 0X34, 0X2E, 0X33, 0X34, 0XA5, 0X6F, 0X74, 0X68, 0X65, 0X72, 0XA7, 0X6E, 0X75, 0X6B, 0X6F, 0X6F, 0X77, 0X65, 0XA4, 0X69, 0X6D, 0X65, 0X69, 0XB0, 0X66, 0X36, 0X32, 0X38, 0X64, 0X30, 0X65, 0X65, 0X62, 0X37, 0X34, 0X62, 0X34, 0X63, 0X36, 0X61, 0XA4, 0X65, 0X6D, 0X69, 0X64, 0XD9, 0X20, 0X65, 0X61, 0X32, 0X63, 0X65, 0X36, 0X39, 0X66, 0X31, 0X33, 0X38, 0X66, 0X34, 0X39, 0X32, 0X66, 0X38, 0X34, 0X37, 0X65, 0X65, 0X66, 0X63, 0X63, 0X61, 0X33, 0X65, 0X61, 0X63, 0X31, 0X39, 0X31, 0XA2, 0X73, 0X6E, 0XD9, 0X20, 0X33, 0X33, 0X37, 0X31, 0X66, 0X62, 0X38, 0X34, 0X38, 0X31, 0X65, 0X36, 0X34, 0X63, 0X37, 0X37, 0X39, 0X61, 0X31, 0X34, 0X33, 0X66, 0X34, 0X36, 0X36, 0X38, 0X32, 0X64, 0X34, 0X36, 0X61, 0X32, 0XA3, 0X65, 0X6E, 0X76, 0XDF, 0X00, 0X00, 0X00, 0X06, 0XA2, 0X69, 0X64, 0X01, 0XA4, 0X70, 0X61, 0X63, 0X6B, 0XCD, 0X04, 0X4D, 0XA3, 0X76, 0X65, 0X72, 0XA5, 0X31, 0X2E, 0X30, 0X2E, 0X30, 0XA7, 0X63, 0X68, 0X61, 0X6E, 0X6E, 0X65, 0X6C, 0XA8, 0X62, 0X62, 0X5F, 0X69, 0X6F, 0X73, 0X5F, 0X31, 0XA5, 0X72, 0X65, 0X66, 0X65, 0X72, 0XA6, 0X31, 0X30, 0X30, 0X39, 0X30, 0X39, 0XA5, 0X6F, 0X74, 0X68, 0X65, 0X72, 0XB0, 0X6F, 0X74, 0X68, 0X65, 0X72, 0X61, 0X70, 0X70, 0X6C, 0X69, 0X63, 0X61, 0X74, 0X69, 0X6F, 0X6E, 0XA4, 0X6E, 0X61, 0X6D, 0X65, 0XA8, 0X36, 0X34, 0X33, 0X38, 0X37, 0X32, 0X33, 0X31, 0XA3, 0X70, 0X77, 0X64, 0XD9, 0X20, 0X30, 0X61, 0X32, 0X37, 0X33, 0X39, 0X66, 0X34, 0X65, 0X37, 0X36, 0X34, 0X34, 0X32, 0X31, 0X62, 0X39, 0X65, 0X63, 0X66, 0X32, 0X66, 0X31, 0X63, 0X38, 0X35, 0X37, 0X62, 0X65, 0X63, 0X66, 0X63, 0XA7, 0X6D, 0X61, 0X63, 0X68, 0X69, 0X6E, 0X65, 0XD9, 0X20, 0X65, 0X31, 0X34, 0X66, 0X39, 0X61, 0X62, 0X64, 0X61, 0X64, 0X36, 0X36, 0X34, 0X30, 0X38, 0X61, 0X38, 0X37, 0X65, 0X64, 0X36, 0X36, 0X34, 0X34, 0X34, 0X39, 0X62, 0X61, 0X34, 0X39, 0X34, 0X66, 0XA4, 0X74, 0X69, 0X6D, 0X65, 0XCE, 0X5C, 0X07, 0XE7, 0X30}

	rqx := new(LoginReq)
	bts, err := rqx.UnmarshalMsg(buf)
	if err != nil {
		t.Fatalf("[%v,%#v]", err, bts)
	}
}

func TestMsgpack(t *testing.T) {
	buf := []byte{0XDF, 0X00, 0X00, 0X00, 0X08, 0XA4, 0X74, 0X79, 0X70, 0X65, 0XCD, 0X04, 0X4D, 0XA6, 0X75, 0X73, 0X65, 0X72, 0X49, 0X64, 0X00, 0XA3, 0X64, 0X65, 0X76, 0XDF, 0X00, 0X00, 0X00, 0X0A, 0XA2, 0X69, 0X64, 0XD9, 0X20, 0X66, 0X38, 0X37, 0X36, 0X38, 0X31, 0X65, 0X62, 0X63, 0X63, 0X66, 0X62, 0X34, 0X63, 0X39, 0X36, 0X61, 0X30, 0X61, 0X39, 0X37, 0X63, 0X61, 0X35, 0X66, 0X37, 0X35, 0X32, 0X63, 0X37, 0X38, 0X65, 0XA4, 0X76, 0X65, 0X6E, 0X64, 0XA5, 0X61, 0X70, 0X70, 0X6C, 0X65, 0XA4, 0X6E, 0X61, 0X6D, 0X65, 0XA7, 0X69, 0X70, 0X68, 0X6F, 0X6E, 0X65, 0X58, 0XA3, 0X6D, 0X61, 0X63, 0XAE, 0X39, 0X39, 0X2D, 0X33, 0X33, 0X2D, 0X66, 0X64, 0X2D, 0X33, 0X34, 0X2D, 0X33, 0X34, 0XA6, 0X6F, 0X73, 0X4C, 0X61, 0X6E, 0X67, 0XA5, 0X7A, 0X68, 0X2D, 0X63, 0X6E, 0XA5, 0X6F, 0X73, 0X56, 0X65, 0X72, 0XA9, 0X73, 0X78, 0X34, 0X2E, 0X33, 0X34, 0X2E, 0X33, 0X34, 0XA5, 0X6F, 0X74, 0X68, 0X65, 0X72, 0XA7, 0X6E, 0X75, 0X6B, 0X6F, 0X6F, 0X77, 0X65, 0XA4, 0X69, 0X6D, 0X65, 0X69, 0XB0, 0X66, 0X36, 0X32, 0X38, 0X64, 0X30, 0X65, 0X65, 0X62, 0X37, 0X34, 0X62, 0X34, 0X63, 0X36, 0X61, 0XA4, 0X65, 0X6D, 0X69, 0X64, 0XD9, 0X20, 0X65, 0X61, 0X32, 0X63, 0X65, 0X36, 0X39, 0X66, 0X31, 0X33, 0X38, 0X66, 0X34, 0X39, 0X32, 0X66, 0X38, 0X34, 0X37, 0X65, 0X65, 0X66, 0X63, 0X63, 0X61, 0X33, 0X65, 0X61, 0X63, 0X31, 0X39, 0X31, 0XA2, 0X73, 0X6E, 0XD9, 0X20, 0X33, 0X33, 0X37, 0X31, 0X66, 0X62, 0X38, 0X34, 0X38, 0X31, 0X65, 0X36, 0X34, 0X63, 0X37, 0X37, 0X39, 0X61, 0X31, 0X34, 0X33, 0X66, 0X34, 0X36, 0X36, 0X38, 0X32, 0X64, 0X34, 0X36, 0X61, 0X32, 0XA3, 0X65, 0X6E, 0X76, 0XDF, 0X00, 0X00, 0X00, 0X06, 0XA2, 0X69, 0X64, 0X01, 0XA4, 0X70, 0X61, 0X63, 0X6B, 0XCD, 0X04, 0X4D, 0XA3, 0X76, 0X65, 0X72, 0XA5, 0X31, 0X2E, 0X30, 0X2E, 0X30, 0XA7, 0X63, 0X68, 0X61, 0X6E, 0X6E, 0X65, 0X6C, 0XA8, 0X62, 0X62, 0X5F, 0X69, 0X6F, 0X73, 0X5F, 0X31, 0XA5, 0X72, 0X65, 0X66, 0X65, 0X72, 0XA6, 0X31, 0X30, 0X30, 0X39, 0X30, 0X39, 0XA5, 0X6F, 0X74, 0X68, 0X65, 0X72, 0XB0, 0X6F, 0X74, 0X68, 0X65, 0X72, 0X61, 0X70, 0X70, 0X6C, 0X69, 0X63, 0X61, 0X74, 0X69, 0X6F, 0X6E, 0XA4, 0X6E, 0X61, 0X6D, 0X65, 0XA8, 0X36, 0X34, 0X33, 0X38, 0X37, 0X32, 0X33, 0X31, 0XA3, 0X70, 0X77, 0X64, 0XD9, 0X20, 0X30, 0X61, 0X32, 0X37, 0X33, 0X39, 0X66, 0X34, 0X65, 0X37, 0X36, 0X34, 0X34, 0X32, 0X31, 0X62, 0X39, 0X65, 0X63, 0X66, 0X32, 0X66, 0X31, 0X63, 0X38, 0X35, 0X37, 0X62, 0X65, 0X63, 0X66, 0X63, 0XA7, 0X6D, 0X61, 0X63, 0X68, 0X69, 0X6E, 0X65, 0XD9, 0X20, 0X65, 0X31, 0X34, 0X66, 0X39, 0X61, 0X62, 0X64, 0X61, 0X64, 0X36, 0X36, 0X34, 0X30, 0X38, 0X61, 0X38, 0X37, 0X65, 0X64, 0X36, 0X36, 0X34, 0X34, 0X34, 0X39, 0X62, 0X61, 0X34, 0X39, 0X34, 0X66, 0XA4, 0X74, 0X69, 0X6D, 0X65, 0XCE, 0X5C, 0X07, 0XE7, 0X30}

	rqx := new(LoginReq)
	d := msgpack.NewDecoder(bytes.NewReader(buf))
	d.UseJSONTag(true)
	err := d.Decode(rqx)
	if err != nil {
		t.Fatalf("[err%v]", err)
	}
}

func TestEncode(t *testing.T) {
	d := &Handshake{
		Code: 2334,
		Seed: 23445,
		Msg:  "中国",
		Ip:   []string{"127.0.0.1"},
	}

	r1, _ := msgpackTest.Encode(d)
	r2, _ := msgTest.Encode(d)

	l := len(r1)
	if l != len(r2) {
		t.Error("测试失败A1")
		t.Logf("r1:%v", r1)
		t.Logf("r2:%v", r2)
		return
	}

	for i := 0; i < l; i++ {
		if r1[i] != r2[i] {
			t.Logf("r1:%v", r1)
			t.Logf("r2:%v", r2)
			t.Error("测试失败A2")
			break
		}
	}
}

func TestDecode(t *testing.T) {
	d := &Handshake{
		Code: 2334,
		Seed: 23445,
		Msg:  "中国",
		Ip:   []string{"127.0.0.1"},
	}

	r1, _ := msgpackTest.Encode(d)
	r2, _ := msgTest.Encode(d)

	t.Logf("r1:%v", r1)
	t.Logf("r2:%v", r2)

	id1, msg1, err1 := msgpackTest.Decode(r2)
	id2, msg2, err2 := msgTest.Decode(r1)

	if id1 != id2 || err1 != nil || err2 != nil {
		t.Error("测试失败1")
	}

	t.Logf("msg1:%v", msg1)
	t.Logf("msg2:%v", msg2)

	m1 := msg1.(*Handshake)
	m2 := msg2.(*Handshake)

	if len(m1.Ip) != len(m2.Ip) {
		t.Error("测试失败2")
	}
	if m1.Ip[0] != m2.Ip[0] {
		t.Error("测试失败3")
	}
}

func TestChange(t *testing.T) {
	buf := []byte{1, 5, 10}
	x := 2
	y := 1
	z := 1

	t1 := buf[z]
	t2 := buf[x]
	t3 := buf[y]

	buf[x] = t1
	buf[y] = t2
	buf[z] = t3
	//buf[x], buf[y], buf[z] = buf[z], buf[x], buf[y]

	if buf[x] != 5 || buf[y] != 1 || buf[z] != 10 {
		t.Errorf("[%v]", buf)
	}

	buf1 := NewPoker(8, false, false)
	buf2 := NewPoker(8, false, true)

	if len(buf1) != len(buf2) {
		t.Error("测试失败3")
	}

	fmt.Printf("\r\nbuf1:%v\r\n", buf1)
	fmt.Printf("\r\nbuf2:%v\r\n", buf2)

	buf3 := NewPoker(1, true, false)
	buf4 := NewPoker(1, true, true)
	buf5 := NewPoker(1, true, true)
	buf6 := NewPoker(1, true, true)
	buf7 := NewPoker(1, true, true)
	buf8 := NewPoker(1, true, true)
	buf9 := NewPoker(1, true, true)
	fmt.Printf("\r\nbuf3:%v", PokerArrayString(buf3, "|"))
	fmt.Printf("\r\nbuf4:%v", PokerArrayString(buf4, "|"))
	fmt.Printf("\r\nbuf5:%v", PokerArrayString(buf5, "|"))
	fmt.Printf("\r\nbuf6:%v", PokerArrayString(buf6, "|"))
	fmt.Printf("\r\nbuf7:%v", PokerArrayString(buf7, "|"))
	fmt.Printf("\r\nbuf8:%v", PokerArrayString(buf8, "|"))
	fmt.Printf("\r\nbuf9:%v", PokerArrayString(buf9, "|"))
}
