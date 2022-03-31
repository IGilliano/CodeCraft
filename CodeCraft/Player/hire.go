package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"fmt"
)

//create struct (or interface) shop

func Hire(p *Player) {

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
	} else if n == 0 {
		fmt.Println("Okay then. Come back later!")
	} else {
		fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
	}
}
