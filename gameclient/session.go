package main

import (
	"io"
	"net"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"

	. "local.com/abc/game/protocol"
	"local.com/abc/game/util"
)

const (
	SESS_INIT     = 0   // 初始化
	SESS_VERCHECK = 1   // 已通过版本检查
	SESS_LOGINING = 2   // 登录中
	SESS_LOGINED  = 3   // 登录成功
	SESS_CLOSE    = 100 // 是否断开
)

// 配置
var (
	connectCount int32
	rpmLimit     int32
	sendChanLen  int
	recvChanLen  int
	readTimeout  time.Duration
	writeTimeout time.Duration
	slowOpNano   int64
)

type Session struct {
	Id     int64  // 连接唯一ID
	Ip     uint32 // IP地址
	Seed   uint32 // 接收加密种子
	Addr   string // 玩家IP地址
	Act    string // 玩家账号
	UserId int32  // 玩家ID
	Coder         // 编码解码

	roomId    int32            // 房间服ID
	Flag      int32            // 会话标记(0:初始化，1:已通过版本检查，2:登录中，3:登录成功, 4:已关闭)
	Created   time.Time        // TCP链接建立时间
	totalSend int64            // 总发送字节数
	totalRecv int64            // 总接收字节数
	conn      net.Conn         // 底层网络连接,连接到agent.
	recvChan  chan []byte      // 接收到的数据,1个生产者,1个消费者，发送值到已经关闭的channel会导致panic
	sendChan  chan interface{} // 等待发送的包,1个消费者，关闭已经关闭的channel会导致panic
	stopSend  chan struct{}    // 发送停止信号
	dieChan   chan struct{}    // 会话关闭信号
	dieOnce   int32            // 会话关闭保护
}

// packet recvice goroutine
func (this *Session) readClientLoop() {
	log.Debugf("start readClientLoop:%v", this.Id)
	defer log.Debugf("end readClientLoop:%v", this.Id)
	defer close(this.recvChan)
	h := [6]byte{}
	head := h[:]
	for this.Flag < SESS_CLOSE {
		n, err := io.ReadFull(this.conn, head)
		if err != nil {
			return
		}
		size := GetHeadLen(head)
		payload := make([]byte, size+HeadLen)
		n, err = io.ReadFull(this.conn, payload[HeadLen:])
		if err != nil || n != int(size) {
			return
		}
		this.totalRecv += int64(n + HeadLen)
		copy(payload[:HeadLen], head)

		//log.Debugf("recv3:%v", payload)
		// deliver the data to the input queue
		select {
		case this.recvChan <- payload:
		case <-this.dieChan:
			return
		}
	}
}

func (this *Session) send(data interface{}) bool {
	select {
	case <-this.stopSend:
		return false
	case this.sendChan <- data:
		return true
	}
}

// packet sending goroutine
func (this *Session) writeClientLoop() {
	log.Debugf("start writeClientLoop:%v", this.Id)
	defer log.Debugf("end writeClientLoop%v", this.Id)
	defer close(this.stopSend)
	var n int
	var err error
	for {
		select {
		case val, ok := <-this.sendChan:
			if ok == false || val == nil {
				return
			}

			var data []byte
			switch val := val.(type) {
			case []byte:
				data = val
			case error:
				if buf, e := this.Encode(errorInfo(0, 1001, "调用失败:"+val.Error())); e == nil {
					data = buf
				}
			default:
				if buf, e := this.Encode(val); e == nil {
					//log.Debugf("send val:%#v", val)
					data = buf
				} else {
					log.Debugf("send encode err:%v", e.Error())
				}
			}
			if data == nil {
				continue
			}

			if n, err = this.conn.Write(data); err != nil {
				switch err := err.(type) {
				case net.Error:
					if err.Timeout() == false {
						return
					}
				default:
					return
				}
			}
			this.totalSend += int64(n)
		case <-this.dieChan:
			return
		}
	}
}

func (this *Session) mainLoop() {
	log.Debugf("start mainloop:%v", this.Id)
	defer log.Debugf("end mainloop:%v", this.Id)
	defer util.PrintPanicStack()
	t := time.NewTicker(10 * time.Second)
	defer t.Stop()
	id := int32(0)
	for this.Flag < SESS_CLOSE {
		select {
		case data, ok := <-this.recvChan: // packet from network
			if ok == false {
				this.Flag = SESS_CLOSE
				return // 读关闭
			}
			if err := this.handle(data); err != nil {
				// 发送通用错误，然后关闭客户端
				id := GetHeadId(data)
				this.Flag = SESS_CLOSE
				log.Errorf("user:%v,msg:%v,addr:%v,err:%v", this.UserId, id, this.Addr, err.Error())
			}
		case <-t.C:
			if this.UserId > 0 {
				id++
				this.send(&HeartBeatReq{Id: id})
			} else {
				// 10秒还没登录成功的断开连接
				this.Flag = SESS_CLOSE
			}
		case <-this.dieChan: // this.Close
			this.Flag = SESS_CLOSE
		case <-this.stopSend: // 发送关闭
			this.Flag = SESS_CLOSE
		case <-signal.Die(): // 服务关闭
			this.Flag = SESS_CLOSE
		}
	}
}

//每个玩家4个线程，客户端网络读1个，客户端网络写1个,主线程1个，房间消息流接收1个
func (this *Session) Start() {
	go this.readClientLoop()
	go this.writeClientLoop()
	this.mainLoop()
}

// 线程安全关闭
func (this *Session) Close() {
	if atomic.CompareAndSwapInt32(&this.dieOnce, 0, 1) {
		close(this.dieChan)
	}
}

func (this *Session) closeRoom() {
	this.roomId = 0
}

func (this *Session) destroy() {
	//this.once.Do(func(){
	this.Close()
	this.closeRoom()
	this.Flag = SESS_CLOSE
	//})
}

// route client protocol
func (this *Session) handle(data []byte) error {
	id, msg, err := this.Coder.Decode(data)
	if err != nil {
		log.Debugf("uid1:%v,msg:%v,addr:%v,err:%v", this.UserId, id, this.Addr, err.Error())
		return nil
	}

	if h := handlers[id]; h != nil {
		h(this, msg)
	} else {
		log.Debugf("uid2:%v,msg:%v,addr:%v,err:%v", this.UserId, id, this.Addr, ErrorUndefined.Error())
	}
	return nil
}

func errorInfo(id int32, code int32, m string) *ErrorInfo {
	return &ErrorInfo{ReqId: id, Code: code, Msg: m}
}
