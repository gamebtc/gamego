package room

import (
	"time"
)

type Timer struct {
	t *time.Timer
	f func()
}

func (t *Timer) Stop() {
	t.t.Stop()
	t.f = nil
}

func (t *Timer) Exec() {
	if t.f != nil {
		t.f()
	}
}
