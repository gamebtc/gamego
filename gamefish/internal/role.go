package internal

import (
	//"fmt"
	log "github.com/sirupsen/logrus"

	"errors"
	"fmt"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol/fish"
	"local.com/abc/game/room"
)

// 每个角色的游戏数据
type Role struct {
	model.User                    // 玩家信息
	Point                         // 位置
	*room.Session                 // 发送消息
	table         *Table          // 桌子
	round         *fish.GameRound // 输赢情况
	Player        *fish.Player    // 玩家信息
	bulletCount   int32           // 子弹数
	lastFireTick  int64
}

func (role *Role) isRobot() bool {
	return role.Job == model.JobRobot
}

func (role *Role) isPlayer() bool {
	return role.Job == model.JobPlayer
}

// 加金币
func (role *Role) addCoin(coin int64) {
	role.Coin += coin
	role.Player.Coin = role.Coin
}

// 发子弹
func (role *Role) shoot(req *fish.ShootReq) error {
	if minBulletBet > req.Bet || maxBulletBet < req.Bet {
		return nil
	}
	if role.bulletCount >= SystemConf.MaxBulletCount {
		return errors.New("炮管散热中!")
	}

	bet := int64(req.Bet)
	if bet > role.Coin {
		return errors.New(fmt.Sprintf("余额不足%v金币，请您充值！", bet))
	}

	if role.Coin < room.Config.PlayMin {
		return errors.New(fmt.Sprintf("余额%v金币以上才可以射击，请您充值！", room.Config.PlayMin))
	}

	table := role.table
	if table == nil {
		return nil
	}

	role.lastFireTick = table.tickMs()
	role.addCoin(-bet)

	bullet := popBullet()
	bullet.Id = table.newBullId()
	bullet.Uid = role.Id
	bullet.Client = req.Client
	bullet.Created = table.tickMs()
	bullet.Bet = req.Bet
	bullet.X = role.X
	bullet.Y = role.Y
	bullet.Fish = req.Fish
	bullet.Direction = req.Direction
	bullet.Job = role.Job
	bullet.Speed = SystemConf.BulletSpeed
	bullet.InitMove()

	role.bulletCount++
	table.bullets[bullet.Id] = bullet
	table.GameUpdateAck.Bullets = append(table.GameUpdateAck.Bullets, &bullet.Bullet)

	return nil
}

// 击中
func (role *Role) hit(bulletId, fishId int32) {
	table := role.table
	if table == nil {
		return
	}

	bullet, ok := table.bullets[bulletId]
	if ok == false {
		return
	}

	gift := int64(0)

	// TODO:查找鱼
	//if fish, ok := table.fishes[fishId]; ok {
	//	  //
	//
	//} else {
	//	fishId = 0
	//}

	if gift != 0 {
		role.addCoin(gift)
	}
	bet := int64(bullet.Bet)
	round := role.round
	round.Bet += bet
	round.Win += gift - bet

	// 只记1元以上的炮或者得分
	if bet >= 1*room.YUAN || gift >= 1*room.YUAN {
		round.Log = append(round.Log, int32(bet), fishId, int32(gift))
	}
	role.tryWriteCoin()
	table.removeBullet(bullet)
	role.bulletCount--
}

func (role *Role) tryWriteCoin() {
	round := role.round
	if len(round.Log) >= 100*3 || round.Bet >= 1000*room.YUAN ||
		round.Win >= 1000*room.YUAN || round.Win <= -1000*room.YUAN {
		role.writeCoin()
	}
}

func (role *Role) writeCoin() {
	role.FlowSn = room.NewKindSn()
	table := role.table
	round := role.round
	round.End = room.Now()
	if role.isRobot() == false {
		round.Id = room.NewGameRoundId()
		flow := &model.CoinFlow{
			Sn:    role.FlowSn,
			Uid:   role.Id,
			Add:   round.Win,
			New:   role.Coin,
			Old:   round.OldCoin,
			Room:  room.RoomId,
			Game:  room.GameId,
			Bet:   round.Bet,
			LogId: round.Id,
		}
		// 写分
		if room.WriteCoin(flow) == nil {
			if role.Coin != flow.New {
				log.Warnf("[%v]金币变化:%v-%v", table.Id, flow.New, role.Coin)
			}
		}
		// 保存牌局
		room.SaveLog(round)
		log.Debugf("[%v]结算:%v", table.Id, flow)
	}

	// 写日志后重置
	round.Bet = 0
	round.Win = 0
	round.Start = round.End
	round.OldCoin = role.Coin
	round.Log = round.Log[0:0]
	round.Note = ""
}
