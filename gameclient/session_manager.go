package main

import (

	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	agentHead  int64
	sessionId  uint32
	locker     sync.RWMutex
	sessions   map[int64]*Session  //在线用户列表，key:连接ID
	userLocker sync.RWMutex
	users      map[int32]*Session  //已登录用户列表
)

func init() {
	sessions = make(map[int64]*Session, 4000)
	users = make(map[int32]*Session, 2000)
}


func delSession(s *Session) {
	removeSession(s)
	s.destroy()
	log.Debugf("del session %v: %v", s.Id, s.Addr)
}

func sessionCount() int {
	locker.RLock()
	l := len(sessions)
	locker.RUnlock()
	return l
}

func getSession(id int64) (session *Session) {
	locker.RLock()
	session = sessions[id]
	locker.RUnlock()
	return
}

func addSession(session *Session) (r bool) {
	if id := session.Id; id != 0 {
		locker.Lock()
		if old, found := sessions[id]; found == false {
			sessions[id] = session
			r = true
		} else if old == session {
			r = true
		}
		locker.Unlock()
	}
	return
}

func removeSession(session *Session) (r bool) {
	if id := session.Id; id != 0 {
		locker.Lock()
		if old := sessions[id]; old == session {
			delete(sessions, id)
			r = true
		}
		locker.Unlock()
	}
	return
}