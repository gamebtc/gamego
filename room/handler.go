package room

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
)

// 主线程调用，玩家上线
func userOnline(sess *Session) {
	log.Debugf("userOnline1:%v", sess)
	// 登录检查
	oldSess := roomer.GetUser(sess.UserId)
	if oldSess == sess {
		log.Debugf("userOnline2:%v", sess)
		return
	}
	if oldSess != nil {
		oldSess.Online = false

		roomer.SetUser(sess)
		sess.Online = true

		log.Debugf("userOnline3:%v", sess)
		roomer.UserReline(oldSess, sess)
		// 发送错误消息,顶掉玩家
		oldSess.Close()
	} else {
		// 锁定玩家
		log.Debugf("userOnline4:%v", sess)
		user, err := driver.LockUserRoom(sess.Id, sess.UserId, Config.Kind, Config.Id)
		if err != nil {
			// 发送错误消息
			sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "登录失败", err.Error())
			sess.Close()
			return
		}
		log.Debugf("userOnline5:%v", user)
		if user == nil {
			log.Debugf("id:%v,uid:%v,kid:%v,room:%v", sess.Id, sess.UserId, Config.Kind, Config.Id)
			sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "登录失败2", "")
			sess.Close()
			return
		}
		// 检查房间配置
		coin := user.Bag[Config.CoinKey]
		if coin < Config.DoorMin || coin > Config.DoorMax {
			log.Debugf("userOnline6:%v", user)
			// 所带金币不符合要求发送错误消息
			sess.SendError(int32(msg.MsgId_LoginRoomReq), 1000, "金币不足", "")
			sess.Close()
			return
		}
		user.Coin = coin
		roomer.UserOnline(sess, user)
		log.Debugf("userOnline7:%v", user)
	}
}

// 主线程调用，玩家下线
func userOffline(s *Session) {
	s.Online = false
	roomer.UserOffline(s)
}

func LoadRobot(agentid int64, userid int32) *model.User {
	// 锁定玩家
	user, err := driver.LockUserRoom(agentid, userid, Config.Kind, Config.Id)
	if err == nil && user != nil {
		// 检查房间配置
		user.Coin = user.Bag[Config.CoinKey]
	}
	return user
}
