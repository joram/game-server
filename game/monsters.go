package game

import (
	"github.com/joram/game-server/utils"
	"github.com/joram/game-server/monsters"
	"math"
	"math/rand"
)

func NewMonsters(minX, minY, maxX, maxY int, count float64, chunk Chunk) []utils.BaseMonsterInterface {
	var m []utils.BaseMonsterInterface

	if count < 1 {
		x := int(100*count)
		if rand.Intn(100) > x {
			return m
		}
	}

	for i:=0; i<int(math.Ceil(count)); i++ {
		isSolid := true
		x := 0
		y := 0
		for isSolid == true {
			x = rand.Intn(maxX-minX)+minX
			y = rand.Intn(maxY-minY)+minY
			isSolid = chunk.IsSolid(x, y)
		}

		if rand.Intn(2) == 0 {
			k := monsters.NewKobold(x,y)
			m = append(m, k)
		} else {
			k := monsters.NewRat(x,y)
			m = append(m, k)
		}

	}
	return m
}

