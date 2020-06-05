package monsters

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/items"
	"github.com/joram/game-server/utils"
	"log"
	"math"
)

type BaseMonster struct {
	*utils.Object
	MaxHealth   int
	Health      int `json:"health"`
	MinDamage   int `json:"min_damage"`
	MaxDamage   int `json:"max_damage"`
	IsAttacking bool
}

func (m *BaseMonster) IsDead() bool {
	return m.Health <= 0
}

func (m *BaseMonster) HealthBar() *string {
	if m.Health == m.MaxHealth || m.IsDead() {
		return nil
	}

	s := []string{
		"/images/dc-misc/mdam_almost_dead.png",
		"/images/dc-misc/mdam_severely_damaged.png",
		"/images/dc-misc/mdam_heavily_damaged.png",
		"/images/dc-misc/mdam_moderately_damaged.png",
		"/images/dc-misc/mdam_lightly_damaged.png",
		"/images/dc-misc/blood_red.png",
	}
	if m.Health <= 0 {
		return &s[0]
	}

	ratio := float64(m.Health)/float64(m.MaxHealth)
	if ratio < 0.2 {
		return &s[0]
	}
	if ratio < 0.4 {
		return &s[1]
	}
	if ratio < 0.6 {
		return &s[2]
	}
	if ratio < 0.8 {
		return &s[3]
	}
	if ratio < 1 {
		return &s[4]
	}
	return &s[5]
}

func (m *BaseMonster) AsString() string {
	originalImages := m.Images
	hb := m.HealthBar()
	if m.IsDead(){
		m.Images = []string{"/images/dc-misc/blood_red.png"}
	}
	if hb != nil {
		m.Images = append(m.Images, *hb)
	}
	jsonString, err := json.Marshal(m)

	m.Images = originalImages
	if err != nil {
		log.Println("write:", err)
	}
	return string(jsonString)
}

func (m *BaseMonster) TakeDamage(damage int, attacker utils.BaseMonsterInterface) {
	m.Health -= damage
	m.Solid = false
	m.Broadcast()
	fmt.Printf("%s[%d] took %d damage from %s[%d]\n", m.Type, m.ID, damage, attacker.GetType(), attacker.GetID())
	if m.IsDead() {
		fmt.Printf("%s[%d] died\n", m.Type, m.ID)
	}
}

var s = items.SWORD.NewInstance(0,0,0,false,true, -32, -1)
var ITEMS = map[int]*items.Item{
	0: &s,
}

func (m *BaseMonster) GetBackpackItems() []*items.Item {
	var myItems []*items.Item
	for _, item := range ITEMS {
		if item.OwnerID == m.ID {
			myItems = append(myItems, item)
		}
	}
	return myItems
}

func (m *BaseMonster) EquipItem(id int) *items.Item {
	fmt.Println("equipping",id)
	ITEMS[id].OwnerID = m.ID
	ITEMS[id].IsCarried = true
	ITEMS[id].IsEquipped = true
	ITEMS[id].EquippedSlot = ITEMS[id].AllowedSlot
	return ITEMS[id]
}

func (m *BaseMonster) UnequipItem(id int) *items.Item {
	fmt.Println("unequipping",id)
	ITEMS[id].OwnerID = m.ID
	ITEMS[id].IsCarried = true
	ITEMS[id].IsEquipped = false
	ITEMS[id].EquippedSlot = -1
	return ITEMS[id]
}

func (m *BaseMonster) DropItem(id int) *items.Item {
	fmt.Println("dropping",id)
	ITEMS[id].IsEquipped = false
	ITEMS[id].IsCarried = false
	ITEMS[id].OwnerID = -1
	ITEMS[id].EquippedSlot = -1
	ITEMS[id].X = m.X
	ITEMS[id].Y = m.Y
	return ITEMS[id]
}

func (m *BaseMonster) GetType() string {
	return m.Type
}

func (m *BaseMonster) Broadcast(){
	for _, client := range utils.ObjectClients {
		client.UpdateObject(m)
	}
}

func (m *BaseMonster) nearestPlayer() (*Player, float64) {
	var nearest *Player
	nearestDistance := -1.0
	for _, p := range PLAYERS {
		if p.IsDead() {
			continue
		}

		x1,y1 := m.GetLocation()
		x2,y2 := p.GetLocation()
		a := math.Abs(float64(x1-x2))
		b := math.Abs(float64(y1-y2))
		distance := math.Sqrt(a*a + b*b)
		if nearest == nil || distance < nearestDistance {
			nearest = p
			nearestDistance = distance
		}
	}
	return nearest, nearestDistance
}


func (m *BaseMonster) isSolid(x,y int) bool {
	return utils.GetPixel(x,y).G > 180
}

func (m *BaseMonster) moveToNearestPlayer(maxDistance float64) *Player {
	player, distance := m.nearestPlayer()

	x := m.X
	y := m.Y
	if math.Round(distance) == 1 {
		return player
	}

	if player != nil && distance <= maxDistance {
		if player.X < m.X {
			x -= 1
		} else if player.X > m.X {
			x += 1
		} else if player.Y < m.Y {
			y -= 1
		} else if player.Y > m.Y {
			y += 1
		}

		if m.isSolid(x,y) {
			return nil
		}

		if utils.IsTown(x,y) {
			return nil
		}

		m.UpdateLocation(x,y)
	}
	return nil
}
