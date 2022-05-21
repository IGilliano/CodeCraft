package shop

import (
	"awesomeProject1/code_craft/player"
	"awesomeProject1/code_craft/units"
	"awesomeProject1/code_craft/utils"
	"fmt"
)

type Shop struct {
	Roster []units.Unit
	maker  map[int64]units.Unit
}

func NewShop() Shop {
	New := Shop{[]units.Unit{units.NewBarbarian(), units.NewHunter()}, make(map[int64]units.Unit)}

	New.maker[1] = units.NewBarbarian()
	New.maker[2] = units.NewHunter()
	New.maker[3] = units.NewKnight()

	return New
}

func (s Shop) ShopMenu(player *player.Player) {
	for {
		fmt.Println("What do you want to do?\n 1: Hire\n 2: Show roster\n 3: My army\n 0: Exit")
		slct := utils.ScanInt()
		switch slct {
		case 1:
			player.Ready = true
			s.SelectUnit(player)
		case 2:
			s.ShowRoster()
		case 3:
			fmt.Println(player.Army)
		case 0:
			return
		default:
			fmt.Println("I dont understand! Please, select one of options!")
		}
	}
}

func (s Shop) ShowRoster() {
	fmt.Println(s.Roster)
}

func (s Shop) SelectUnit(player *player.Player) {
	var n int64 = 1
	for n != 0 {
		fmt.Println("Which one do you want to hire?\n 1: Barbarian\n 2: Hunter\n 3: Knight\n 0: Exit")
		n = utils.ScanInt()

		if n < 0 || n > 3 {
			fmt.Println("I dont get it. You need to choose one of my offers. Please, try again")
			return
		} else if n == 0 {
			return
		}
		unit := s.maker[n]
		s.Hire(unit, player)
		fmt.Println("Done! Wonna hire more?\n 1: Yes 2: No")
		n = utils.ScanInt()
		if n == 2 {
			n = 0
			fmt.Println("Okay then. Come back later!")
		}
	}
}

func (s Shop) Hire(unit units.Unit, p *player.Player) {
	cost := unit.GetCost()
	if cost > p.Gold {
		fmt.Println("Oh! Its too expensive for you, buddy!")
		return
	}
	p.Army = append(p.Army, unit)
	p.Gold = p.Gold - cost
	fmt.Println("Done!")
}
