package internal

import (
	"local.com/abc/game/protocol/zjh"
)

type GameRound struct {
	zjh.GameRound
	winner []int32   // 获胜者的ID
	prize  []int64   // 获胜者分得的钱
}