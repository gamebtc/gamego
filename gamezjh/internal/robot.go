package internal

import (
	log "github.com/sirupsen/logrus"

	//"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

// 机器人AI接口
type Robot interface {
	// 每一局开始的时候调用
	Start(role *Role)
	// 自己说话
	Play(role *Role)
	// 已弃牌
	Discard(role *Role)
	// 已失败
	Lose(role *Role)
	// 每一局结束的时候调用
	End(role *Role)
}

type RobotAi struct {
	step int   // 步数
	ring int32 // 已做过决定的轮数
	down int32 // 这一步的时间倒计时
	rand int32 // 这一步的随机数[0-1000)
	my   bool  // 是否应该我执行命令
}

// 每一局开始的时候调用
func (t *RobotAi) Start(role *Role) {
	t.step = -1
	t.ring = 0
	t.down = 0
	t.rand = 0
	t.my = false
}

// 每一局结束的时候调用
func (t *RobotAi) End(role *Role) {

}

func (t *RobotAi) Play(role *Role) {
	table := role.table
	round := table.round
	step := len(round.Log)
	if t.step != step {
		t.step = step
		t.rand = gameRand.Int31n(1000)
		t.my = table.isRunner(role.User.Id)
		if t.my {
			t.down = gameRand.Int31n(4) + 1
		} else {
			t.down = gameRand.Int31n(6) + 1
		}
	}
	if t.down > 0 {
		t.down--
		return
	}

	if round.Ring == 1 {
		if t.my {
			t.ring1my(role)
		}else{
			t.ring1(role)
		}
	} else if round.Ring == 2 {
		if t.my {
			t.ring2my(role)
		}else{
			t.play(role)
		}
	} else {
		if t.my {
			t.playmy(role)
		}else{
			t.play(role)
		}
	}
}

// 第1轮,只能弃牌/有条件看牌？
func (t *RobotAi) ring1(role *Role) {
	// 弃牌
	player := role.player
	if player.Look {
		if role.poker.Power < 10000 {

		}
	} else {
		if player.Bet > room.Config.Ante {
			if t.rand < 50 {
				role.Look()
			}
		}
	}
}

// 第2轮以上, 只能弃牌/看牌
func (t *RobotAi) play(role *Role) {
	// 1000分之1的概率在不看牌时弃牌
	if role.player.Look {
		// 每一秒都有100分之5弃牌
		if t.rand < 5 {
			role.Look()
		}
		// 如果牌为单张，小于Q,则放弃
		if role.poker.Zilch() {
			if role.poker.Power < 12439 {
				role.Discard(false)
			}
		}
		return
	}
	// 每一秒都有100分之5看牌
	if t.rand < 5 {
		role.Look()
	}
}

// 我的第1轮弃牌/跟注/加注
func (t *RobotAi) ring1my(role *Role){
	// 弃牌+跟注/加注
	role.AddBet(0)
	return
}

// 我的第2轮, 只能弃牌/看牌/跟注/加注
func (t *RobotAi) ring2my(role *Role) {
	if t.rand < 500 {
		role.AddBet(0)
	}
	return
}


// 我的第2轮以上弃牌/看牌/跟注/加注/比牌/allin
func (t *RobotAi) playmy(role *Role) {
	if role.poker.Power >= 20000 {
		role.Allin()
		return
	}

	if t.rand < 30 {
		role.AddBet(0)
		return
	}

	if t.rand < 50 && role.table.firstAllin == nil {
		// 选择一个对手
		opp := make([]int32, 0, 2)
		for _, player := range role.table.players {
			if player.State == zjh.Player_Playing && player != role.player {
				opp = append(opp, player.Id)
			}
		}
		if len(opp) > 0 {
			oppid := opp[gameRand.Int31n(int32(len(opp)))]
			role.Compare(oppid)
			return
		} else {
			log.Debugf("选择对手失败")
		}
		return
	}

	if t.rand < 30 {
		role.Allin()
	}
}

// 已弃牌
func (*RobotAi) Discard(role *Role) {
	// TODO：钱不够就退出桌子补充

}

// 已失败
func (*RobotAi) Lose(role *Role) {
	// TODO：钱不够就退出桌子补充

}
