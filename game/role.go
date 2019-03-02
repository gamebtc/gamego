package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)

// 每个玩家的游戏数据
type Role struct {
	*model.User           // 玩家信息
	session *room.Session // 为空则为机器人
	table   *Table        // 桌子ID
	bill    *msg.GameBill // 输赢情况
	flowSn  int64         // 最后的写分序号，返回时用于验证
	online  bool          // 是否在线
	play    bool          // 是否正在游戏中
}

// 检查投注项位置0:龙赢,1:虎赢,2:和
func (role *Role) RobotCanBet(item int32) bool {
	if role.bill != nil {
		if item == 0 {
			return role.bill.Group[1] == 0
		} else if item == 1 {
			return role.bill.Group[0] == 0
		}
	}
	return true
}

func (role *Role) Reset() {
	if role.bill == nil || role.bill.Bet > 0 {
		role.bill = &msg.GameBill{
			Uid:   role.Id,
			Coin:  role.Coin,
			Group: make([]int64, betItem),
			Job:   role.Job,
		}
	}
}

// 投注
func (role *Role) AddBet(bet msg.BetReq) bool {
	// 检查金币
	coin := int64(bet.Coin)
	if coin <= 0 || coin > role.Coin {
		// 金币不足
		log.Debugf("金币不足:%v,%v", role.Coin, coin)
		return false
	}
	round := role.table.round
	if round == nil {
		log.Debugf("round is null:%v,%v", role.Coin, coin)
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
	}

	// 首次投注
	if bill.Bet == 0 {
		round.Bill = append(round.Bill, bill)
	}

	round.Flow = append(round.Flow, role.Id, bet.Item, bet.Coin)
	round.Group[bet.Item] += coin

	role.Coin -= coin
	bill.Group[bet.Item] += coin
	bill.Bet += coin
	// 有真实玩家下注
	if role.session != nil {
		role.table.round.Real = true
		log.Debugf("玩家下注:%v,%v", role.Id, bet)
	}
	return true
}

// 结算,亚赔方式计算,赔率放大100倍
const radix = 100

func (role *Role) Balance() *model.CoinFlow {
	bill := role.bill
	round := role.table.round
	if bill == nil || bill.Bet <= 0 || round == nil {
		return nil //没有投注
	}
	odds := round.Odds
	prize := int64(0)
	tax := int64(0)
	betIndex := -1 //最后一个有下注的项
	for i := 0; i < betItem; i++ {
		// 下注金额大于0
		if bet := bill.Group[i]; bet > 0 {
			betIndex = i
			//有钱回收,包含输1半
			if odd := int64(odds[i]); odd > -radix {
				w := bet * odd / radix
				if w > 0 {
					// 赢钱了收税，税率按千分比配置，需除以1000
					if r := role.table.dealer.Tax(i); r > 0 {
						tax += w * r / 1000
					}
				}
				//前面是亚赔方式计算,所以需要把本金加上，变为欧赔
				prize += w + bet
			}
		}
	}
	// 扣税后返给玩家的钱
	prize -= tax
	role.Coin += prize
	// 实际输赢要扣掉本金
	addCoin := prize - bill.Bet

	bill.Win = addCoin
	bill.Tax = tax
	round.Tax += tax
	round.Win -= addCoin

	ulog := &msg.FolksUserLog{
		Tab:   role.table.Id,
		Bet:   bill.Bet,
		Group: bill.Group[0 : betIndex+1],
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
