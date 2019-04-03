package internal

import (
	"math/rand"
	"strings"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol/zjh"
	"local.com/abc/game/room"
)

const chairCount = 5

type GameState int32

const (
	// 无
	GameStateNone GameState = 0
	// 准备游戏
	GameStateReady GameState = 1
	// 游戏中
	GameStatePlaying GameState = 2
	// 结算
	GameStateCheckout GameState = 3
)

const (
	// 准备时间，15秒
	readySecond = 15
	// 等待超时3秒
	waitSecond = 20
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
	Roles      [chairCount]*Role // 所有游戏玩家,按位置,玩家离开后Role退出
	round      *GameRound        // 1局
	waitSecond int32             // 等待秒数
	poker      []byte            // 所有的牌
	playIndex  int               // 游戏指针
	players    []*Player         // 游戏中的玩家
	playerInfo []*zjh.Player     // 玩家信息
	winner     []*Role           // 最后的赢家
	firstAllin *Role             // 第一个全压的人
	playCount  int32             // 剩余正在玩的人数
	lookCount  int32             // 已看牌的人数
	compCount  int32             // 被PK掉的玩家数
}

func NewTable() *Table {
	t := &Table{}
	return t
}

func (table *Table) reset() {
	table.round = nil
	table.waitSecond = 0
	table.poker = nil
	table.playIndex = 0
	table.players = nil
	table.playerInfo = nil
	table.winner = nil
	table.firstAllin = nil
}

func (table *Table) mustWin() bool {
	return room.Config.WinRate > gameRand.Int31n(1000)
}

// 等待发命令者
func (table *Table) runner() (player *Player) {
	return table.players[table.playIndex]
}

// 等待发命令者
func (table *Table) isRunner(uid int32) bool {
	return uid == table.players[table.playIndex].Player.Id
}

func (table *Table) findRole(id int32) *Role {
	for _, role := range table.Roles {
		if role != nil && role.Id == id {
			return role
		}
	}
	return nil
}

// 试着结算
func (table *Table) tryCheckout() bool {
	if table.winner != nil {
		return true
	}

	// Allin
	if table.firstAllin != nil {
		for _, player := range table.players {
			if player.State == zjh.Player_Playing {
				return false
			}
		}
		table.autoComp(true)
		return true
	}

	var winner *Role
	// 只有一个玩家为Playing时，玩家胜利
	for _, player := range table.players {
		if player.State == zjh.Player_Playing {
			if winner != nil {
				return false
			}
			winner = table.Roles[player.Chair]
		}
	}
	if winner == nil {
		return false
	}
	table.winner = []*Role{winner}
	// 给赢家结算
	round := table.round
	prize := round.Sum
	winner.Win(prize)

	round.winner = []int32{winner.User.Id}
	round.prize = []int64{prize}
	table.gameClose()
	return true
}

// 进行PK，选出所有获胜的玩家
func (table *Table) selectWinners(allinRoles []*Role, allIn bool) (winners []*Role, loses []*Role) {
	role := allinRoles[0]
	winners = make([]*Role, 0, 3)
	loses = make([]*Role, 0, 3)
	for _, opp := range allinRoles {
		// 将每一个最后的玩家都加入PK列表
		for _, b := range allinRoles {
			if opp != b {
				opp.Bill.Pk = append(opp.Bill.Pk, b.User.Id)
			}
		}
		if role != opp {
			if role.Poker.Power > opp.Poker.Power {
				loses = append(loses, opp)
			} else if role.Poker.Power < opp.Poker.Power {
				loses = append(loses, role)
				role = opp
				winners = winners[:0]
			} else {
				winners = append(winners, opp)
			}
		}
	}
	winners = append(winners, role)
	winCount := len(winners)
	if allIn {
		// 如果有多个获胜者，去掉第一个全押的人
		if len(winners) > 1 {
			for i := 0; i < winCount; i++ {
				if role := winners[i]; role == table.firstAllin {
					winCount--
					winners = append(winners[:i], winners[i+1:]...)
					loses = append(loses, role)
					break
				}
			}
		}
	}
	return
}

