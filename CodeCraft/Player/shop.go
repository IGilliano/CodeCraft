package Player

import (
	"awesomeProject1/CodeCraft/Units"
	"awesomeProject1/CodeCraft/Utils"
	"fmt"
)

type Shop struct {
	Roster []Units.Unit
	maker  map[int64]Units.Unit
}

func NewShop() Shop {
	New := Shop{[]Units.Unit{Units.NewBarbarian(), Units.NewHunter()}, make(map[int64]Units.Unit)}

	New.maker[1] = Units.NewBarbarian()
	New.maker[2] = Units.NewHunter()
	New.maker[3] = Units.NewHunter() //Knight

	return New
}

func (s Shop) ShopMenu(player *Player) {
	fmt.Println("What do you want to do?\n 1: Hire\n 2: Show roster\n 3: Exit")
	slct := Utils.ScanInt()
	switch slct {
	case 1:
		s.SelectUnit(player)
	case 2:
		s.ShowRoster()
	case 3:
		return
	default:
		fmt.Println("I dont understand! Please, select one of options!")
	}
}
func (s Shop) ShowRoster() {
	fmt.Println(s.Roster)
}

func (s Shop) SelectUnit(player *Player) {

	fmt.Println("Which one do you want to hire?\n 1: Barbarian\n 2: Hunter\n 3: Knight\n 0: Exit")
	n := Utils.ScanInt()
	if n == 0 {
		fmt.Println("Okay then. Come back later!")
		return
	} else if n < 0 || n > 3 {
		fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
		return
	}
	unit := s.maker[n]
	s.Hire(unit, player)
	//fmt.Println("Done! Wonna hire more?\n 1: Yes 2: No")
}

func (s Shop) Hire(unit Units.Unit, p *Player) {
	cost := unit.GetCost()
	if cost > p.Gold {
		fmt.Println("Oh! Its too expensive for you, buddy!")
		return
	}
	p.Army = append(p.Army, unit)
	p.Gold = p.Gold - cost
	fmt.Println("Done!")
}
