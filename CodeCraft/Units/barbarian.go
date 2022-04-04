package Units

import (
	"awesomeProject1/CodeCraft/Player"
	"fmt"
)

type Barbarian struct {
	Class  string
	HP     int64
	Armor  int64
	Damage int64
	Cost   int64
	//Weapon string
}

func NewBarbarian() Unit {
	return &Barbarian{"Warrior", 100, 0, 50, 100}
}

func (b Barbarian) Attack(enemy Unit) bool {
	atk := b.Damage
	return enemy.GetHit(atk)
}

func (b *Barbarian) GetHit(atk int64) bool {
	b.HP = b.HP + b.Armor - atk
	b.Rage()
	if b.HP > 0 {
		return true
	}
	return false
}

func (b Barbarian) String() string {
	return fmt.Sprintf("This is %s with stats: HP = %d, atk = %d\n", b.Class, b.HP, b.Damage)
}
func (b *Barbarian) Rage() {
	if b.HP <= 50 {
		b.Damage = 75
	}
}

func (b Barbarian) Price(p *Player.Player) bool {
	if p.Gold >= b.Cost {
		p.Gold = p.Gold - b.Cost
		fmt.Println("Done!")
		return true
	} else {
		fmt.Println("Oh! Its too expensive for you, buddy!")
		return false
	}
}

//Knight := Unit{"Warrior", 100, 25, 40, 100}
//Hunter := Unit{"Archer", 75, 0, 75, 100}
