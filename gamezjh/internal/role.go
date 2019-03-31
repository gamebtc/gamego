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
	dealer   = model.NewGoldenFlowerDealer(false)
	goldPool = int64(0)
)

type Player struct {
}

// 每个角色的游戏数据
type Role struct {
	model.User                  // 玩家信息
	*room.Session               // 发送消息
	table         *Table        // 桌子ID
	bill          *zjh.GameBill // 输赢情况
	player        *zjh.Player   // 玩家信息
	robot         Robot         // 机器人行为
	poker         *model.GoldenFlowerGroup
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

func ChangeGoldPool(add int64) {
	goldPool += add
}

// 输了，结算
func (role *Role) Lose() {
	player := role.player
	player.State = zjh.Player_Lose

	table := role.table
	bill := role.bill
	bet := bill.Bet
	addCoin := -bet
	bill.Win = addCoin
	role.AddWinBet(addCoin, bet)
	log.Debugf("[%v]lose:%v-%v", table.Id, player.Id, addCoin)

	if role.IsRobot() {
		// 机器人输了不用写分直接返回
		return
	}

	round := table.round
	round.Win -= addCoin
	// 写分
	role.FlowSn = room.NewKindSn()
	poker := bill.Poker
	flow := &model.CoinFlow{
		Sn:    role.FlowSn,
		Uid:   role.Id,
		Add:   addCoin,
		New:   role.Coin,
		Old:   bill.OldCoin,
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
	log.Debugf("[%v]结算:%v", table.Id, flow)
}

// 赢了，结算
const poolTax = 10 // 千分之10进入彩金池
func (role *Role) Win(prize int64) {
	player := role.player
	player.State = zjh.Player_Win

	table := role.table
	// 下注金额
	bill := role.bill
	bet := bill.Bet
	winCoin := prize - bet
	tax := int64(0)
	water := int64(0) //抽水进入彩金池
	if winCoin > 0 {
		water = winCoin * poolTax / 1000
		tax = winCoin * room.Config.Tax / 1000
	}
	round := role.table.round
	// TODO:检查是否有彩金,彩金不交税
	lucky := int64(0)
	if role.poker.IsThreeKind() {
		// 豹子5%
		lucky = round.Pool * 5 / 100
		log.Debugf("[%v]lucky:%v-%v-%v", table.Id, model.PokerArrayString(bill.Poker), player.Id, lucky)
	} else if role.poker.IsStraightFlush() {
		// 顺金1%
		lucky = round.Pool * 1 / 100
		log.Debugf("[%v]lucky:%v-%v-%v", table.Id, model.PokerArrayString(bill.Poker), player.Id, lucky)
	}
	// 扣税后返给玩家的币
	prize = prize - tax - water + lucky
	role.Coin += prize
	player.Coin = role.Coin
	// 实际输赢要扣掉本金(用于写分)
	addCoin := prize - bet
	bill.Tax = tax
	bill.Win = addCoin
	bill.Water = water
	bill.Lucky = lucky
	role.AddWinBet(addCoin, bet)
	log.Debugf("[%v]win:%v-%v", table.Id, player.Id, addCoin)

	// 更新全局池子
	if lucky > 0 {
		ChangeGoldPool(-lucky)
	}
	if water > 0 {
		ChangeGoldPool(water)
	}

	if role.IsRobot() {
		// 机器人赢了不用写分直接返回
		return
	}
	round.Tax += tax
	round.Win -= addCoin
	round.Water += water
	round.Lucky += lucky
	// 写分
	role.FlowSn = room.NewKindSn()
	poker := bill.Poker
	flow := &model.CoinFlow{
		Sn:    role.FlowSn,
		Uid:   role.Id,
		Add:   addCoin,
		New:   role.Coin,
		Old:   bill.OldCoin,
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
			log.Warnf("[%v]金币变化:%v-%v", table.Id, flow.New, role.Coin)
		}
	}
	log.Debugf("[%v]结算:%v", table.Id, flow)
}

// 减金币
func (role *Role) decCoin(bet int64) {
	role.Coin -= bet
	role.player.Coin = role.Coin
	role.player.Bet += bet
	role.bill.Bet += bet
	role.table.round.Sum += bet
}

// 新的一局
func(role *Role)newGameRound(poker []byte)*zjh.GameBill {
	bill := &zjh.GameBill{
		Uid:     role.User.Id,
		Job:     role.User.Job,
		OldCoin: role.Coin,
		Poker:   poker,
	}
	role.bill = bill
	role.player.State = zjh.Player_Playing
	role.player.Look = false
	role.player.Down = 0
	role.poker = dealer.GetGroup(poker)
	role.bill = bill
	return bill
}

// 准备
func (role *Role) Ready() {
	player := role.player
	if player.State != zjh.Player_None {
		return
	}
	player.State = zjh.Player_Ready
	// 发送准备信息
	role.table.SendToAll(&zjh.ActionAck{
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionReady,
	})
}

// 看牌
func (role *Role) Look() {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	table := role.table
	round := table.round
	// 第一轮还没有下注时不能看牌
	if player.Bet <= room.Config.Ante {
		return
	}

	// 发送消息给自己
	if !role.IsRobot() {
		role.UnsafeSend(&zjh.ActionAck{
			Uid:   role.User.Id,
			Type:  zjh.ActionType_ActionLook,
			Poker: role.bill.Poker,
		})
	}

	if player.Look {
		// 重复的看牌命令
		return
	}

	log.Debugf("[%v]look:%v", table.Id, player.Id)
	player.Look = true
	// 发送消息给其它人
	ack := &zjh.ActionAck{
		Uid:  role.User.Id,
		Type: zjh.ActionType_ActionLook,
	}
	table.SendToOther(ack, role)
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  zjh.ActionType_ActionLook,
	}
	round.Log = append(round.Log, log)
}

