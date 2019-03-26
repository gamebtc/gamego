package internal

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"local.com/abc/game/model"
	"local.com/abc/game/protocol"
)

const BAN_LOGIN = 0  //禁止登录
const BAN_REGIST = 1 //禁止注册

func loginFailAck(code int32, args ...interface{}) *protocol.LoginFailAck {
	m := driver.GetHint(code)
	if len(args) != 0 {
		m = fmt.Sprintf(m, args...)
	}
	return &protocol.LoginFailAck{Code: code, Msg: m}
}

func getBan(ban []int32, i int) int32 {
	if len(ban) > i {
		return ban[i]
	}
	return 0
}

func getRandName() string {
	return ""
}

// 登录
func userLogin(ctx context.Context, in interface{}) interface{} {
	uctx := ctx.(*protocol.UserContext)
	uid, agent := uctx.UserId, uctx.AgentId
	req := in.(*protocol.LoginReq)
	log.Infof("login a:%v,u:%v,i:%v, req: %#v", agent, uid, uctx.Ip, req)
	now := driver.Now()
	nowInt := int32(now.Unix())
	//检查包配置
	chanConf := driver.GetChanConf(req.Env.Chan)
	if chanConf == nil || getBan(chanConf.Ban, BAN_LOGIN) > nowInt {
		return loginFailAck(10001, req.Env.Chan)
	}
	ip := model.IP(uctx.Ip)
	//检查IP
	ipInfo := driver.GetIpInfo(ip)
	if ipInfo != nil && getBan(ipInfo.Ban, BAN_LOGIN) > nowInt {
		return loginFailAck(10003)
	}
	//检查机器码
	udid := driver.GetMachineInfo(req.Udid)
	if udid != nil && getBan(udid.Ban, BAN_LOGIN) > nowInt {
		return loginFailAck(10004)
	}

	// 获取账号
	acc, err := driver.GetAccount(req.Env.Id, req.Type, req.Name)
	if err != nil {
		if getBan(chanConf.Ban, BAN_REGIST) > nowInt {
			return loginFailAck(10101)
		}
		// 创建账号
		if ipInfo != nil && getBan(ipInfo.Ban, BAN_REGIST) > nowInt {
			return loginFailAck(10102)
		}
		// 检查IP
		ipInfo := driver.GetIpInfo(ip)
		if ipInfo != nil && getBan(ipInfo.Ban, BAN_REGIST) > nowInt {
			return loginFailAck(10103)
		}
		if acc == nil {
			acc = new(model.Account)
		}
		id, err := driver.NewId()
		if err != nil {
			//获取ID失败
			return loginFailAck(10105, req.Env.Chan)
		}
		acc.Id = id
		acc.App = req.Env.Id
		acc.Type = req.Type
		acc.Name = req.Name
		acc.Pwd = req.Pwd
		acc.Pack = req.Env.Pack
		acc.Chan = chanConf.Id
		acc.Ip0 = ip
		acc.Udid = req.Udid
		acc.Users = make([]int32, 0, 1)
		acc.Init = now
		acc.Up = now
		err = driver.CreateAccount(acc, req)
		if err != nil {
			return loginFailAck(10104)
		}
	}
	if req.Pwd == "" || acc.Pwd != req.Pwd {
		// 账号不存在或者密码错误
		return loginFailAck(10007)
	}

	var user *model.User
	if len(acc.Users) == 0 {
		// 创建玩家
		bag := appConf.GetInitCoin()
		user = &model.User{
			App:  req.Env.Id,
			Act:  acc.Id,
			Pack: req.Env.Pack,
			Chan: acc.Chan,
			Ip0:  ip,
			Last: now,
			Ip:   ip,
			Bag:  bag,
			Init: now,
			Up:   now,
		}
		user.Name = getRandName()
		if err = driver.CreateUser(user, req); err == nil {
			acc.Users = append(acc.Users, user.Id)
		}
	} else {
		// 获取玩家
		if uid == 0 {
			uid = acc.Users[0]
		}
		user, err = driver.LoadUser(uid, ip)
		log.Debugf("login%#v, err:%v", user, err)
	}

	if err != nil || user == nil || user.Id == 0 {
		// 创建玩家失败，请重试
		return loginFailAck(10008, err)
	}

	// 锁定账号
	var lock *model.UserLocker
	if lock, err = driver.LockUser(agent, user.Id, user.Ip, now, req); err != nil {
		// 锁定玩家失败，请重试
		return loginFailAck(10009, err)
	}

	if lock.Agent != agent {
		// TODO：通知服务器其它地方登录
		log.Debugf("login%#v", user)
	}

	return &protocol.LoginSuccessAck{
		Id:    int32(user.Id),
		Agent: agent,
		Icon:  user.Icon,
		Sex:   user.Sex,
		Vip:   user.Vip,
		Act:   user.Act.Hex(),
		Name:  user.Name,
		Phone: acc.Phone,
		Bag:   user.Bag,
		Kind:  lock.Kind,
		Room:  lock.Room,
	}
}
