package Shop

import "awesomeProject1/CodeCraft/Player"

type Shop interface {
	Hire(p *Player.Player)
}
