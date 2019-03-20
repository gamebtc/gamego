package internal

import (
	"time"

	log "github.com/sirupsen/logrus"
	//"local.com/abc/game/protocol/folks"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

//type GameDriver interface {
//	// 准备游戏, 状态1
//	Ready(table *Table)
//	// 开始下注, 状态2
//	Open(table *Table)
//    // 游戏中, 状态2
//    Play(table *Table)
//	// 停止下注, 状态3
//	Stop(table *Table)
//	// 发牌结算, 状态4
//	Deal(table *Table)
//}

const chairCount = 5

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	Id        int32             // 桌子ID
	CurId     int32             // 当前局的ID
	State     int32             // 0:等待玩家进入;1:等待玩家同意;2:游戏中;3:结算
	Roles     [chairCount]*Role // 所有游戏玩家,按位置
	Player    []*Role           // 游戏中的玩家
	round     *GameRound        // 1局
	continued int               // 持续秒数
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
				State: table.State,
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
	gameReady(table)
}

func (table *Table) Update() {
	table.continued--
	switch table.State {
	case 0:
		if table.continued == 0 {
			table.continued = 10
			gameOpen(table)
		}
	case 2:
		if table.continued == 0 {
			table.continued = 1
			gameStop(table)
		} else {
			gamePlay(table)
		}
	case 3:
		if table.continued == 0 {
			table.continued = 2
			gameDeal(table)
		}
	case 4:
		if room.Config.Pause == 0 {
			if table.continued == 0 {
				table.continued = 1
				gameReady(table)
			}
		} else {
			// 暂停
			table.continued++
			log.Debugf("Pause: %v", time.Now())
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

// 准备
func gameReady(table *Table) {
	table.State = 1
	table.CurId += 1
	log.Debugf("%v准备:%v", gameName, table.CurId)
	table.NewGameRound()
}

// 开始
func gameOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("开始下注:%v", table.CurId)

}

func gamePlay(table *Table) {
}

// 停止下注
func gameStop (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
}

// 发牌结算
func gameDeal(table *Table) {
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算


	round := table.round
	round.End = room.Now()
	room.SaveLog(round)

}