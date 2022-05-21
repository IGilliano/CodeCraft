package units

import "fmt"

type Knight struct {
	Class  string
	HP     int64
	Armor  int64
	Damage int64
	Cost   int64
	//Weapon string
}

func NewKnight() Unit {
	return &Knight{"Warrior", 50, 50, 50, 100}
}

func (k Knight) Attack(enemy Unit) bool {
	atk := k.Damage
	return enemy.GetHit(atk)
}

func (k *Knight) GetHit(atk int64) bool {
	k.HP = k.HP + k.Armor - atk
	if k.HP > 0 {
		return true
	}
	return false
}

func (k Knight) String() string {
	return fmt.Sprintf("Barbarian (HP = %d, atk = %d)", k.HP, k.Damage)
}

func (k *Knight) Rage() {
	if k.HP <= 50 {
		k.Damage = 75
		fmt.Println("Barbarian is getting furious!")
	}
}

func (k Knight) GetCost() int64 {
	return k.Cost
}
