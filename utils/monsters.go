package utils

type BaseMonsterInterface interface {
	IsDead() bool
	AsString() string
	GetID() int
	GetType() string
	GetLocation() (x,y int)
	UpdateLocation(x,y int)
	UpdateDeltaLocation(x,y int)
	TakeDamage(damage int, attacker BaseMonsterInterface)
}

