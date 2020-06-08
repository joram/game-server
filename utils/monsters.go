package utils

import "github.com/joram/game-server/items"

type BaseMonsterInterface interface {
	AsString() string
	GetID() int
	GetType() string
	GetImages() []string

	GetLocation() (x,y int)
	UpdateLocation(x,y int)
	UpdateDeltaLocation(x,y int)

	Attack(BaseMonsterInterface)
	TakeDamage(damage int, attacker BaseMonsterInterface)
	IsDead() bool

	GetBackpackItems() []*items.Item
	EquipItem(id int) *items.Item
	UnequipItem(id int) *items.Item
	DropItem(id int) *items.Item
}

