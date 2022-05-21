package units

type Unit interface {
	Attack(enemy Unit) bool
	GetHit(int64) bool
	GetCost() int64
}
