package monsters

import (
	"github.com/joram/game-server/utils"
)

type BaseMonster struct {
	*utils.Object
	Health int `json:"health"`
	MinDamage int `json:"min_damage"`
	MaxDamage int `json:"max_damage"`
}

func LoadMonsters() []utils.ObjectInterface {
	var objects []utils.ObjectInterface
	k := NewKobold(-5,-5)
	objects = append(objects, &k)
	return objects
}