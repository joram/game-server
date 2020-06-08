package items

import (
	"github.com/joram/game-server/ids"
)

const (
	SLOT1 = 1
	SLOT2 = 2
	SLOT3 = 3
	SLOT4 = 4
	SLOT5 = 5
	SLOT6 = 6
)
type ItemType struct {
	Name string `json:"name"`
	EquippedImage string `json:"equipped_image"`
	DroppedImage string `json:"dropped_image"`
	AllowedSlot int `json:"allowed_slot"`
	MinDamage int `json:"min_damage"`
	MaxDamage int `json:"max_damage"`
}

var SWORD1 = ItemType{
	"sword",
	"/images/player/hand1/short_sword.png",
	"/images/item/weapon/long_sword1.png",
	SLOT3,
	1,
	3,
}

var SWORD2 = ItemType{
	"sword",
	"/images/player/hand1/short_sword2.png",
	"/images/item/weapon/long_sword2.png",
	SLOT3,
	2,
	4,
}


func (it *ItemType) NewInstance(ownerID int) Item {
	return Item{
		ItemType:  it,
		ID:        ids.NextID("item"),
		OwnerID:   ownerID,
		IsCarried: true,
	}
}

type Item struct {
	*ItemType
	X int `json:"x"`
	Y int `json:"y"`
	ID int `json:"id"`
	IsEquipped bool  `json:"is_equipped"`
	IsCarried bool `json:"is_carried"`
	OwnerID int `json:"owner_id"`
	EquippedSlot int `json:"equipped_slot"`
}