// 达到指定下注轮数或者下一轮有玩家金币不足时自动开牌
func (table *Table) autoComp(allIn bool) {
	var playerState zjh.Player_State
	var actionType zjh.ActionType
	if allIn {
		// 全压时选择所有全压的玩家
		playerState = zjh.Player_Allin
		actionType = zjh.ActionType_ActionAllinCompare
	} else {
		// 自动开牌选择所有游戏中的玩家
		playerState = zjh.Player_Playing
		actionType = zjh.ActionType_ActionAutoCompare
	}

	// 选出最后进行PK的所有玩家
	allinRoles := make([]*Role, 0, 3)
	players := make([]int32, 0, 3)
	for _, player := range table.players {
		if player.State == playerState {
			players = append(players, player.Id)
			allinRoles = append(allinRoles, table.Roles[player.Chair])
		}
	}

	winners, loses := table.selectWinners(allinRoles, allIn)
	table.winner = winners

	// 结算输家
	for _, role := range loses {
		role.Lose()
	}

	// 结算赢家
	winCount := len(winners)
	round := table.round
	prize := round.Sum / int64(winCount)
	round.winner = make([]int32, winCount)
	round.prize = make([]int64, winCount)
	for i, winner := range winners {
		winner.Win(prize)
		round.winner[i] = winner.User.Id
		round.prize[i] = prize
	}

	table.SendToAll(&zjh.ActionAck{
		Type:    actionType,
		Players: players,
		Winners: round.winner,
	})
	// 添加日志
	log := &zjh.ActionLog{
		Start:   room.Now(),
		Type:    actionType,
		Players: players,
		Winners: round.winner,
	}
	table.round.Log = append(table.round.Log, log)

	table.gameClose()
}

// 进入下一轮
func (table *Table) nextRing() bool {
	// 最多10轮
	const maxRing = 10
	round := table.round
	round.Ring++
	log.Debugf("[%v]ring:%v", table.Id, round.Ring)
	if table.firstAllin == nil {
		// 达到最大轮数限制
		if round.Ring > maxRing {
			table.autoComp(false)
			return false
		}
		// 下一轮玩家存在金币不足的情况
		for _, player := range table.players {
			if player.State == zjh.Player_Playing && player.Coin < maxBetItem {
				table.autoComp(false)
				return false
			}
		}
	}
	return true
}

// 下一个等待者
func (table *Table) nextWait() bool {
	if table.winner != nil {
		return false
	}
	playCount := len(table.players)
	for i := 1; i < playCount; i++ {
		table.playIndex++
		if table.playIndex == playCount {
			table.playIndex = 0
			if table.nextRing() == false {
				return false
			}
		}
		if player := table.runner(); player.State == zjh.Player_Playing {
			//log.Debugf("wait player:%v, %v", table.ring, player.Id)
			player.Down = waitSecond
			return true
		}
	}
	return false
}

// 查找一个空位
func (table *Table) findEmptyChair() int32 {
	for i, v := range table.Roles {
		if v == nil {
			return int32(i)
		}
	}
	return -1
}

// 增加玩家
func (table *Table) addRole(role *Role) bool {
	i := table.findEmptyChair()
	if i < 0 {
		return false
	}
	role.Player = &Player{
		Player: zjh.Player{
			Id:    role.Id,
			Icon:  role.Icon,
			Vip:   role.Vip,
			Name:  role.Name,
			Coin:  role.Coin,
			State: zjh.Player_Ready,
			Chair: i,
		},
	}
	role.table = table
	table.Roles[i] = role
	table.RoleCount++
	if !role.IsRobot() {
		role.Player.Robot = &RobotAi{}
		table.sendGameInit(role)
	}else{
		role.Player.Robot = &RobotAi{}
	}
	log.Debugf("[%v]addRole:%v", table.Id, role.Id)
	return true
}

func (table *Table) removeRole(i int) bool {
	role := table.Roles[i]
	if role == nil {
		return false
	}
	table.RoleCount--
	table.Roles[i] = nil
	role.table = nil
	role.Player = nil
	log.Debugf("[%v]removeRole:%v", table.Id, role.Id)
	return true
}

// 所有机器人退出
func (table *Table) freeRobots() {
	for i, role := range table.Roles {
		if role != nil && role.IsRobot() {
			table.removeRole(i)
			PutWaitRobot(role)
		}
	}
}

// 初始化场景
func (table *Table) sendGameInit(role *Role) {
	if !role.IsRobot() {
		ack := &zjh.GameInitAck{
			Table:   table.Id,
			Id:      table.CurId,
			State:   int32(table.State),
			Players: table.playerInfo,
			Poker:   nil,
		}
		if round := table.round; round != nil {
			ack.Ring = round.Ring
			ack.Pool = round.Pool
		}
		role.Send(ack)
	}
}

