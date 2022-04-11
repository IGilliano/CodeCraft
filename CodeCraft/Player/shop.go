package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"awesomeProject1/CodeCraft/Utils"
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

func (s Shop) SelectUnit(p *Player) bool {
	var r bool
	fmt.Println("Which one do you want to hire?\n 1: Barbarian\n 2: Hunter\n 3: Knight\n 0: Exit")
	u := Utils.ScanInt()
	if u > 0 && u < 4 {
		r = s.Hire(u, p, r)
	} else if u == 0 {
		fmt.Println("Okay then. Come back later!")
		r = false
		return r
	} else {
		fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
	}
	return r
}

func (s Shop) Hire(u int64, p *Player, r bool) bool {
	newUnit := s.maker[u]()
	c := newUnit.GetCost()
	if c <= p.Gold {
		p.Army = append(p.Army, newUnit)
		p.Gold = p.Gold - c
		fmt.Println("Done! Wonna hire more?\n 1: Yes 2: No")
		r = p.IsReady()
	} else {
		fmt.Println("Oh! Its too expensive for you, buddy!")
	}
	return r
}
