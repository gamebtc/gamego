package room

import (
	log "github.com/sirupsen/logrus"

	//"local.com/abc/game/db"
	//"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

// 主线程调用，玩家上线
func userOnline(sess *Session, uid int32) {
	log.Debugf("userOnline:%v,uid:%v", sess.AgentId, uid)
	// 登录检查
	oldSess := GetUser(uid)

	if oldSess == nil && Config.Lock > 0 {
		// 检查房间配置
		log.Debugf("id:%v,uid:%v,kid:%v,room:%v,登录失败:%v", sess.AgentId, uid, KindId, RoomId, "房间已锁定")
		sess.SendError(int32(protocol.MsgId_LoginRoomReq), 1000, "房间已锁定", "")
		sess.Close()
		return
	}

	if oldSess == sess {
		return
	}

	// 锁定玩家
	win, bet, round := int64(0), int64(0), int32(0)
	if oldSess != nil && oldSess.Role != nil {
		win, bet, round = sess.TotalWin, sess.TotalBet, sess.TotalRound
	}
	user, err := sess.lockRoom(uid, win, bet, round)
	if err != nil {
		// 发送错误消息
		log.Debugf("id:%v,uid:%v,kid:%v,room:%v,登录失败:%v", sess.AgentId, uid, KindId, RoomId, err.Error())
		sess.SendError(int32(protocol.MsgId_LoginRoomReq), 1000, "登录失败", err.Error())
		sess.Close()
		return
	}
	if user == nil || user.Id == 0{
		log.Debugf("id:%v,uid:%v,kid:%v,room:%v,登录失败2", sess.AgentId, uid, KindId, RoomId)
		sess.SendError(int32(protocol.MsgId_LoginRoomReq), 1000, "登录失败2", "")
		sess.Close()
		return
	}

	sess.UserId = uid
	sess.locked = true
	if oldSess != nil {
		sess.Role = oldSess.Role
		oldSess.UserId = 0
		oldSess.locked = false
		AddUser(sess)
		// 发送登录游戏信息
		sess.UnsafeSend(&protocol.LoginRoomAck{
			Room: RoomId,
			Kind: KindId,
		})

		hall.UserReline(oldSess, sess)
		// 发送错误消息,顶掉玩家
		oldSess.Close()
	} else {
		// 检查所带金币
		coin := user.Coin
		if coin < Config.DoorMin || coin > Config.DoorMax {
			log.Debugf("id:%v,uid:%v,kid:%v,room:%v,登录失败:%v", sess.AgentId, uid, KindId, RoomId, "金币不足")
			sess.SendError(int32(protocol.MsgId_LoginRoomReq), 1000, "金币不足", "")
			sess.Close()
			return
		}
		user.Online = true
		AddUser(sess)
		// 发送登录游戏信息
		sess.UnsafeSend(&protocol.LoginRoomAck{
			Room: RoomId,
			Kind: KindId,
		})
		hall.UserOnline(sess, user)
	}
}

// 主线程调用，玩家下线
func userOffline(sess *Session) {
	log.Debugf("userOffline:%v,uid:%v", sess.AgentId, sess.UserId)
	hall.UserOffline(sess)
}
