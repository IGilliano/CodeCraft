package Player

import (
	"awesomeProject1/CodeCraft/Units"
)

type Player struct {
	Gold int64
	Army []Units.Unit
}

func NewPlayer() Player {
	return Player{1000, make([]Units.Unit, 0)}
}
