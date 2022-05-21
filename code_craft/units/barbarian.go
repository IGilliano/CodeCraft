package units

import "fmt"

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
	return fmt.Sprintf("Barbarian (HP = %d, atk = %d)", b.HP, b.Damage)
}

func (b *Barbarian) Rage() {
	if b.HP <= 50 {
		b.Damage = 75
		fmt.Println("Barbarian is getting furious!")
	}
}

func (b Barbarian) GetCost() int64 {
	return b.Cost
}
