package monsters

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/items"
	"github.com/joram/game-server/utils"
	"log"
	"math"
	"math/rand"
)

type BaseMonster struct {
	*utils.Object
	MaxHealth   int
	Health      int `json:"health"`
	MinDamage   int `json:"min_damage"`
	MaxDamage   int `json:"max_damage"`
	IsAttacking bool
}

func (m BaseMonster) IsDead() bool {
	return m.Health <= 0
}

func (m BaseMonster) HealthBar() *string {
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

func (m BaseMonster) AsString() string {
	originalImages := m.Images
	m.Images = m.GetImages()
	jsonString, err := json.Marshal(m)
	m.Images = originalImages

	if err != nil {
		log.Println("write:", err)
	}
	return string(jsonString)
}

func (m BaseMonster) GetImages() []string {

	if m.IsDead(){
		return []string{"/images/dc-misc/blood_red.png"}
	}

	hb := m.HealthBar()
	images := m.Images
	if hb != nil {
		images = append(images, *hb)
	}
	return images
}

func (m *BaseMonster) InitialItems(minItems, maxItems int, options map[int]items.ItemType) []items.Item {
	total := 0
	for v, _ := range options {
		total += v
	}
	numItems := rand.Intn(maxItems-minItems+1)+minItems

	var initialItems []items.Item
	for i :=0; i<numItems; i++ {
		r := rand.Intn(total)
		s := 0
		for v, itemType := range options {
			s += v
			if r < s {
				item := itemType.NewInstance(m.ID)
				m.PickUpItem(&item)
				initialItems = append(initialItems, item)
				break
			}
		}
	}
	return initialItems
}

func (m *BaseMonster) Attack(target utils.BaseMonsterInterface) {
	damage := rand.Intn(m.MaxDamage - m.MinDamage) + m.MinDamage
	target.TakeDamage(damage, m)
}

func (m *BaseMonster) TakeDamage(damage int, attacker utils.BaseMonsterInterface) {
	m.Health -= damage
	m.Solid = false
	m.Broadcast()
	fmt.Printf("%s[%d] took %d damage from %s[%d]\n", m.Type, m.ID, damage, attacker.GetType(), attacker.GetID())
	if m.IsDead() {
		fmt.Printf("%s[%d] died\n", m.Type, m.ID)
		m.DropAllItems()
	}
}

var s = items.DullSword.NewInstance( -32)
var a = items.LeatherArmour.NewInstance( -32)
var h = items.LeatherHelmet.NewInstance( -32)
var ITEMS = map[int]*items.Item{
	s.ID: &s,
	a.ID: &a,
	h.ID: &h,
}

func (m BaseMonster) GetBackpackItems() []*items.Item {
	var myItems []*items.Item
	for _, item := range ITEMS {
		if item.OwnerID == m.ID {
			myItems = append(myItems, item)
		}
	}
	return myItems
}

func (m BaseMonster) EquipItem(id int) *items.Item {
	fmt.Println("equipping",id)
	ITEMS[id].OwnerID = m.ID
	ITEMS[id].IsCarried = true
	ITEMS[id].IsEquipped = true
	ITEMS[id].EquippedSlot = ITEMS[id].Slot
	return ITEMS[id]
}

func (m BaseMonster) UnequipItem(id int) *items.Item {
	fmt.Println("unequipping",id)
	ITEMS[id].OwnerID = m.ID
	ITEMS[id].IsCarried = true
	ITEMS[id].IsEquipped = false
	ITEMS[id].EquippedSlot = -1
	return ITEMS[id]
}

func (m BaseMonster) DropAllItems() {
	for _, item := range m.GetBackpackItems() {
		ITEMS[item.ID].OwnerID = -1
		ITEMS[item.ID].IsCarried = false
		ITEMS[item.ID].IsEquipped = false
		ITEMS[item.ID].X = m.X
		ITEMS[item.ID].Y = m.Y
		for _, c := range utils.ObjectClients {
			c.SendBackpackItem(ITEMS[item.ID])
		}
		fmt.Printf("%s[%d] dropped %s[%d]\n", m.Type, m.ID, ITEMS[item.ID].Name, item.ID)
	}

}

func (m BaseMonster) DropItem(id int) *items.Item {
	fmt.Println("dropping",id)
	ITEMS[id].IsEquipped = false
	ITEMS[id].IsCarried = false
	ITEMS[id].OwnerID = -1
	ITEMS[id].EquippedSlot = -1
	ITEMS[id].X = m.X
	ITEMS[id].Y = m.Y
	return ITEMS[id]
}

func (m BaseMonster) PickUpItem(item *items.Item)  {
	item.IsEquipped = false
	item.IsCarried = true
	item.OwnerID = m.ID
	item.EquippedSlot = -1
	item.X = m.X
	item.Y = m.Y
	ITEMS[item.ID] = item
	for _, c := range utils.ObjectClients {
		if c.Player.GetID() == m.ID {
			c.SendBackpackItem(item)
			break
		}
	}
	fmt.Printf("%s[%d] picked up %s[%d]\n", m.Type, m.ID, item.Name, item.ID)
}

func (m BaseMonster) GetType() string {
	return m.Type
}

func (m *BaseMonster) Broadcast(){
	for _, client := range utils.ObjectClients {
		client.UpdateMonster(m)
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