func (table *Table) newGameRound() {
	table.reset()
	count := table.RoleCount
	ante := room.Config.Ante // 底注
	id := room.NewGameRoundId()
	round := &GameRound{
		GameRound: zjh.GameRound{
			Id:    id,
			Start: room.Now(),
			Room:  room.RoomId,
			Tab:   table.Id,
			Bill:  make([]*zjh.GameBill, 0, count),
			Ante:  int32(ante),
			Ring:  1,
		},
	}
	table.round = round
	offset := int32(0)
	pokers := model.NewPoker(1, false, true)
	players := make([]*Player, 0, count)
	playerInfo := make([]*zjh.Player, 0, count)
	// 随机首家
	first := gameRand.Int31n(count)
	note := &strings.Builder{}
	note.Grow(120)
	for i := first; i < chairCount+first; i++ {
		role := table.Roles[i%chairCount]
		if role == nil {
			continue
		}
		//note.WriteString(strconv.Itoa(int(role.Id)))
		poker := pokers[offset : offset+3]
		model.PokerStringAppend(note, poker)
		offset += 3
		note.WriteString("|")
		role.newGameRound(poker)
		players = append(players, role.Player)
		playerInfo = append(playerInfo, &role.Player.Player)
		round.Bill = append(round.Bill, role.Bill)
		// 打底
		role.decCoin(ante)
	}
	round.Note = note.String()
	table.waitSecond = 0
	table.poker = pokers
	table.playIndex = 0
	table.players = players
	table.playerInfo = playerInfo
	table.winner = nil
	table.firstAllin = nil
	table.playCount = int32(len(players))
	table.lookCount = 0
	table.compCount = 0
	// 设置第一个玩家的超时时间
	players[0].Down = waitSecond
}

func (table *Table) Init() {

}

func (table *Table) Start() {
	table.gameReset()
}

