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

func test3() {
	v := []int32{1, 2, 3, 4, 5, 6}
	for i, t := range v {
		if t == 2 {
			v = append(v[:i], v[i+1:]...)
		}
		log.Printf("test3:%v", t)
	}
	// output:124566
}

func test4() {
	m := map[string]int64{}
	t := m["bb"]
	log.Printf("test4:%v", t)
}

func main() {
	test4()
	test3()
	test2()
	test1()
}