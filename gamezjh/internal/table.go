package internal

import (
	log "github.com/sirupsen/logrus"
	//"local.com/abc/game/protocol/folks"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

const chairCount = 5

type GameState int32
const(
	// 等待玩家进入
	GameStateWait GameState = 0
	// 准备游戏
	GameStateReady GameState = 1
	// 游戏中
	GameStatePlaying GameState = 2
	// 结算
	GameStateDeal GameState = 3
)

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	Id        int32             // 桌子ID
	CurId     int32             // 当前局的ID
	State     GameState         // 0:等待玩家进入;1:准备游戏;2:游戏中;3:结算
	Roles     [chairCount]*Role // 所有游戏玩家,按位置
	Player    []*Role           // 游戏中的玩家
	round     *GameRound        // 1局
	continued int32             // 持续秒数
}

func NewTable() *Table {
	t := &Table{
	}
	return t
}

func(table *Table)MustWin()bool {
	return  mustWinRate > mustWinRand.Int31n(100)
}

// 增加真实的玩家
func (table *Table) AddRole(role *Role) bool {
	for i := int32(0); i < chairCount; i++ {
		if table.Roles[i] == nil {
			role.Chair = i
			table.Roles[i] = role
			// 真实玩家
			ack := &zjh.GameInitAck{
				Id:    table.round.Id,
				State: int32(table.State),
				//Sum:   table.round.Group,
			}
			//if bill := role.bill; bill != nil {
			//	ack.Bet = bill.Group
			//}
			role.Send(ack)

			return true
		}
	}
	return false
}

func (table *Table) RoleCount() int {
	i := 0
	for _, role := range table.Roles {
		if role != nil {
			i++
		}
	}
	return i
}

func (table *Table) NewGameRound() {
	i := table.RoleCount()
	players := make([]*Role, 0, i)
	bills := make([]*zjh.GameBill, 0, i)
	// 随机首家
	first := mustWinRand.Intn(i)
	for i := first; i < chairCount+first; i++ {
		if role := table.Roles[i%chairCount]; role != nil {
			role.Reset()
			players = append(players, role)
			bills = append(bills, role.bill)
		}
	}
	table.Player = players
	id := room.NewGameRoundId()
	round := &GameRound{
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

func (table *Table) Init(){

}


func (table *Table) Start() {
	table.continued = 1
	table.State = GameStateWait
}

func (table *Table) Update() {
	table.continued--
	switch table.State {
	case GameStateWait:
		if room.Config.Pause == 0 {
			if table.continued <= 0 {
				table.continued = 1
				gameReady(table)
			}
		} else {
			table.continued++
		}
	case GameStateReady:
		if table.continued <= 0 {
			table.continued = 12
			gameOpen(table)
		}
	case GameStatePlaying:
		if table.continued != 0 {
			gamePlay(table)
		} else {
			table.continued = 5
			gameDeal(table)
		}
	case GameStateDeal:
		if table.continued <= 0 {
			table.continued = 5
			gameWait(table)
		}
	}
}


// 发送消息给所有在线玩家
func(table *Table)SendToAll(val interface{}) {
	if len(table.Roles) > 0 {
		if val, err := room.Coder.Encode(val); err != nil {
			for _, role := range table.Roles {
				role.UnsafeSend(val)
			}
		}
	}
}

// 等待
func gameWait (table *Table) {
	table.State = GameStateWait
	log.Debugf("停止下注:%v", table.CurId)
}

// 准备
func gameReady(table *Table) {
	table.State = GameStateReady
	table.CurId += 1
	log.Debugf("%v准备:%v", gameName, table.CurId)
	table.NewGameRound()
}

// 开始
func gameOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = GameStatePlaying
	log.Debugf("开始下注:%v", table.CurId)
}

func gamePlay(table *Table) {

}


// 发牌结算
func gameDeal(table *Table) {
	table.State = GameStateDeal
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算

	round := table.round
	round.End = room.Now()
	room.SaveLog(round)
}