package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"awesomeProject1/CodeCraft/Utils"
	"fmt"
)

type Player struct {
	Gold int64
	Army []Units.Unit
	Name string
}

func (p *Player) ChooseName() {
	fmt.Println("May I know your name?")
	p.Name = Utils.ScanText()
	fmt.Println("Greetings,", p.Name, "!")

}

func NewPlayer() Player {
	p := Player{100, make([]Units.Unit, 0), ""}
	p.ChooseName()
	return p

}

func (Player) IsReady() bool {
	i := Utils.ScanInt()
	if i == 1 {
		return true
	} else if i == 2 {
		return false
	}
	return false
}
