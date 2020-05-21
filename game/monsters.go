package game

import (
	"github.com/joram/game-server/utils"
	"github.com/joram/game-server/monsters"
	"math/rand"
)

func NewMonsters(minX, minY, maxX, maxY, count int, chunk Chunk) []utils.BaseMonsterInterface {
	var m []utils.BaseMonsterInterface
	for i:=0; i<count; i++ {
		isSolid := true
		x := 0
		y := 0
		for isSolid == true {
			x = rand.Intn(maxX-minX)+minX
			y = rand.Intn(maxY-minY)+minY
			isSolid = chunk.IsSolid(x, y)
		}

		k := monsters.NewKobold(x,y)
		m = append(m, k)
	}
	return m
}

