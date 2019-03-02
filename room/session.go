package room

import (
	"io"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	. "local.com/abc/game/msg"
	"local.com/abc/game/util"
)

const (
	SESS_INIT     = 0   // 初始化
	SESS_VERCHECK = 1   // 已通过版本检查
	SESS_LOGINING = 2   // 登录中
	SESS_LOGINED  = 3   // 登录成功
	SESS_CLOSE    = 100 // 是否断开
)

// 会话是一个单独玩家的上下文，在连入后到退出前的整个生命周期内存在
type Session struct {
	AgentId  int64            // 连接唯一ID
	Ip       model.IP         // IP地址
	UserId   int32            // 玩家ID
	Online   bool             // 是否在线(只有在线才处理接收到的消息)
	Reline   bool             // 是否是重新上线
	Playing  bool             // 是否在游戏中
	Role     interface{}      // 游戏角色数据
	Created  time.Time        // TCP链接建立时间
	Flag     int32            // 会话标记(0:初始化，1:已通过版本检查，2:登录中，3:登录成功, 4:已关闭)
	sendChan chan interface{} // 发送出去的数据
	stopSend chan struct{}    // 发送停止信号
	dieChan  chan struct{}    // 会话关闭信号
	dieOnce  int32            // 会话关闭信号
}

// 不安全的发送，被发送的消息在另外一个线程编码，使用此方法，需保证被发送的对象线程安全
func (this *Session) UnsafeSend(data interface{}) bool {
	if this.stopSend != nil {
		select {
		case <-this.stopSend:
			return false
		case this.sendChan <- data:
			return true
			//default:
			//	log.WithFields(log.Fields{"user": this.UserId, "ip": this.Ip}).Warning("pending full")
			//	return false
		}
	}
	return true
}

func (this *Session) Send(val interface{}) bool {
	if this.stopSend != nil {
		if data, err := coder.Encode(val); err == nil {
			return this.UnsafeSend(data)
		}
		return false
	}
	return true
}

func (this *Session) SendError(id int32, code int32, m string, k string) {
	this.UnsafeSend(&ErrorInfo{ReqId: id, Code: code, Msg: m, Key: k})
}

func (this *Session) Start(stream GameStream) {
	if stream != nil && this.stopSend != nil {
		go this.sendLoop(stream)
		this.recvLoop(stream)
	}
}

// 写消息循环
func (this *Session) sendLoop(stream GameStream) {
	defer util.PrintPanicStack()
	defer close(this.stopSend)
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
			default:
				if buf, err := coder.Encode(val); err == nil {
					data = buf
				}
			}
			if data != nil {
				if err := stream.Send(data); err != nil {
					if err != io.EOF {
						log.Error(err)
					}
					return
				}
			}
		case <-this.dieChan:
			return
		}
	}
}

func (this *Session) recvLoop(stream GameStream) {
	defer util.PrintPanicStack()
	for this.Flag < SESS_CLOSE {
		if data, err := stream.Recv(); err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			return
		} else {
			if len(data) < HeadLen {
				return
			}
			id, arg, e := coder.Decode(data)
			log.Debugf("coder:%v,%v,%v,%v", coder.Name(), id, arg, e)
			if e != nil || this.Call(id, arg) == false {
				return
			}
		}
	}
}

func (this *Session) Call(id int32, arg interface{}) bool {
	select {
	case <-signal.Die():
		this.Flag = SESS_CLOSE
	// 所有用户消息，放入主循环处理
	case messageChan <- &NetMessage{Session: this, Id: id, Arg: arg}:
		return true
	}
	return false
}

func (this *Session) Close() {
	if atomic.CompareAndSwapInt32(&this.dieOnce, 0, 1) {
		close(this.dieChan)
		// 通知服务器退出
	}
}
