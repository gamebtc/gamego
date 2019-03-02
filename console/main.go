package main

import (
"time"

log "github.com/sirupsen/logrus"
)

func main() {
	i := 0
	period := time.Duration(2000) * time.Millisecond
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for i < 20 {
		select {
		case <-ticker.C: // 帧更新
			i++
			log.Printf("TestTicker count:%v", i)
			time.Sleep(time.Duration(3) * time.Second)
		}
	}
}