package internal

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol/folks"
	"local.com/abc/game/room"
)

const lastBillCount = 20

// 每个角色的游戏数据
type Role struct {
	model.User             // 玩家信息
	*room.Session          // 发送消息
	table  *Table          // 桌子ID
	bill   *folks.GameBill // 输赢情况
	player *folks.Player   // 玩家信息

	LastBet      [lastBillCount]int64 // 最后20局的下注金额
	LastWin      [lastBillCount]byte  // 最后20局的输赢金额
	LastBetSum   int64
	LastWinCount byte
}

func (role *Role) GetPlayer() *folks.Player {
	role.player.Coin = role.Coin
	return role.player
}

func (role *Role) IsRobot() bool {
	return role.User.Job == model.JobRobot
}

func (role *Role) IsPlayer() bool {
	return role.User.Job == model.JobPlayer
}

// 检查投注项位置
func (role *Role) RobotCanBet(item int32, bet int32) bool {
	if int64(bet) > role.Coin || role.Coin < room.Config.PlayMin {
		return false
	}
	if role.bill != nil && item < int32(len(badBet)) {
		if role.bill.Group[badBet[item]] > 0 {
			return false
		}
	}
	return true
}

func (role *Role) Reset() {
	role.bill = nil
}

// 存在指定的下注金额
func ExistsBetItem(bet int32) bool {
	return (bet >= betItems[0]) && (bet <= betItems[len(betItems)-1]) && (bet%100 == 0)
}

// 投注
func (role *Role) AddBet(req folks.BetReq) error {
	// 检查游戏状态
	round := role.table.round
	if round == nil || role.table.State != GameStatePlaying {
		return errors.New("本局游戏已停止下注，请您稍后再试！")
	}

	// 检查投注项
	i := req.Item
	bet := int64(req.Bet)
	if i >= int32(betItemCount) || bet <= 0 {
		return errors.New(fmt.Sprintf("错误的投注项:%v,%v", i, bet))
	}

	if !ExistsBetItem(req.Bet) {
		return errors.New(fmt.Sprintf("错误的投注金额:%v|%v", i, bet))
	}

	// 检查金币
	need := bet
	if multipleLost > 1 {
		totalBet := int64(0)
		if role.bill != nil {
			totalBet = role.bill.Bet
		}
		need = multipleLost*(totalBet+bet) - totalBet
	}
	if need > role.Coin {
		return errors.New(fmt.Sprintf("余额不足%v金币，请您充值！", need))
	}

	if role.Coin < room.Config.PlayMin {
		return errors.New(fmt.Sprintf("余额%v金币以上才可以下注，请您充值！", room.Config.PlayMin))
	}

	bill := role.bill
	if bill == nil {
		bill = &folks.GameBill{
			Uid:   role.Id,
			Coin:  role.Coin,
			Group: make([]int64, betItemCount),
			Job:   role.User.Job,
		}
		role.bill = bill
		// 首次投注
		round.Bill = append(round.Bill, bill)
	}

	round.Flow = append(round.Flow, role.Id, i, req.Bet)
	round.Group[i] += bet

	role.Coin -= bet
	bill.Group[i] += bet
	bill.Bet += bet
	// 有真实玩家下注
	if role.IsPlayer() {
		if round.UserBet == nil {
			round.UserBet = make([]int64, betItemCount)
		}
		round.UserBet[i] += bet
		log.Debugf("%v下注:%v_%v", role.Id, i, bet/100)
	}
	return nil
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

	i := role.TotalRound % lastBillCount
	role.LastBetSum += bet - role.LastBet[i]
	role.LastBet[i] = bet

	if role.LastWin[i] == 1 {
		role.LastWinCount--
	}
	if addCoin >= 0 {
		role.LastWin[i] = 1
		role.LastWinCount++
	} else {
		role.LastWin[i] = 0
	}

	role.AddWinBet(addCoin, bet)

	bill.Tax = tax
	bill.Win = addCoin

	if role.IsRobot() {
		return nil
	}

	if role.IsPlayer() {
		round.Tax += tax
		round.Win -= addCoin
	}

	ulog := &folks.UserLog{
		Tab:   role.table.Id,
		Bet:   bet,
		Group: room.TrimEndZero(bill.Group),
		Log:   round.Id,
		Poker: round.Poker,
	}

	// 写分
	role.FlowSn = room.NewSn(1)
	flow := &model.CoinFlow{
		Sn:     role.FlowSn,
		Uid:    role.Id,
		Add:    addCoin,
		Tax:    tax,
		Expect: role.Coin,
		Room:   room.RoomId,
		Kind:   room.KindId,
		Type:   1,
		Note:   round.Note,
		Att:    ulog,
	}
	return flow
}
