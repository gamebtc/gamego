package util

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type AppSignal struct {
	sync.WaitGroup
	die chan struct{}
}

func NewAppSignal() *AppSignal {
	a := new(AppSignal)
	a.die = make(chan struct{})
	return a
}

// handle unix signals
func(app *AppSignal) Run(f func()) {
	defer PrintPanicStack()
	f()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-ch
	close(app.die)
	app.Wait()
	os.Exit(0)
}

func(app *AppSignal) Die() <-chan struct{} {
	return app.die
}