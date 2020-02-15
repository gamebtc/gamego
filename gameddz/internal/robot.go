package internal

import (
	//"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
)

// 机器人AI接口,机器人控制打牌像真人就行了
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
	step    int   // 步数
	lastAdd int   // 我最后一次加注的位置
	ring    int32 // 已做过决定的轮数
	down    int32 // 这一步的时间倒计时
	my      bool  // 是否应该我执行命令
	first   bool  // 我是否是首家
}

// 每一局开始的时候调用
func (t *RobotAi) Start(role *Role) {
	t.step = -1
	t.lastAdd = -1
	t.ring = 0
	t.down = 0
	t.my = false
	t.first = role.Player == role.table.players[0]
}

// 每一局结束的时候调用
func (t *RobotAi) End(role *Role) {

}

func rn() int32 {
	return gameRand.Int31n(1000)
}

func (t *RobotAi) Play(role *Role) {
	table := role.table
	round := table.round
	step := len(round.Log)
	if t.step != step {
		t.step = step
		t.my = table.isRunner(role.Id)
		if t.my {
			// 自己，1-3秒做一个决定
			t.down = 1 + gameRand.Int31n(3)
		} else {
			// 没有轮到自己，1-3秒做一个决定
			t.down = gameRand.Int31n(4)
		}
	}
	if t.down > 0 {
		t.down--
		return
	}
	if t.my {
		//log.Debugf("[%v]robot:%v", round.Ring, role.Id)
		switch round.Ring {
		case 1:
			t.ring1my(role)
		case 2:
			t.ring2my(role)
		default:
			if table.firstAllin == nil {
				t.playmy(role)
			} else {
				t.playmyAllin(role)
			}
		}
		t.ring = round.Ring
	} else if round.Ring == 1 {
		//if role.player.Bet > room.Config.Ante {
		//	t.play(role) //可以看牌/弃牌
		//} else {
		//	// TODO:只能弃牌
		//}
	} else {
		if table.firstAllin == nil {
			t.play(role)
		} else {
			t.playAllin(role)
		}
	}
}

// 第2轮以上, 只能弃牌/看牌/忽略
func (t *RobotAi) play(role *Role) {
	table := role.table
	// 有PK胜利或者只有2个玩家时忽略
	if len(role.Bill.Pk) > 0 || (table.playCount <= 2 && len(table.players) > 2) {
		return
	}

	r := rn()
	if !role.Player.Look {
		// 每一次都有1000分之10看牌
		if r < (table.lookCount+1)*150 {
			role.Look()
		}
		return
	}

	// 我已看牌并下注,其它人没有加过注则不主动放弃
	if t.lastAdd > -1 && (!t.otherAddBet(role)) {
		return
	}

	if r < 400 {
		//  牌小于♣K♠0♥2=11099,则放弃
		minPower := 10000 + (2000*table.compCount + 500*table.playCount) + r
		if role.Bill.Weight <= minPower {
			role.Discard(false)
		}
	}
}

func (t *RobotAi) otherAddBet(role *Role) bool {
	logs := role.table.round.Log
	for i := t.lastAdd; i < len(logs); i++ {
		t := logs[i].Type
		if t == zjh.ActionType_ActionCompare ||
			t == zjh.ActionType_ActionAddBet ||
			t == zjh.ActionType_ActionAllin {
			return true
		}
	}
	return false
}

func (t *RobotAi) addBet(role *Role, rand bool) {
	if rand {
		role.randAddBet()
	} else {
		role.AddBet(0)
	}
	if role.Player.Look {
		t.lastAdd = len(role.table.round.Log)
	}
}

//  有人全压, 只能弃牌/看牌/忽略
func (t *RobotAi) playAllin(role *Role) {
	table := role.table
	// 有PK胜利或者只有2个玩家时忽略
	if len(role.Bill.Pk) > 0 || table.playCount <= 2 {
		return
	}

	if !role.Player.Look {
		// 每一次都有1000分之50看牌
		if rn() < (50*table.lookCount + 50) {
			role.Look()
		}
		return
	}

	if r := rn(); r < 100 {
		//  牌小于♣K♠0♥2=11099,则放弃
		minPower := 10000 + (1000*table.compCount + 1000*table.playCount) + r
		if role.Bill.Weight <= minPower {
			role.Discard(false)
		}
	}
}

