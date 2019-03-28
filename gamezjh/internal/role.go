package internal

import (
	//"errors"
	//"fmt"
	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

var (
	dealer = model.NewGoldenFlowerDealer(true)
	goldPool int64 = 0
)

// 每个角色的游戏数据
type Role struct {
	model.User                  // 玩家信息
	*room.Session               // 发送消息
	table         *Table        // 桌子ID
	bill          *zjh.GameBill // 输赢情况
	player        *zjh.Player   // 玩家信息
}

func (role *Role) GetPlayer() *zjh.Player {
	return role.player
}

func (role *Role) IsRobot() bool {
	return role.User.Job == model.JobRobot
}

func (role *Role) IsPlayer() bool {
	return role.User.Job == model.JobPlayer
}

//func (role *Role) Reset() {
//	role.bill = &zjh.GameBill{
//		Uid:  role.User.Id,
//		Job:  role.User.Job,
//		Coin: role.Coin,
//	}
//
//	role.player.State = zjh.Player_None
//	role.player.Look = false
//	role.player.Bet = 0
//	role.player.Down = 0
//	role.player.Coin = role.Coin
//}

func ChangeGoldPool(add int64){
	goldPool+=add
}


// 准备
func (role *Role) Ready() {
	player := role.player
	if player.State == zjh.Player_None {
		player.State = zjh.Player_Ready
		// 发送准备信息
		role.table.SendToAll(&zjh.ActionAck{})
	}
}

// 看牌
func (role *Role) Look() {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	if player.Bet <= role.table.round.Ante {
		// 没有下注不能看牌
		return
	}
	if player.Look == false {
		player.Look = true
	}
	// 发送消息给自己
	if !role.IsRobot() {
		role.UnsafeSend(&zjh.ActionAck{
			Uid:   role.User.Id,
			Type:  zjh.ActionType_ActionLook,
			Poker: role.bill.Poker,
		})
	}
	// 发送消息给其它人
	ack := &zjh.ActionAck{
		Uid:  role.User.Id,
		Type: zjh.ActionType_ActionLook,
	}
	role.table.SendToOther(ack, role)
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionLook,
	}
	role.table.round.Log = append(role.table.round.Log, log)
}

// 输了，结算
func (role *Role) Lose() {
	if role.IsRobot() {
		// 机器人输了不用写分直接返回
		return
	}

	bill := role.bill
	bet := bill.Bet
	round := role.table.round
	round.Win += bet
	// 写分
	role.FlowSn = room.NewKindSn()
	poker := bill.Poker
	flow := &model.CoinFlow{
		Sn:    role.FlowSn,
		Uid:   role.Id,
		Add:   -bet,
		New:   role.Coin,
		Old:   bill.Coin,
		Tax:   0,
		Room:  room.RoomId,
		Kind:  room.KindId,
		Note:  model.PokerArrayString(poker),
		Bet:   bet,
		LogId: round.Id,
		Poker: poker,
	}
	if room.WriteCoin(flow) == nil {
		if role.Coin != flow.New {
			log.Warnf("金币变化:%v-%v", flow.New, role.Coin)
		}
	}
	log.Debugf("结算:%v", flow)
}

// 赢了，结算
func (role *Role) Win(prize int64) {
	// 下注金额
	bill := role.bill
	bet := bill.Bet
	winCoin := prize - bet
	tax := int64(0)
	water := int64(0) //抽水进入彩金池
	if winCoin > 0 {
		const poolTax= 10 // 千分之10进入彩金池
		water = winCoin * poolTax / 1000
		tax = winCoin * room.Config.Tax / 1000
	}
	round := role.table.round
	// TODO:检查是否有彩金,彩金不交税
	lucky := int64(0)
	ag := dealer.GetGroup(bill.Poker)
	if ag.IsThreeKind() {
		// 豹子5%
		lucky = round.Pool * 5 / 100
	} else if ag.IsStraightFlush() {
		// 顺金1%
		lucky = round.Pool * 1 / 100
	}
	// 扣税后返给玩家的币
	prize = prize - tax - water + lucky
	role.Coin += prize
	role.player.Coin = role.Coin
	bill.Tax = tax
	bill.Water = water
	bill.Lucky = lucky

	if role.IsRobot() {
		// 机器人赢了不用写分直接返回
		return
	}

	// 实际输赢要扣掉本金(用于写分)
	addCoin := prize - bet
	round.Tax += tax
	round.Water += water
	round.Lucky += lucky
	round.Win -= addCoin
	// 写分
	role.FlowSn = room.NewKindSn()
	poker := bill.Poker
	flow := &model.CoinFlow{
		Sn:    role.FlowSn,
		Uid:   role.Id,
		Add:   addCoin,
		New:   role.Coin,
		Old:   bill.Coin,
		Tax:   tax,
		Room:  room.RoomId,
		Kind:  room.KindId,
		Note:  model.PokerArrayString(poker),
		Bet:   bet,
		LogId: round.Id,
		Poker: poker,
	}
	if room.WriteCoin(flow) == nil {
		if role.Coin != flow.New {
			log.Warnf("金币变化:%v-%v", flow.New, role.Coin)
		}
	}
	log.Debugf("结算:%v", flow)
}

// 减金币
func (role *Role) decCoin(bet int64) {
	role.Coin -= bet
	role.player.Coin = role.Coin
	role.player.Bet += bet
	role.bill.Bet += bet
	role.table.round.Sum += bet
}

