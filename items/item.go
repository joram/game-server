package items

import (
	"github.com/joram/game-server/ids"
)

type ItemType struct {
	Name string `json:"name"`
	EquippedImage string `json:"equipped_image"`
	DroppedImage string `json:"dropped_image"`
	AllowedSlot int `json:"allowed_slot"`
}

var SWORD = ItemType{
	"sword",
	"/images/player/hand1/short_sword.png",
	"/images/item/weapon/long_sword1.png",
	3,
}

func (it *ItemType) NewInstance(x,y int, isEquipped, isCarried bool, ownerID, equippedSlot int) Item {
	return Item{
		ItemType:     it,
		ID:           ids.NextID("item"),
		X:            x,
		Y:            y,
		IsEquipped:   isEquipped,
		IsCarried:    isCarried,
		OwnerID:      ownerID,
		EquippedSlot: equippedSlot,
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