func (table *Table) Update() {
	switch table.State {
	case GameStateReady:
		if room.Config.Pause == 0 {
			table.gameReady()
		}
	case GameStatePlaying:
		table.gamePlay()
	case GameStateCheckout:
		table.gameCheckout()
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

// 重置
func (table *Table) gameReset() {
	// 第1次准备
	table.State = GameStateReady
	table.waitSecond = 0
	table.CurId += 1
	log.Debugf("[%v]准备:%v", table.Id, table.CurId)

	// 设置玩家举手倒计时
	for _, role := range table.Roles {
		if role == nil {
			continue
		}
		player := role.Player
		if player.State != zjh.Player_Ready {
			player.State = zjh.Player_None
			player.Down = readySecond
		}
	}
}

func (table *Table) gameReady() {
	table.waitSecond++
	// 真人数量
	realCount := 0
	// 所有人都已经准备好
	allReady := true
	for _, role := range table.Roles {
		if role == nil {
			continue
		}
		if !role.IsRobot() {
			realCount++
		}
		player := role.Player
		if player.State == zjh.Player_None {
			player.Down--
			if role.IsRobot() {
				// 机器人准备时间0-4秒,不大于5秒
				if gameRand.Int31n(4) == 0 || table.waitSecond >= 5 {
					player.State = zjh.Player_Ready
					log.Debugf("[%v]举手:%v", table.Id, player.Id)
				}
			} else if player.Down < 0 {
				player.State = zjh.Player_Ready
				// TODO： 真人超时了，强制退出
				//table.removeRole(i)
			}
		}
		if player.State != zjh.Player_Ready {
			allReady = false
		}
	}

	if realCount == 0 {
		//// 没有真人了机器人退出
		//table.freeRobots()
		//return
	}

	// 所有人都已准备，有2个人就可以开始了
	if allReady && (table.RoleCount >= 2) {
		table.gameOpen()
	}
}

// 开始
func (table *Table) gameOpen() {
	// 发送开始下注消息给所有玩家
	table.State = GameStatePlaying
	table.waitSecond = 0
	table.newGameRound()
	log.Debugf("[%v]开始:%v,%v", table.Id, table.CurId, table.round.Note)
	// 发送消息给玩家
	table.SendToAll(&zjh.GameStartAck{
		Id:      table.CurId,
		Players: table.playerInfo,
	})

	// 启动机器人
	for _, role := range table.Roles {
		if role == nil || role.Robot == nil {
			continue
		}
		role.Robot.Start(role)
	}
}

func (table *Table) gamePlay() {
	table.waitSecond++
	player := table.runner()
	if player.State == zjh.Player_Playing {
		player.Down--
		if player.Down < 0 {
			// 超时弃牌
			table.Roles[player.Chair].Discard(true)
		}
	}
	table.robotPlay()
}

// 机器人玩牌
func (table *Table) robotPlay() {
	for _, role := range table.Roles {
		if table.State != GameStatePlaying {
			return
		}
		if role == nil || role.Player == nil {
			continue
		}
		if robot := role.Player.Robot; robot != nil {
			switch role.Player.State {
			case zjh.Player_Playing:
				robot.Play(role)
			case zjh.Player_Discard:
				robot.Discard(role)
			case zjh.Player_Lose:
				robot.Lose(role)
			}
		}
	}
}

// 结算结果发给玩家
func (table *Table) sendGameResult() {
	round := table.round
	for _, role := range table.Roles {
		if role == nil || role.IsRobot() {
			continue
		}
		// 看自己的牌+PK过的牌
		poker := make([]byte, 3*chairCount)
		for _, opp := range table.players {
			if opp.Id == role.Id || role.IsOpponent(opp.Id) {
				i := opp.Chair
				copy(poker[i:i+3], table.poker[i:i+3])
			}
		}
		// 发送游戏结束
		role.UnsafeSend(&zjh.GameResultAck{
			Id:     table.CurId,
			Prize:  round.prize,
			Winner: round.winner,
			Coin:   role.Coin,
			Poker:  poker,
			Lucky:  round.Lucky,
		})
	}
}

// 结算
func (table *Table) gameClose() {
	table.State = GameStateCheckout
	table.waitSecond = 0
	log.Debugf("[%v]结算:%v", table.Id, table.CurId)

	// 计算玩家和机器人之间的输赢
	robotWin := true //
	playerWin := true
	for _, winner := range table.winner {
		if winner.IsRobot() {
			playerWin = false
		} else {
			robotWin = false
		}
	}

	round := table.round
	if robotWin {
		//计算真实玩家输的钱
		for _, bill := range round.Bill {
			if bill.Job != model.JobRobot {
				bill.Robot = -bill.Bet
			}
		}
	} else if playerWin {
		// 计算所有机器人输的钱
		sumRobotLose := int64(0)
		for _, bill := range round.Bill {
			if bill.Job == model.JobRobot {
				sumRobotLose += bill.Bet
			}
		}
		// 所有赢钱的玩家平滩机器人输的钱
		sumRobotLose -= sumRobotLose * (room.Config.Tax + poolTax) / 1000
		argWin := sumRobotLose / int64(len(table.winner))
		for _, winner := range table.winner {
			winner.Bill.Robot = argWin
		}
	} else {
		// 机器和玩家都有赢家则忽略计算
	}

	// 保存牌局
	round.End = room.Now()
	room.SaveLog(&round.GameRound)

	// 结算结果发给玩家
	table.sendGameResult()

	// 结束机器人
	for _, role := range table.Roles {
		if role == nil || role.Robot == nil {
			continue
		}
		role.Robot.End(role)
	}

	// TODO:清理钱不够的机器人和玩家
	table.clearRole()
}

func (table *Table) gameCheckout() {
	table.waitSecond++
	if table.waitSecond == 5 {
		// 5秒后开始下一局
		table.gameReset()
	}
}

//
func (table *Table) clearRole() {
	// 删除钱不足或者钱多的机器人
	var ids []model.UserId
	for i, role := range table.Roles {
		if role == nil {
			continue
		}
		if role.IsRobot() {
			if role.TotalRound > rand.Int31n(64)+10 ||
				role.Coin < room.Config.PlayMin ||
				role.Coin > room.Config.PlayMax ||
				role.TotalWin > 10000*100 {
				ids = append(ids, role.Id)
				table.removeRole(i)
				role.Online = false
			}
		} else {
			if role.Online == false {
				table.removeRole(i)
			}
		}
	}
	if len(ids) > 0 {
		unloadRobots(ids)
		log.Debugf("[%v][%v]clearRobot:%v", table.Id, table.State, ids)
	}
}