// 我的第1轮弃牌/跟注/加注
func (t *RobotAi) ring1my(role *Role) {
	// 我是首家
	//// 弃牌千分之1
	//if r < 1 {
	//	if first == false {
	//		role.Discard(false)
	//		return
	//	}
	//}

	r := rn()
	// 我是首家1000分之100加满注, 不是首家1000分之20加注
	if (t.first && r < 100) || r < 20 {
		t.addBet(role, true)
		return
	}

	// 默认跟注
	role.AddBet(0)
	return
}

// 我的第2轮, 只能弃牌/看牌/跟注/加注
func (t *RobotAi) ring2my(role *Role) {
	table := role.table
	player := role.Player

	r := rn()
	if !player.Look {
		if r < 100 {
			role.Look()
		} else if r < 100+200 {
			role.randAddBet()
		} else {
			role.AddBet(0)
		}
		return
	}

	if table.playCount > 2 && r < 800 {
		//  已看牌,牌小于Power:10019, number:861986,key:|♦K♥7♥2|,则放弃
		minPower := 10000 + (1200*table.compCount + 500*table.playCount) - r
		if int32(role.Bill.Weight) <= minPower {
			role.Discard(false)
			return
		}
	}

	if r < 400 && role.Bill.Weight < 10000+rn() {
		role.Discard(false)
		return
	}

	r = rn()
	if r < 300 {
		role.randAddBet()
	} else {
		role.AddBet(0)
	}
	return
}

// 我的第3轮及以上有人全压, 弃牌/看牌/allin
func (t *RobotAi) playmyAllin(role *Role) {
	player := role.Player
	r := rn()
	if !player.Look {
		if r < 500 {
			//  Weight:19631, number:4001058,key:|♠K♦K♥2|
			if role.Bill.Weight >= 19631 {
				role.Allin()
				return
			}
		}
		role.Look()
		return
	}

	//Weight:17615, number:2496034,key:|♥6♣6♥2|
	if role.Bill.Weight < 17615+r {
		role.Discard(false)
		return
	}

	role.Allin()
}

// 我的第3轮及以上,弃牌/看牌/跟注/加注/比牌/allin
func (t *RobotAi) playmy(role *Role) {
	table := role.table
	player := role.Player

	r := rn()
	if !player.Look {
		if r < 100 {
			role.Look()
		} else if r < 100+200 {
			role.randAddBet()
		} else {
			role.AddBet(0)
		}
		return
	}

	if r < 50 {
		role.Allin()
		return
	}

	if r < 300 {
		// 选择一个对手
		opp := t.findOpp(role)
		if opp > 0 {
			role.Compare(opp)
			return
		}
	}

	if r < 300+200 {
		// Weight:19631, number:4001058,key:|♠K♦K♥2|
		if role.Bill.Weight >= 19631+r*5 {
			role.Allin()
			return
		}
		//  已看牌,牌小于Power:10019, number:861986,key:|♦K♥7♥2|,则放弃
		minPower := 10000 + (1200*table.compCount + 500*table.playCount) - r
		if int32(role.Bill.Weight) <= minPower {
			role.Discard(false)
			return
		}
	}

	r = rn()
	if r < 300 {
		role.randAddBet()
	} else {
		role.AddBet(0)
	}
	return
}

// 选择一个玩家PK，优化选择已比赢的，再选择已看牌的，再选择没有看牌的
func (t *RobotAi) findOpp(role *Role) int32 {
	bill := role.table.round.Bill
	players := role.table.players
	for i, player := range players {
		if player.State == zjh.Player_Playing &&
			len(bill[i].Pk) > 0 {
			return player.Id
		}
	}
	for _, player := range players {
		if player.State == zjh.Player_Playing &&
			player.Look {
			return player.Id
		}
	}
	for _, player := range players {
		if player.State == zjh.Player_Playing {
			return player.Id
		}
	}
	return 0
}

// 已弃牌
func (*RobotAi) Discard(role *Role) {
	// TODO：钱不够就退出桌子补充

}

// 已失败
func (*RobotAi) Lose(role *Role) {
	// TODO：钱不够就退出桌子补充

}
