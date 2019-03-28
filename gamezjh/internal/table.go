package internal

import (
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

const chairCount = 5

type GameState int32

const (
	// 准备游戏
	GameStateReady GameState = 0
	// 游戏中
	GameStatePlaying GameState = 1
	// 结算
	GameStateDeal GameState = 2
)

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	Id          int32             // 桌子ID
	CurId       int32             // 当前局的ID
	State       GameState         // 0:准备游戏;1:游戏中;2:结算
	RoleCount   int32             // 玩家数量
	Roles       [chairCount]*Role // 所有游戏玩家,按位置,玩家离开后Role退出
	round       *GameRound        // 1局
	delay       int32             // 持续秒数
	waitSecond  int32             // 等待秒数
	poker       []byte            // 所有的牌
	pokerOffset int32             // 牌的位置
	ring        int32             // 下注轮数
	firstChair  int32             // 首家位置
	waitChair   int32             // 等待位置
	players     []*zjh.Player     // 游戏中的玩家
	winner      *Role             // 最后的赢家
	curAnte     int32             // 没有看牌的当前底注,
	allinBet    int32             // 全压下注
}

func NewTable() *Table {
	t := &Table{}
	return t
}

func (table *Table) MustWin() bool {
	return room.Config.WinRate > gameRand.Int31n(1000)
}

func (table *Table) GetPlayer() []*zjh.Player {
	// TODO: 获取玩家信息
	return table.players
}

func (table *Table) FindRole(id int32) *Role {
	for _, role := range table.Roles {
		if role != nil && role.User.Id == id {
			return role
		}
	}
	return nil
}

// 试着结算
func (table *Table) TryDeal() bool {
	if table.winner != nil {
		return true
	}
	var winner *Role
	// 只有一个玩家为Playing或者Allin时，玩家胜利
	for _, role := range table.Roles {
		if role != nil && (role.player.State == zjh.Player_Playing || role.player.State == zjh.Player_Allin) {
			if winner != nil {
				return false
			}
			winner = role
		}
	}
	if winner == nil {
		return false
	}
	round := table.round
	table.winner = winner
	winner.player.State = zjh.Player_Win
	// 给赢家结算
	winCoin := round.Sum
	winner.Win(winCoin)

	// 重置玩家和机器人之间的输赢

	// 更新全局池子
	//if lucky > 0{
	//	ChangeGoldPool()
	//}

	// 保存牌局
	round.End = room.Now()
	room.SaveLog(round)
	// 结算结果发给玩家
	for _, role := range table.Roles {
		if role != nil && role.IsRobot() == false {
			// 看自己的牌+PK过的牌
			poker := make([]byte, 3*chairCount)
			for _, opp := range table.players {
				if opp.Id == role.User.Id || role.IsOpponent(opp.Id) {
					i := opp.Chair - 1
					copy(poker[i:i+3], table.poker[i:i+3])
				}
			}
			// 发送游戏结束
			role.UnsafeSend(&zjh.GameEndAck{
				Id:     table.CurId,
				Win:    winCoin,
				Winner: winner.User.Id,
				Coin:   role.Coin,
				Poker:  poker,
				Lucky:  round.Lucky,
			})
		}
	}
	return true
}

