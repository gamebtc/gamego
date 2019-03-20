package main

import (
"time"

log "github.com/sirupsen/logrus"
)


func test1(){
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

func test2() {
	i := 0
	period := time.Duration(2000) * time.Millisecond
	timer := time.NewTimer(period)
	defer timer.Stop()

	time.Sleep(time.Duration(1) * time.Second)
	go func(){
		timer.Stop()
		timer.Stop()
		timer.Stop()
	}()

	select {
	case <-timer.C: // 帧更新
		i++
		log.Printf("TestTicker count:%v", i)
		time.Sleep(time.Duration(3) * time.Second)
	default:
		log.Printf("default count:%v", i)
	}
}


func main() {
	test2()
	test1()
}