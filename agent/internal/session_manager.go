package internal

import (
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	. "local.com/abc/game/protocol"
)

var (
	locker     sync.RWMutex
	sessions   map[int64]*Session //在线用户列表，key:连接ID
	userLocker sync.RWMutex
	users      map[UserId]*Session //已登录用户列表
)

func init() {
	sessions = make(map[int64]*Session, 4000)
	users = make(map[UserId]*Session, 2000)
}

// create a new session object for the connection
func newSession(ip uint32, addr string, conn net.Conn) {
	id := newAgentId()
	for id < 0 {
		id = newAgentId()
	}
	signal.Add(1)
	defer signal.Done()
	s := &Session{
		Id:       id,
		Ip:       ip,
		Addr:     addr,
		Created:  time.Now(),
		Coder:    coder,
		conn:     conn,
		dieChan:  make(chan struct{}),
		stopRecv: make(chan struct{}),
	}
	defer delSession(s)
	s.Start()
	log.Debugf("new session %v: %v", s.Id, s.Addr)
	return
}

func delSession(s *Session) {
	removeUser(s)
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

func sessionRange(f func(session *Session)) {
	locker.RLock()
	defer locker.RUnlock()
	for _, v := range sessions {
		if v != nil {
			f(v)
		}
	}
}

func userCount() int {
	userLocker.RLock()
	l := len(users)
	userLocker.RUnlock()
	return l
}

func getUser(id int32) (user *Session) {
	userLocker.RLock()
	user = users[id]
	userLocker.RUnlock()
	return
}

// addUser 添加在线用户
func addUser(user *Session) (old *Session) {
	if id := user.UserId; id != 0 {
		userLocker.Lock()
		if old = users[id]; old != user {
			users[id] = user
		}
		userLocker.Unlock()
	}
	return
}

func removeUser(user *Session) (r bool) {
	if id := user.UserId; id != 0 {
		userLocker.Lock()
		if old := users[id]; old == user {
			delete(users, id)
			r = true
		}
		userLocker.Unlock()
	}
	return
}
