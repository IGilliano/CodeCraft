package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"fmt"
)

type Shop struct {
	roster []Units.Unit
	maker  map[int64]func() Units.Unit
}

func NewShop() Shop {
	New := Shop{[]Units.Unit{}, make(map[int64]func() Units.Unit)}

	New.maker[1] = Units.NewBarbarian
	New.maker[2] = Units.NewHunter
	New.maker[3] = Units.NewHunter //Knight

	return New
}

func (p *Player) Hire(s Shop) {
	var n int64
	fmt.Println("Which one do you want to hire?\n 1: Barbarian\n 2: Hunter\n 3: Knight\n 0: Exit")
	fmt.Scanf("%d", &n)
	if n > 0 && n < 4 {
		newUnit := s.maker[n]()
		if newUnit.Price(p) == true {
			p.Army = append(p.Army, newUnit)
			fmt.Println("Done!")
		} else {
			fmt.Println("Oh! Its too expensive for you, buddy!")
		}
	} else if n == 0 {
		fmt.Println("Okay then. Come back later!")
	} else {
		fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
	}
}
