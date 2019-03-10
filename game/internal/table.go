package internal

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/db"
	//"local.com/abc/game/model"
	"local.com/abc/game/msg"
	"local.com/abc/game/room"
)

const RicherCount = 6

type Plan struct {
	// 要执行的函数
	f func(*Table)
	// 下一步执行的时间
	d time.Duration
}

type GameDriver interface {
	// 准备游戏, 状态1
	Ready(table *Table)
	// 开始下注, 状态2
	Open(table *Table)
    // 游戏中, 状态2
    Play(table *Table)
	// 停止下注, 状态3
	Stop(table *Table)
	// 发牌结算, 状态4
	Deal(table *Table)
}

// 开始下注：下注时间12秒
// 停止下注：发牌开：2秒
// 结算：    2秒
// 游戏桌子
type Table struct {
	GameDriver
	Id     int32       //桌子ID
	CurId  int32       //当前局的ID
	LastId int32       //最后的局ID
	Log    []byte      //最后60局的输赢情况0:龙赢,1:虎赢,2:和
	State  int32       //0:暂停;1:洗牌;2:下注;3:结算
	Roles  []*Role     //所有真实游戏玩家
	Robot  []*Role     //所有机器人
	Richer []*msg.User //富豪
	round  *GameRound

	timer     *time.Timer
	planIndex int
}

func NewTable() *Table {
	dealer := newDriver()
	t := &Table{
		Roles:      make([]*Role, 0, 100),
		Robot:      make([]*Role, 0, 100),
		GameDriver: dealer,
	}
	return t
}

// 增加真实的玩家
func (table *Table) AddRole(role *Role) {
	table.Roles = append(table.Roles, role)
	// 真实玩家
	ack := &msg.FolksGameInitAck{
		Id:    table.round.Id,
		State: table.State,
		Sum:   table.round.Group,
		Log:   table.Log,
	}
	if bill := role.bill; bill != nil {
		ack.Bet = bill.Group
	}
	ack.Rich = table.Richer
	role.Send(ack)
}

// 查找1位赌神和5位富豪
func  (table *Table) FindRicher() {
	roleCount := len(table.Roles) + len(table.Robot)
	if roleCount == 0 {
		table.Richer = []*msg.User{}
		return
	}

	roles := make([]*Role, roleCount)

	i := 0
	for _, role := range table.Robot {
		roles[i] = role
		i++
	}
	for _, role := range table.Roles {
		roles[i] = role
		i++
	}

	// 查找1位赌神
	richIndex := 0
	rich := roles[0]
	for i := 1; i < roleCount; i++ {
		b := roles[i]
		if rich.TotalWin < b.TotalWin {
			roles[i] = rich
			rich = b
			richIndex = i
		}
	}
	richer := []*msg.User{rich.GetMsgUser()}
	roles = append(roles[0:richIndex], roles[richIndex+1:]...)
	roleCount--

	// 查找5位富豪
	for c := 0; c < 5 && c < roleCount; c++ {
		rich := roles[c]
		for i := c + 1; i < roleCount; i++ {
			b := roles[i]
			if rich.Coin < b.Coin {
				// 交换位置
				roles[i] = rich
				rich = b
			}
		}
		richer = append(richer, rich.GetMsgUser())
	}
	table.Richer = richer
}

func (table *Table) NewGameRound() {
	id := room.NewGameRoundId()
	g := make([]int64, 2*betItem)
	round := &GameRound{
		Id:        id,
		Start:     room.Now(),
		Room:      room.RoomId,
		Tab:       table.Id,
		Group:     g[0:betItem],
		UserGroup: g[betItem:],
	}

	for _, role := range table.Roles {
		role.Reset()
	}
	table.FindRicher()

	table.round = round
}

func (table *Table) Init(){
}


func (table *Table) Start() {
	table.timer = time.NewTimer(time.Microsecond)
}

func (table *Table) Update() {
	select {
	case <-table.timer.C:
		cur := table.planIndex
		table.timer.Reset(schedule[cur].d)
		if cur >= len(schedule)-1 {
			table.planIndex = 0
		} else {
			table.planIndex = cur + 1
		}
		schedule[cur].f(table)
	default:
	}
}

// 返回系统输赢
func (table *Table)CheckWin(odds []int32) int64 {
	prize, _, bet := Balance(table.round.UserGroup, odds)
	return bet - prize
}

// 投注
func betReq(m *room.NetMessage) {
	log.Debugf("betReq:%v", m)

	role, ok := m.Session.Role.(*Role)
	if ok == false || role == nil {
		log.Debugf("betReq: role is nil")
		return
	}
	//获取参数
	req, ok := m.Arg.(*msg.BetReq)
	if ok == false || req == nil {
		log.Debugf("betReq: is nil")
		return
	}
	log.Debugf("add bet: %v", req)
	role.AddBet(*req)
}

func(table *Table) loadRobot(count int32) {
	if count < 0 {
		// 退出部分机器人
	} else if count > 0 {

		// 增加机器人
		robots := db.Driver.LoadRobot(room.RoomId, count)
		for _,user:=range robots{
			robot := &Role{
				Sender: roleSender,
				User: user,
			}
			robot.table = table
			robot.Online = true
			robot.Reset()
			table.Robot = append(table.Robot,robot)
		}
	}
}

// 准备
func gameReady(table *Table) {
	// 加载机器人
	robotConf := room.Config.Robot
	end := len(robotConf) / 6
	now := time.Now()
	minute := int32(now.Hour()*60 + now.Minute())
	for i := 0; i < end; i++ {
		if minute <= robotConf[i] && minute > robotConf[i+1] {
			min := robotConf[i+2]  //最小人数
			max := robotConf[i+3]  //最大人数
			base := robotConf[i+4] //基础人数
			rate := robotConf[i+5] //真实玩家的百分比人数
			//真人数量
			roleCount := int32(len(table.Roles))
			count := base + roleCount*rate/100
			if count < min {
				count = min
			} else if count > max {
				count = max
			}
			table.loadRobot(int32(len(table.Robot)) - count)
			break
		}
	}

	table.State = 1
	table.CurId += 1
	log.Debugf("%v准备:%v", gameName, table.CurId)
	table.NewGameRound()
	table.Ready(table)
}

// 开始
func gameOpen (table *Table){
	// 发送开始下注消息给所有玩家
	table.State = 2
	log.Debugf("开始下注:%v", table.CurId)
	table.Open(table)
}

func gamePlay(table *Table) {
	table.Play(table)
	for _, role := range table.Robot {
		if role.Id%5 == rand.Int31n(5) {
			bet := msg.BetReq{
				Item: rand.Int31n(int32(betItem)),
				Coin: 100 + rand.Int31n(100)*100,
			}
			if role.RobotCanBet(bet.Item) {
				role.AddBet(bet)
				//log.Debugf("机器人下注:%v,%v", bet, r)
			}
		}
	}
}

// 停止下注
func gameStop (table *Table) {
	table.State = 3
	log.Debugf("停止下注:%v", table.CurId)
	table.Stop(table)
}

// 发牌结算
func gameDeal(table *Table) {
	table.State = 4
	log.Debugf("发牌结算:%v", table.CurId)
	// 发牌结算
	table.Deal(table)

	// 结算结果发给玩家
	table.LastId = table.CurId
	round := table.round
	round.End = room.Now()
	room.SaveLog(round)

}