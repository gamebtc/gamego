package internal

import (
	"local.com/abc/game/protocol/zjh"
)

type GameRound struct {
	zjh.GameRound
	winner []int32
	prize  []int64
}