// 弃牌
func (role *Role) Discard(overtime bool) {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	table := role.table
	log.Debugf("[%v]dis:%v-%v", table.Id, player.Id, overtime)

	player.State = zjh.Player_Discard
	role.Lose()

	actionType := zjh.ActionType_ActionDiscard
	if overtime {
		actionType = zjh.ActionType_ActionOvertime
	}
	// 发送消息所有人
	table.SendToAll(&zjh.ActionAck{
		Uid:  role.User.Id,
		Type: actionType,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start: room.Now(),
		Uid:   role.User.Id,
		Type:  actionType,
	}
	table.round.Log = append(table.round.Log, log)

	// 已经有人胜利
	if table.tryCheckout() {
		return
	}

	// 轮到自己下注时放弃
	if table.isRunner(player.Id) {
		table.nextWait()
	}
}

// 是否是PK过的对手
func (role *Role) IsOpponent(id int32) bool {
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
	if table.round.Ring < 3 {
		return
	}
	if !table.isRunner(player.Id) {
		return
	}
	if table.firstAllin != nil {
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}

	round := table.round
	bet := round.Ante * rate
	if role.Coin < int64(bet) {
		// 钱不够
		return
	}

	opp := table.findRole(opponent)
	if opp == nil || opp == role {
		return
	}

	// 如果PK时对手已弃牌，转换为下注
	if opp.player.State != zjh.Player_Playing {
		role.AddBet(bet)
		return
	}

	log.Debugf("[%v]vs%v-%v", table.Id, player.Id, opp.User.Id)

	role.decCoin(int64(bet))

	// PK
	role.bill.Pk = append(role.bill.Pk, opp.User.Id)
	opp.bill.Pk = append(opp.bill.Pk, role.User.Id)
	var winner []int32
	if role.poker.Power > opp.poker.Power {
		opp.player.State = zjh.Player_Lose
		opp.Lose()
		winner = []int32{role.User.Id}
	} else {
		player.State = zjh.Player_Lose
		role.Lose()
		winner = []int32{opp.User.Id}
	}

	players := []int32{opponent}
	table.SendToAll(&zjh.ActionAck{
		Uid:     role.User.Id,
		Type:    zjh.ActionType_ActionCompare,
		Bet:     bet,
		Players: players,
		Winners: winner,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start:   room.Now(),
		Uid:     role.User.Id,
		Type:    zjh.ActionType_ActionCompare,
		Bet:     bet,
		Players: players,
		Winners: winner,
	}
	round.Log = append(round.Log, log)

	// 已经有人胜利
	if table.tryCheckout() {
		return
	}
	table.nextWait()
}

// 下注(跟注+加注)
func (role *Role) AddBet(bet int32) {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}
	table := role.table
	if !table.isRunner(player.Id) {
		return
	}

	// 转为全压
	if table.firstAllin != nil {
		role.Allin()
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}

	round := table.round
	minBet := round.Ante * rate
	if bet == 0 {
		// 跟注时将下注额设置为当前最小底注
		bet = minBet
	}
	if bet < minBet {
		// 下注错误
		return
	}

	if role.Coin < int64(bet) {
		return
	}

	if bet > minBet {
		// 加注, 检查下注项
		newAnte := bet / rate
		if ExistsBetItem(newAnte) == false {
			return
		}
		round.Ante = newAnte
	}

	log.Debugf("[%v]bet:%v-%v", table.Id, player.Id, bet)

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
	round.Log = append(round.Log, log)

	table.nextWait()
}

// 全压(3轮后,有玩家的钱小于最小下注金额时开启)
func (role *Role) Allin() {
	player := role.player
	if player.State != zjh.Player_Playing {
		return
	}

	table := role.table
	// 3轮后到自己才可以全压
	if table.round.Ring < 3 {
		return
	}
	if !table.isRunner(player.Id) {
		return
	}

	rate := int32(1)
	if player.Look {
		rate = 2
	}

	round := table.round
	// 第一次全压，底注设置为带钱最少的人
	if table.firstAllin == nil {
		table.firstAllin = role
		allinBet := (role.Coin - (role.Coin % 100)) / int64(rate)
		for _, other := range table.players {
			if other.State == zjh.Player_Playing && other != player {
				coin := (other.Coin - (other.Coin % 100)) / 2
				if coin > 0 && coin < allinBet {
					allinBet = coin
				}
			}
		}
		round.Ante = int32(allinBet)
	}

	bet := round.Ante * rate
	if role.Coin < int64(bet) {
		// 钱不够
		log.Debugf("[%v]allin err:%v-%v-%v", table.Id, player.Id, bet, role.Coin)
		return
	}
	log.Debugf("[%v]all:%v-%v", table.Id, player.Id, bet)

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
	round.Log = append(round.Log, log)
	// 自动比牌
	if table.tryCheckout() {
		return
	}
	table.nextWait()
}

// 换桌玩
func (role *Role) RenewDesk() {
	player := role.player
	if player.State == zjh.Player_Playing || player.State == zjh.Player_Allin{
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
