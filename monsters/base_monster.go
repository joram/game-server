package monsters

import (
	"encoding/json"
	"fmt"
	//"github.com/joram/game-server/game"
	"github.com/joram/game-server/utils"

	//"github.com/joram/game-server/game"

	//"github.com/joram/game-server/monsters"
	"log"
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

func (m BaseMonster) AsString() string {
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
	m.Broadcast()
	fmt.Printf("%s[%d] took %d damage from %s[%d]\n", m.Type, m.ID, damage, attacker.GetType(), attacker.GetID())
	if m.IsDead() {
		fmt.Printf("%s[%d] died\n", m.Type, m.ID)
	}
}

func (m *BaseMonster) GetType() string {
	return m.Type
}

func (o *BaseMonster) Broadcast(){
	for _, client := range utils.ObjectClients {
		client.UpdateObject(o)
	}
}