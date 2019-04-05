package main

import (
	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

// 创建socket连接
func StartSession(s *Session) {
	log.Debugf("new session: %v", s.Addr)
	id := atomic.AddInt32(&connectCount, 1)
	s.Ip = uint32(id)
	conn := s.conn
	defer func() {
		atomic.AddInt32(&connectCount, -1)
		conn.Close()
	}()
	signal.Add(1)
	defer func(){
		signal.Done()
		delSession(s)
	}()
	s.Start()
}

