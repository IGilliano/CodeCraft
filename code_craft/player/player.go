package player

import (
	"awesomeProject1/code_craft/units"
	"awesomeProject1/code_craft/utils"
	"fmt"
)

type Player struct {
	Gold  int64
	Army  []units.Unit
	Name  string
	Ready bool
}

func (p *Player) ChooseName() {
	fmt.Println("May I know your name?")
	p.Name = utils.ScanText()
	fmt.Println("Greetings,", p.Name, "!")

}

func NewPlayer() Player {
	p := Player{1000, make([]units.Unit, 0), "", false}
	return p

}

func (p Player) Inventory() {
	fmt.Println(p.Army)
}
