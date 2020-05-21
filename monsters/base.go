package monsters

import (
	"github.com/joram/game-server/utils"
)


func LoadMonsters() []utils.ObjectInterface {
	var objects []utils.ObjectInterface
	k := NewKobold(-5,-5)
	objects = append(objects, &k)
	return objects
}

