package internal

import (
	log "github.com/sirupsen/logrus"

	//"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	//"local.com/abc/game/room"
)

// 机器人接口
type Robot interface {
	// 每一局开始的时候调用
	Start (role *Role)
	// 自己说话
	Play(role *Role, ring int32, my bool)
	// 已弃牌
	Discard(role *Role)
	// 已失败
	Lose(role *Role)
	// 每一局结束的时候调用
	End(role *Role)
}

type TestRobot struct {
   first int32
}

func (*TestRobot)rand(i int32)int32 {
	return gameRand.Int31n(i)
}

// 每一局开始的时候调用
func(t *TestRobot)Start(role *Role){
	t.first = 0
}


// 每一局结束的时候调用
func(t *TestRobot)End(role *Role){

}

func(t *TestRobot)Play(role *Role, ring int32, my bool){
	if ring == 1 {
		t.ring1(role,my)
	} else if ring == 2 {
		t.ring2(role,my)
	}else {
		t.play(role,ring,my)
	}
}

// 第1轮,只能弃牌+跟注/加注
func(t *TestRobot) ring1(role *Role, my bool){
	//// 1000分之1的概率在不看牌时弃牌
	//if t.first == 0 {
	//	t.first++
	//	if t.rand(1000) == 1 {
	//		role.Discard(false)
	//	}
	//}
	if my {
		role.AddBet(0)
	}
}

// 第2轮, 只能弃牌/看牌+跟注/加注
func(t *TestRobot) ring2(role *Role, my bool) {
	if my {
		if t.rand(100) < 50{
			role.AddBet(0)
		}
		return
	}
	// 1000分之1的概率在不看牌时弃牌
	if role.player.Look {
		// 每一秒都有100分之5弃牌
		if t.rand(100) < 5 {
			role.Look()
		}

		// 如果牌为单张，小于Q,则放弃
		if role.poker.Zilch() {
			if role.poker.Power < 12439 {
				role.Discard(false)
			}
		}


	} else {
		// 每一秒都有100分之5看牌
		if t.rand(100) < 5 {
			role.Look()
		}
	}
}

// 第2轮以上, 只能弃牌/看牌+跟注/加注/比牌/allin
func(t *TestRobot)play(role *Role, ring int32, my bool) {
	if my {
		if t.rand(100) < 30 {
			role.AddBet(0)
		}
		if t.rand(100) < 50 && role.table.firstAllin == nil {
			// 选择一个对手
			opp := make([]int32, 0, 2)
			for _, player := range role.table.players {
				if player.State == zjh.Player_Playing && player != role.player {
					opp = append(opp, player.Id)
				}
			}
			if len(opp) > 0 {
				oppid := opp[t.rand(int32(len(opp)))]
				role.Compare(oppid)
			} else {
				log.Debugf("选择对手失败")
			}
		}
		if t.rand(100) < 30 {
			role.Allin()
		}
		return
	}
	// 1000分之1的概率在不看牌时弃牌
	if role.player.Look == false {
		// 每一秒都有100分之5弃牌
		if t.rand(100) < 5 {
			role.Look()
		}
	}
}

// 已弃牌
func(*TestRobot)Discard(role *Role){
	// TODO：钱不够就退出桌子补充

}

// 已失败
func(*TestRobot)Lose(role *Role){
	// TODO：钱不够就退出桌子补充

}