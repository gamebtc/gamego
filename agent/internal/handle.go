package internal

import (
	"bytes"
	"errors"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"
	. "local.com/abc/game/protocol"
)

// 握手
func handshakeHandler(sess *Session, data []byte) (interface{}, error) {
	return nil, nil
}

// 心跳检查
func heartBeatHandler(sess *Session, data []byte) (interface{}, error) {
	arg := HeartBeatReq{}
	if err := sess.Unmarshal(data[HeadLen:], &arg); err != nil {
		return nil, err
	}
	log.Debugf("heartBeat:uid%v, id:%v", sess.UserId, arg.Id)
	ack := &HeartBeatAck{Id: arg.Id}
	return ack, nil
}

const checkVerSign = true
// 版本检查
func verCheckHandler(sess *Session, data []byte) (interface{}, error) {
	if checkVerSign && (sess.Flag != SESS_INIT) {
		return nil, nil
	}

	arg := VerCheckReq{}
	if err := sess.Unmarshal(data[HeadLen:], &arg); err != nil {
		return nil, err
	}

	packId := strconv.Itoa(int(arg.Env.Pack))
	// 检查签名
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(arg.Check)))
	buffer.WriteString(strconv.Itoa(int(arg.Time)))
	buffer.WriteString(strconv.Itoa(int(arg.Env.Id)))
	buffer.WriteString(packId)
	buffer.WriteString(arg.Env.Ver)
	buffer.WriteString(arg.Env.Chan)
	buffer.WriteString(arg.Env.Refer)
	buffer.WriteString(arg.Env.Other)

	//sign := md5.Sum(buffer.Bytes())
	////校验签名
	//if checkVerSign && (sign[0] == 234) { //应该应该为!=0
	//	log.Debugf("verCheck:Id:%v, has:%v", sess.Id, sign)
	//	return nil, ErrorSign
	//}

	//转发到服务器处理
	ret, err := sess.callServer(int32(MsgId_VerCheckReq), data)
	if sess.Flag == SESS_INIT {
		if ret, ok := ret.([]byte); ok && len(ret) > HeadLen {
			if GetHeadId(ret) == int32(MsgId_VerCheckAck) {
				sess.Flag = SESS_VERCHECK
				sess.Seed = uint32(arg.Check)
			}
		}
	}
	log.Debugf("verCheck:Id:%v, flag:%v", sess.Id, sess.Flag )
	return ret, err
}

// 玩家登录鉴权
func userLoginHandler(sess *Session, data []byte) (interface{}, error) {
	oldFlag := sess.Flag
	if checkVerSign && (oldFlag != SESS_VERCHECK) {
		// 未版本
		log.Debugf("userLogin:Id:%v, flag:%v", sess.Id, oldFlag)
		return &LoginFailAck{Code: 999, Msg: ErrorSign.Error()}, nil
	}
	if oldFlag == SESS_LOGINED {
		//重复登录
		return &LoginFailAck{Code: 999, Msg: ErrorDuplicateLogin.Error()}, nil
	}
	sess.Flag = SESS_LOGINING
	//转发到服务器处理
	ret, err := sess.callServer(int32(MsgId_UserLoginReq), data)
	if ret, ok := ret.([]byte); ok && len(ret) > HeadLen {
		if GetHeadId(ret) == int32(MsgId_UserLoginSuccessAck) {
			// 登录成功
			ack := LoginSuccessAck{}
			if err = sess.Unmarshal(ret[HeadLen:], &ack); err == nil {
				log.Debugf("LoginSuccess:%v", ack.Id)
				sess.SetUser(ack.Id)
				sess.Act = ack.Act
				if old := addUser(sess); old != nil && old != sess {
					old.Close()
					removeSession(old)
				}
				sess.Flag = SESS_LOGINED
			} else {
				log.Debugf("LoginFail:%v", err.Error())
				sess.Flag = oldFlag
			}
			return ret, err
		}
	}
	sess.Flag = oldFlag
	return ret, err
}

// 玩家连接游戏房间
func loginRoomHandler(sess *Session, date []byte) (interface{}, error) {
	req := LoginRoomReq{}
	if err := sess.Unmarshal(date[HeadLen:], &req); err != nil {
		return nil, err
	}
	roomId := int32(req.Id)
	if sess.RoomId != 0 {
		sess.closeRoom()
	}

	ret, err := loginRoom(sess, roomId, date)
	if err != nil {
		return &LoginRoomAck{
			Room: roomId,
			Code: 1009,
			Msg:  err.Error(),
		}, nil
	}
	return ret, err
}

// 断开房间连接
func exitRoomHandler(sess *Session, date []byte) (interface{}, error) {
	//req := ExitRoomReq{}
	//if err := sess.Unmarshal(date[HeadLen:], &req); err != nil {
	//	return nil, err
	//}
	//roomId := int32(req.Id)
	if sess.RoomId == 0 {
		return &ExitRoomAck{
			Code: 1010,
			Msg:  "房间号错误",
		}, nil
	}

	sess.closeRoom()
	return &ExitRoomAck{
		Code: 0,
		Msg:  "success",
	}, nil
}

////
//func loginRoom(sess *Session, roomId int32, v []byte) (interface{}, error) {
//	// 连接到已选定游戏房间服务器
//	conn := roomServicePool.GetService(roomId)
//	if conn == nil {
//		log.Debugf("cannot get room:%v", roomId)
//		return nil, errors.New("cannot get room:" + strconv.Itoa(int(roomId)))
//	}
//	// 开启到游戏服的流
//	cli := NewGameClient(conn)
//	if s, err := cli.Send(sess.callCtx); err != nil {
//		log.Warnf("room error %v: %s", roomId, err.Error())
//		return nil, errors.New("cannot connect room:" + strconv.Itoa(int(roomId)))
//	} else {
//		stream := &GrpcStream{ClientStream: s}
//		return sess.loginRoom(roomId, v, stream)
//	}
//}

//
func loginRoom(sess *Session, roomId int32, v []byte) (interface{}, error) {
	// 连接到已选定游戏房间服务器
	addr := rpcServicePool.GetValue(strconv.Itoa(int(roomId)))
	if addr == nil {
		log.Debugf("cannot get room:%v", roomId)
		return nil, errors.New("cannot get room:" + strconv.Itoa(int(roomId)))
	}

	conn, err := net.Dial("tcp", string(addr))
	if err != nil {
		log.Warnf("room error %s: %s", string(addr), err.Error())
		return nil, errors.New("cannot connect room:" + strconv.Itoa(int(roomId)))
	}

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetReadBuffer(tcpWriteBuf * 2)
		tcpConn.SetWriteBuffer(tcpReadBuf * 2)
	}

	// 发送头信息
	head := NewUserHead(sess.Id, sess.UserId, sess.Ip)
	if _, err := conn.Write(head); err != nil {
		return nil, err
	}
	// 发送登录消息
	stream := &NetStream{conn: conn}
	return sess.loginRoom(roomId, v, stream)
}