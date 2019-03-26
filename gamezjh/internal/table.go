package internal

import (
	log "github.com/sirupsen/logrus"
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
	Id         int32             // 桌子ID
	CurId      int32             // 当前局的ID
	State      GameState         // 0:准备游戏;1:游戏中;2:结算
	RoleCount  int32             // 玩家数量
	Roles      [chairCount]*Role // 所有游戏玩家,按位置
	round      *GameRound        // 1局
	delay      int32             // 持续秒数
	waitSecond int32             // 等待秒数
}

func NewTable() *Table {
	t := &Table{}
	return t
}

func (table *Table) MustWin() bool {
	return room.Config.WinRate > gameRand.Int31n(1000)
}

func(table *Table)GetPlayer()[]*zjh.Player {
	// TODO: 获取玩家信息
	if table.round != nil {
		return table.round.players
	}
	return nil
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
	i := table.RoleCount
	players := make([]*zjh.Player, 0, i)
	bills := make([]*zjh.GameBill, 0, i)
	// 随机首家
	first := gameRand.Int31n(i)
	for i := first; i < chairCount+first; i++ {
		if role := table.Roles[i%chairCount]; role != nil {
			role.Reset()
			role.player.State = zjh.Player_Playing
			players = append(players, role.GetPlayer())
			bills = append(bills, role.bill)
		}
	}
	id := room.NewGameRoundId()
	round := &GameRound{
		firstChair: first,
		players:    players,
		GameRound: zjh.GameRound{
			Id:    id,
			Start: room.Now(),
			Room:  room.RoomId,
			Tab:   table.Id,
			Bill:  bills,
		},
	}
	table.round = round
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
		if role.player.State != zjh.Player_Read {
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
					role.player.State = zjh.Player_Read
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
		if role.player.State != zjh.Player_Read {
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
		table.delay = 2
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
		Id:   table.CurId,
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
	round.End = room.Now()
	room.SaveLog(round)
}
