package room

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

// 主线程调用，玩家上线
func userOnline(sess *Session) {
	log.Debugf("userOnline:%v", sess)
	// 登录检查
	oldSess := roomer.GetUser(sess.UserId)
	if oldSess == sess {
		return
	}
	// 锁定玩家
	user, err := driver.LockUserRoom(sess.AgentId, sess.UserId, KindId, RoomId)
	if err != nil {
		// 发送错误消息
		sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "登录失败", err.Error())
		sess.Close()
		return
	}
	if user == nil {
		log.Debugf("id:%v,uid:%v,kid:%v,room:%v", sess.AgentId, sess.UserId, KindId, RoomId)
		sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "登录失败2", "")
		sess.Close()
		return
	}
	if oldSess != nil {
		oldSess.Online = false

		roomer.SetUser(sess)
		sess.Online = true

		roomer.UserReline(oldSess, sess)
		// 发送错误消息,顶掉玩家
		oldSess.Close()
	} else {
		// 检查房间配置
		coin := user.Bag[CoinKey]
		if coin < Config.DoorMin || coin > Config.DoorMax {
			// 所带金币不符合要求发送错误消息
			sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "金币不足", "")
			sess.Close()
			return
		}
		user.Coin = coin
		roomer.UserOnline(sess, user)
	}
}

// 主线程调用，玩家下线
func userOffline(sess *Session) {
	log.Debugf("userOffline:%v", sess)
	defer func() {
		if sess.Playing == false {
			driver.UnlockUserRoom(sess.AgentId, sess.UserId, Config.Id)
		}
	}()
	sess.Online = false
	roomer.UserOffline(sess)
}

func LoadRobot(agentId int64, userId int32) *model.User {
	// 锁定玩家
	user, err := driver.LockUserRoom(agentId, userId, KindId, RoomId)
	if err == nil && user != nil {
		// 检查房间配置
		user.Coin = user.Bag[CoinKey]
	}
	return user
}
