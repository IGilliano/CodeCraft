package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"fmt"
)

type Player struct {
	Gold int64
	Army []Units.Unit
	Name string
}

func NewPlayer() Player {
	return Player{1000, make([]Units.Unit, 0), ""}
}

func (p *Player) Hire() {

	maker := make(map[int64]func() Units.Unit)
	maker[1] = Units.NewBarbarian
	maker[2] = Units.NewHunter
	maker[3] = Units.NewHunter //Knight

	var n int64
	fmt.Println("Which one do you want to hire?\n 1: Barbarian\n 2: Hunter\n 3: Knight\n 0: Exit")
	fmt.Scanf("%d", &n)
	if n > 0 && n < 4 {
		newUnit := maker[n]()
		p.Army = append(p.Army, newUnit)
		newUnit.Price(p)
	} else if n == 0 {
		fmt.Println("Okay then. Come back later!")
	} else {
		fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
	}
}
func (p Player) ChooseName() {
	fmt.Println("May I know your name?")
	fmt.Scanf("%s", &p.Name)
	fmt.Println("Greetings,", p.Name, "!")
}
