package room

import (
	"io"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
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
	AgentId    int64       // 连接唯一ID
	UserId     int32       // 玩家ID
	Ip         model.IP    // IP地址
	Online     bool        // 是否在线(只有在线才处理接收到的消息)
	Disposed   bool        // 是否已销毁(被强制退出)
	Role       interface{} // 游戏角色数据
	Created    time.Time   // 建立时间
	Flag       int32       // 会话标记(0:初始化，1:已通过版本检查，2:登录中，3:登录成功, 4:已关闭)
	TotalWin   int64       // 进入之后的赢钱金额
	TotalBet   int64       // 进入之后的下注金额
	TotalRound int32       // 有下注的总局数

	sendChan chan interface{} // 发送出去的数据
	stopSend chan struct{}    // 发送停止信号
	dieOnce  int32            // 会话关闭信号
}

func(sess *Session) AddWinBet(win int64, bet int64)() {
	sess.TotalWin += win
	sess.TotalBet += bet
	sess.TotalRound++
}


// 不安全的发送，被发送的消息在另外一个线程编码，使用此方法，需保证被发送的对象线程安全
func (sess *Session) UnsafeSend(val interface{}) bool {
	if sess.sendChan == nil {
		return true
	}
	select {
	case <-sess.stopSend:
		return false
	case sess.sendChan <- val:
		return true
	}
}

func(sess *Session) SendRaw(val []byte) bool{
       return sess.UnsafeSend(val)
}

func (sess *Session) Send(val interface{}) bool {
	if sess.sendChan == nil {
		return true
	}
	if val, err := Coder.Encode(val); err == nil {
		return sess.UnsafeSend(val)
	}
	return false
}

func (sess *Session) SendError(id int32, code int32, m string, k string) {
	sess.UnsafeSend(&protocol.ErrorInfo{ReqId: id, Code: code, Msg: m, Key: k})
}

func (sess *Session) Start(stream protocol.GameStream) {
	if stream != nil {
		sess.sendChan = make(chan interface{}, 512)
		sess.stopSend = make(chan struct{})
		go sess.sendLoop(stream)
		sess.recvLoop(stream)
		select {
		case <-sess.stopSend:
			//等级发送结束再关闭连接
		}
	}
}

// 写消息循环
func (sess *Session) sendLoop(stream protocol.GameStream) {
	defer util.PrintPanicStack()
	defer close(sess.stopSend)
	for {
		select {
		case val, ok := <-sess.sendChan:
			if ok == false || val == nil {
				return
			}
			var data []byte
			switch val := val.(type) {
			case []byte:
				data = val
			default:
				if buf, err := Coder.Encode(val); err == nil {
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
		}
	}
}

func (sess *Session) recvLoop(stream protocol.GameStream) {
	defer util.PrintPanicStack()
	defer sess.Close()
	for sess.Flag < SESS_CLOSE {
		select {
		case <-sess.stopSend:
			sess.Flag = SESS_CLOSE
		case <-signal.Die():
			sess.Flag = SESS_CLOSE
		default:
			if data, err := stream.Recv(); err != nil {
				if err != io.EOF {
					log.Error(err)
				}
				return
			} else {
				if len(data) < protocol.HeadLen {
					return
				}
				id, arg, e := Coder.Decode(data)
				log.Debugf("Coder:%v,%v,%v,%v", Coder.Name(), id, arg, e)
				if e != nil || sess.Call(id, arg) == false {
					return
				}
			}
		}
	}
}

func (sess *Session) Call(id int32, arg interface{}) bool {
	select {
	case <-sess.stopSend:
		return false
	// 所有用户消息，放入主循环处理
	case messageChan <- &NetMessage{Session: sess, Id: id, Arg: arg}:
		return true
	}
}

func (sess *Session) Close() {
	if atomic.CompareAndSwapInt32(&sess.dieOnce, 0, 1) {
		if sess.sendChan == nil {
			return
		}
		sess.sendChan <- nil
	}
}

func(sess *Session)LockRoom(uid int32, win int64, bet int64, round int32) (*model.User, error) {
	return db.Driver.LockUserRoom(sess.AgentId, uid, KindId, RoomId, win, bet, round)
}

func(sess *Session)UnlockRoom() bool {
	log.Debugf("UnlockRoom:%v,%v", sess.AgentId, sess.UserId)
	if sess.Disposed == false && sess.UserId != 0 {
		sess.Disposed = true
		if db.Driver.UnlockUserRoom(sess.AgentId, sess.UserId, RoomId, sess.TotalWin, sess.TotalBet, sess.TotalRound) {
			log.Debugf("UnlockRoom success:%v,%v", sess.AgentId, sess.UserId)
			return true
		}
	}
	log.Debugf("UnlockRoom fail:%v,%v", sess.AgentId, sess.UserId)
	return false
}