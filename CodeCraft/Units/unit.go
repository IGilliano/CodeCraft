package Units

type Unit interface {
	Attack(enemy Unit) bool
	GetHit(int64) bool
}
