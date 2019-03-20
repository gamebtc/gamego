package util

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

type AppSignal struct {
	sync.WaitGroup
	die chan struct{}
	close chan struct{}
	closeOnce int32
}

func NewAppSignal() *AppSignal {
	a := new(AppSignal)
	a.die = make(chan struct{})
	a.close = make(chan struct{})
	return a
}

// handle unix signals
func(app *AppSignal) Run(f func()) {
	defer PrintPanicStack()
	f()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	select{
		case <-ch:
		case <-app.close:
	}
	close(app.die)
	app.Wait()
	os.Exit(0)
}

func(app *AppSignal) Die() <-chan struct{} {
	return app.die
}

func(app *AppSignal) Close() bool {
	if atomic.CompareAndSwapInt32(&app.closeOnce, 0, 1) {
		close(app.close)
		return true
	}
	return false
}