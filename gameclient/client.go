package main

import (
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

// 创建socket连接
func StartSession(s *Session) {
	id := atomic.AddInt32(&connectCount, 1)
	s.Ip = uint32(id)
	conn := s.conn
	defer func() {
		atomic.AddInt32(&connectCount, -1)
		conn.Close()
	}()
	signal.Add(1)
	defer func(){
		signal.Add(-1)
		delSession(s)
	}()
	log.Debugf("new session: %v", s.Addr)
	s.Start()
}

