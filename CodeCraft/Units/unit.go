package Units

import "awesomeProject1/CodeCraft/Player"

type Unit interface {
	Attack(enemy Unit) bool
	GetHit(int64) bool
	Price(p *Player.Player) bool
}
