package utils

import "github.com/joram/game-server/items"

type BaseMonsterInterface interface {
	AsString() string
	GetID() int
	GetType() string

	GetLocation() (x,y int)
	UpdateLocation(x,y int)
	UpdateDeltaLocation(x,y int)

	TakeDamage(damage int, attacker BaseMonsterInterface)
	IsDead() bool

	GetBackpackItems() []*items.Item
	EquipItem(id int) *items.Item
}

