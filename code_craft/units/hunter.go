package units

import (
	"fmt"
)

type Hunter struct {
	Class  string
	HP     int64
	Armor  int64
	Damage int64
	Cost   int64
	//Weapon string
}

func NewHunter() Unit {
	return Hunter{"Archer", 75, 0, 75, 100}
}

func (h Hunter) Attack(enemy Unit) bool {
	atk := h.Damage
	return enemy.GetHit(atk)
}

func (h Hunter) GetHit(atk int64) bool {
	h.HP = h.HP + h.Armor - atk
	if h.HP > 0 {
		return true
	}
	return false
}

func (h Hunter) String() string {
	return fmt.Sprintf("Hunter (HP = %d, atk = %d)", h.HP, h.Damage)
}

func (h Hunter) GetCost() int64 {
	return h.Cost
}
