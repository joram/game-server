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
	WEAPON = "weapon"
	ARMOUR = "armour"
	HELMET = "helmet"
)

type ItemType struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	EquippedImage string `json:"equipped_image"`
	DroppedImage  string `json:"dropped_image"`
	Slot          int    `json:"slot"`
	MinDamage     int    `json:"min_damage"`
	MaxDamage     int    `json:"max_damage"`
	AC            int    `json:"ac"`
}

var DullSword = ItemType{
	"dull sword",
	WEAPON,
	"/images/player/hand1/short_sword.png",
	"/images/item/weapon/long_sword1.png",
	SLOT3,
	1,
	3,
	0,
}

var SharpSword = ItemType{
	"sharp sword",
	WEAPON,
	"/images/player/hand1/short_sword2.png",
	"/images/item/weapon/long_sword2.png",
	SLOT3,
	2,
	4,
	0,
}

var LeatherArmour = ItemType{
	Name:          "leather armour",
	Type:          ARMOUR,
	EquippedImage: "/images/player/body/leather_armour.png",
	DroppedImage:  "/images/item/armour/leather_armour1.png",
	Slot:          SLOT4,
	AC:            2,
}

var LeatherHelmet = ItemType{
	Name:          "leather armour",
	Type:          HELMET,
	EquippedImage: "/images/player/head/bear.png",
	DroppedImage:  "/images/item/armour/headgear/elven_leather_helm.png",
	Slot:          SLOT1,
	AC:            2,
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
