package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)


type Sender interface {
    Send(val interface{})bool
}

// 每个玩家的游戏数据
type Role struct {
	*model.User            // 玩家信息
	Sender                 // 发送消息
	Online   bool          // 是否在线
	TotalWin int64         // 进入之后的赢钱总数
	Coin     int64         // 当前房间使用的币
	table    *Table        // 桌子ID
	bill     *msg.GameBill // 输赢情况
	flowSn   int64         // 最后的写分序号，返回时用于验证
}

func (role *Role)GetMsgUser() *msg.User {
	//return &this.User
	return &msg.User{
		Id:    role.Id,
		Icon:  role.Icon,
		Vip:   role.Vip,
		Name:  role.Name,
		Coin:  role.Coin,
	}
}

func(role *Role)IsRobot()bool {
	return role.User.Job == 10
}

// 检查投注项位置
func (role *Role) RobotCanBet(item int32) bool {
	if role.bill != nil {
		for k, v := range role.bill.Group {
			if k != int(item) && v > int64(0) {
				return false
			}
		}
	}
	return true
}

func (role *Role) Reset() {
	role.bill = nil
}

// 投注
func (role *Role) AddBet(bet msg.BetReq) bool {
	// 检查投注项
	i := bet.Item
	if i >= int32(betItem) {
		log.Debugf("投注项错误:%v,%v", i, betItem)
		return false
	}
	// 检查金币
	coin := int64(bet.Coin)
	if coin <= 0 || coin > role.Coin {
		log.Debugf("金币不足:%v,%v", role.Coin, coin)
		return false
	}

	round := role.table.round
	if round == nil || role.table.State != 2 {
		log.Debugf("下游已停止:%v,%v", role.Coin, coin)
		return false
	}

	bill := role.bill
	if bill == nil {
		bill = &msg.GameBill{
			Uid:   role.Id,
			Coin:  role.Coin,
			Group: make([]int64, betItem),
			Job:   role.Job,
		}
		role.bill = bill
		// 首次投注
		round.Bill = append(round.Bill, bill)
	}

	round.Flow = append(round.Flow, role.Id, i, bet.Coin)
	round.Group[i] += coin

	role.Coin -= coin
	bill.Group[i] += coin
	bill.Bet += coin
	// 有真实玩家下注
	if !role.IsRobot(){
		round.Real = true
		round.UserGroup[i] += coin
		log.Debugf("玩家下注:%v,%v", role.Id, bet)
	}
	return true
}

// 结算,欧赔方式计算,赔率放大100倍
const radix = 100
const lostRadix = 0

func (role *Role) Balance() *model.CoinFlow {
	bill := role.bill
	round := role.table.round
	if bill == nil || bill.Bet <= 0 || round == nil {
		return nil //没有投注
	}

	prize, tax, bet := Balance(bill.Group, round.Odds)
	if bet != bill.Bet {
		log.Errorf("Balance error:uid:%v,bet1:%v,bet2:%v", role.Id, bet, bill.Bet)
	}

	// 扣税后返给玩家的钱
	prize -= tax
	role.Coin += prize
	// 实际输赢要扣掉本金(用于写分)
	addCoin := prize - bet
	if addCoin > 0 {
		role.TotalWin += addCoin
	}

	bill.Win = addCoin
	bill.Tax = tax
	round.Tax += tax
	round.Win -= addCoin

	ulog := &msg.FolksUserLog{
		Tab:   role.table.Id,
		Bet:   bet,
		Group: TrimEndZero(bill.Group),
		Log:   round.Id,
		Poker: round.Poker,
	}

	// 写分
	role.flowSn = room.NewSn(1)
	flow := &model.CoinFlow{
		Sn:     role.flowSn,
		Uid:    role.Id,
		Add:    addCoin,
		Tax:    tax,
		Expect: role.Coin,
		Room:   room.RoomId,
		Kind:   room.KindId,
		Type:   1,
		Note:   round.Note + fmt.Sprintf("%v", ulog.Group),
		Att:    ulog,
	}
	return flow
}