func (table *Table) TryAutoComp() bool {
	allinRoles := make([]*Role, 0, 2)
	for _, role := range table.Roles {
		if role != nil {
			if role.player.State == zjh.Player_Allin {
				allinRoles = append(allinRoles, role)
			} else if role.player.State == zjh.Player_Playing {
				allinRoles = nil
				break
			}
		}
	}
	if len(allinRoles) < 2 {
		return false
	}
	// 2个以上玩家进行PK
	// 设置PK列表
	for _, a := range allinRoles {
		for _, b := range allinRoles {
			if a != b {
				a.bill.Pk = append(a.bill.Pk, b.User.Id)
			}
		}
	}
	// 开始PK
	a := allinRoles[0]
	for i := 1; i < len(allinRoles); i++ {
		if b := allinRoles[i]; a.pk(b) == false {
			a = b
		}
	}
	// 发送自动比牌结果
	table.SendToAll(&zjh.ActionAck{
		Uid:  a.Id, //赢家ID
		Type: zjh.ActionType_ActionAutoCompare,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   a.Id,
		Type:  zjh.ActionType_ActionAutoCompare,
	}
	table.round.Log = append(table.round.Log, log)
	return true
}

// 下一个等待者
func (table *Table) NextWait() bool {
	if table.winner != nil {
		return false
	}
	const waitSecond = 15
	for i := table.waitChair + 1; i < table.waitChair+chairCount; i++ {
		chair := i % chairCount
		if chair == table.firstChair {
			table.ring++
		}
		role := table.Roles[chair]
		if role != nil && role.player.State == zjh.Player_Playing {
			table.waitChair = role.player.Chair
			role.player.Down = waitSecond
			return true
		}
	}
	return false
}

// 增加玩家
func (table *Table) addRole(role *Role) bool {
	for i, v := range table.Roles {
		if v == nil {
			role.player = &zjh.Player{
				Id:    role.Id,
				Icon:  role.Icon,
				Vip:   role.Vip,
				Name:  role.Name,
				Coin:  role.Coin,
				State: zjh.Player_None,
			}
			role.table = table
			role.player.Chair = int32(i + 1)
			table.Roles[i] = role
			table.RoleCount++
			// 真实玩家
			if !role.IsRobot() {
				table.sendGameInit(role)
			}
			return true
		}
	}
	return false
}

// 所有机器人退出
func (table *Table) freeRobots() {
	for i, role := range table.Roles {
		if role != nil && role.IsRobot() {
			table.RoleCount++
			table.Roles[i] = nil
			role.player = nil
			role.table = nil
			role.bill = nil
		}
	}
}

// 初始化场景
func (table *Table) sendGameInit(role *Role) {
	if !role.IsRobot() {
		ack := &zjh.GameInitAck{
			Table:  table.Id,
			Id:     table.CurId,
			State:  int32(table.State),
			Player: nil,
			Poker:  nil,
		}
		if round := table.round; round != nil {
			ack.Ring = table.round.Ring
			ack.Player = table.GetPlayer()
		}
		if bill := role.bill; bill != nil {
			//ack.Ring = bill.
		}
		role.Send(ack)
	}
}

func (table *Table) newGameRound() {
	count := table.RoleCount
	ante := room.Config.Ante // 底注
	id := room.NewGameRoundId()
	table.round = &GameRound{
		Id:    id,
		Start: room.Now(),
		Room:  room.RoomId,
		Tab:   table.Id,
		Bill:  make([]*zjh.GameBill, 0, count),
		Ante:  ante,
	}
	table.players = make([]*zjh.Player, 0, count)
	table.poker = model.NewPoker(1, false, true)
	table.pokerOffset = 0
	// 随机首家
	first := gameRand.Int31n(count)
	table.firstChair = first
	table.waitChair = first
	for i := first; i < chairCount+first; i++ {
		if role := table.Roles[i%chairCount]; role != nil {
			role.bill = &zjh.GameBill{
				Uid:   role.User.Id,
				Job:   role.User.Job,
				Coin:  role.Coin,
				Poker: table.poker[table.pokerOffset : table.pokerOffset+3],
			}
			table.pokerOffset += 3
			role.player.State = zjh.Player_Playing
			role.player.Look = false
			role.player.Down = 0
			// 打底
			role.decCoin(ante)

			table.players = append(table.players, role.GetPlayer())
			table.round.Bill = append(table.round.Bill, role.bill)
		}
	}
	// 第一轮自动
	table.ring = 1
}

func (table *Table) Init() {

}

func (table *Table) Start() {
	table.gameReady()
}

func (table *Table) Update() {
	table.delay--
	switch table.State {
	case GameStateReady:
		if room.Config.Pause != 0 {
			table.delay++
			return
		}
		table.gameWait()
	case GameStatePlaying:
		if table.delay > 0 {
			table.gamePlay()
		} else {
			table.gameDeal()
		}
	case GameStateDeal:
		if table.delay <= 0 {
			table.gameReady()
		}
	}
}

// 发送消息给所有在线玩家
func (table *Table) SendToAll(val interface{}) {
	if val, err := room.Encode(val); err != nil {
		for _, role := range table.Roles {
			if role != nil {
				role.UnsafeSend(val)
			}
		}
	}
}

// 发送消息给其它玩家
func (table *Table) SendToOther(val interface{}, my *Role) {
	if val, err := room.Encode(val); err != nil {
		for _, role := range table.Roles {
			if role != nil && role != my {
				role.UnsafeSend(val)
			}
		}
	}
}

// 检查是否有真实玩家
func (table *Table) ExistsRealPlayer() bool {
	for _, role := range table.Roles {
		if role.IsRobot() == false {
			return true
		}
	}
	return false
}

// 准备时间，15秒
const readySecond = 15

// 准备
func (table *Table) gameReady() {
	table.delay = 2
	table.State = GameStateReady
	table.waitSecond = 0
	table.CurId += 1
	log.Debugf("%v准备:%v", table.Id, table.CurId)
	// 设置玩家举手倒计时
	for _, role := range table.Roles {
		if role == nil {
			continue
		}
		if role.player.State != zjh.Player_Ready {
			role.player.State = zjh.Player_None
			role.player.Down = readySecond
		}
	}
}

func (table *Table) gameWait() {
	table.waitSecond++
	// 真人数量
	realCount := 0
	// 所有人都已经准备好
	allReady := true
	for _, role := range table.Roles {
		if role == nil {
			continue
		}
		if role.player.State == zjh.Player_None {
			role.player.Down--
			if role.IsRobot() {
				// 机器人准备时间0-4秒,不大于5秒
				if gameRand.Int31n(4) == 0 || table.waitSecond >= 5 {
					role.player.State = zjh.Player_Ready
				}
			} else {
				// TODO： 真人超时了，强制退出
				if role.player.Down < 0 {
					continue
				}
			}
		}
		if !role.IsRobot() {
			realCount++
		}
		if role.player.State != zjh.Player_Ready {
			allReady = false
		}
	}

	if realCount == 0 {
		// 没有真人了机器人退出
		table.freeRobots()
		return
	}

	if allReady && table.RoleCount >= 2 {
		// 所有人都已准备，有2个人就可以开始了
		table.gameOpen()
	} else {
		table.delay = readySecond
		log.Debugf("%v等待玩家:%v", table.Id, table.CurId)
	}
}

// 开始
func (table *Table) gameOpen() {
	// 发送开始下注消息给所有玩家
	table.delay = 15
	table.State = GameStatePlaying
	table.waitSecond = 0
	log.Debugf("%v开始下注:%v", table.Id, table.CurId)
	table.newGameRound()

	// 发送消息给玩家
	table.SendToAll(&zjh.GameStartAck{
		Id:     table.CurId,
		Player: table.GetPlayer(),
	})

}

func (table *Table) gamePlay() {

}

// 结算
func (table *Table) gameDeal() {
	table.delay = 5
	table.State = GameStateDeal
	log.Debugf("%v发牌结算:%v", table.Id, table.CurId)
	// 发牌结算

	round := table.round
	round.Ring = table.ring
	round.End = room.Now()
	room.SaveLog(round)
}