// 弃牌
func (role *Role) Discard(overtime bool) {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}
	player.State = zjh.Player_Discard
	role.Lose()

	table := role.table
	// 发送消息所有人
	table.SendToAll(&zjh.ActionAck{
		Uid:  role.User.Id,
		Type: zjh.ActionType_ActionDiscard,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionDiscard,
	}
	table.round.Log = append(table.round.Log, log)

	// 已经有人胜利
	if table.TryDeal() {
		return
	}

	// 轮到自己下注时放弃
	if player.Chair == table.waitChair {
		table.NextWait()
	}
}

func (role *Role) pk(opp *Role) bool {
	ag := dealer.GetGroup(role.bill.Poker)
	bg := dealer.GetGroup(opp.bill.Poker)
	win := ag.Power > bg.Power
	if win {
		opp.player.State = zjh.Player_Lose
		opp.Lose()
	} else {
		role.player.State = zjh.Player_Lose
		role.Lose()
	}
	return win
}

// 是否是PK过的对手
func(role *Role) IsOpponent(id int32)bool {
	if role.bill != nil && role.bill.Pk != nil {
		for _, opp := range role.bill.Pk {
			if opp == id {
				return true
			}
		}
	}
	return false
}

// 比牌
func (role *Role) Compare(opponent int32) {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	table := role.table

	// 3轮后到自己才可以比牌
	if table.round.Ring < 3 || player.Chair != table.waitChair {
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}
	bet := table.curAnte * rate
	if role.Coin < int64(bet) {
		// 钱不够
		return
	}

	opp := table.FindRole(opponent)
	if opp == nil || opp == role {
		return
	}

	// 如果PK时对手已弃牌，转换为下注
	if opp.player.State != zjh.Player_Playing {
		role.AddBet(bet)
		return
	}

	role.decCoin(int64(bet))

	// PK
	role.bill.Pk = append(role.bill.Pk, opp.User.Id)
	opp.bill.Pk = append(opp.bill.Pk, role.User.Id)
	win := role.pk(opp)

	table.SendToAll(&zjh.ActionAck{
		Uid:      role.User.Id,
		Type:     zjh.ActionType_ActionCompare,
		Bet:      bet,
		Opponent: opponent,
		Win:      win,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start:    room.Now(),
		Uid:      role.User.Id,
		Type:     zjh.ActionType_ActionCompare,
		Bet:      bet,
		Opponent: opponent,
		Win:      win,
	}
	table.round.Log = append(table.round.Log, log)

	// 已经有人胜利
	if table.TryDeal() {
		return
	}

	table.NextWait()
}

// 下注(跟注+加注)
func (role *Role) AddBet(bet int32) {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}
	if player.Chair != role.table.waitChair {
		return
	}

	table := role.table
	// 转为全压
	if table.allinBet > 0 {
		role.Allin()
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}

	minBet := table.curAnte * rate
	if bet == 0 {
		// 跟注时将下注额设置为当前最小底注
		bet = minBet
	}
	if bet < minBet {
		// 下注错误
		return
	}

	if role.Coin < int64(bet)+100 {
		// 下注后至少保留1块钱，钱不够
		return
	}

	if bet > minBet {
		// 加注, 检查下注项
		newAnte := bet / rate
		if ExistsBetItem(newAnte) == false {
			return
		}
		table.curAnte = newAnte
	}

	role.decCoin(int64(bet))

	table.SendToAll(&zjh.ActionAck{
		Uid:  role.User.Id,
		Type: zjh.ActionType_ActionAddBet,
		Bet:  bet,
		Coin: role.Coin,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionAddBet,
		Bet:   bet,
	}
	table.round.Log = append(table.round.Log, log)

	table.NextWait()
}

// 全压(3轮后,有玩家的钱小于最小下注金额时开启)
func (role *Role) Allin() {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	table := role.table
	// 3轮后到自己才可以全压
	if table.round.Ring < 3 || player.Chair != table.waitChair {
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}

	// 第一次全压，底注设置为带钱最少的人
	if table.allinBet == 0 {
		allinBet := (role.Coin - (role.Coin % 100)) / int64(rate)
		for _, item := range table.Roles {
			if item != nil && item != role {
				coin := (item.Coin - (item.Coin % 100)) / 2
				if coin > 0 && coin < allinBet {
					allinBet = coin
				}
			}
		}
		table.allinBet = int32(allinBet)
	}
	if table.allinBet == 0 {
		return
	}

	bet := table.allinBet * rate
	if role.Coin < int64(bet) {
		// 钱不够
		return
	}

	player.State = zjh.Player_Allin

	role.decCoin(int64(bet))

	table.SendToAll(&zjh.ActionAck{
		Uid:  role.User.Id,
		Type: zjh.ActionType_ActionAllin,
		Bet:  bet,
		Coin: role.Coin,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionAllin,
		Bet:   bet,
	}
	table.round.Log = append(table.round.Log, log)
	// 自动比牌
	table.TryAutoComp()

	// 已经有人胜利
	if table.TryDeal() {
		return
	}
	table.NextWait()
}

// 换桌玩
func (role *Role) RenewDesk() {
	player := role.player
	if player.State == zjh.Player_Playing {
		return
	}
}

// 存在指定的下注金币
func ExistsBetItem(bet int32) bool {
	for _, item := range betItems {
		if item == bet {
			return true
		}
	}
	return false
}
