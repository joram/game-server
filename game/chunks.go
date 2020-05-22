package game

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/towns"
	"github.com/joram/game-server/utils"
	"net/http"
)

var ChunkSize = 10
var ActiveChunks = map[string]Chunk{}
var Towns = []towns.Town{towns.LoadTown()}

func allMonsters() []utils.BaseMonsterInterface {
	var all []utils.BaseMonsterInterface
	for _, c := range ActiveChunks {
		all = append(all, c.Monsters...)
	}
	return all
}

type Chunk struct {
	X int `json:"x"`
	Y int `json:"y"`
	Size int `json:"size"`
	Pixels []utils.Pixel `json:"pixels"`
	Monsters []utils.BaseMonsterInterface
}

func (c *Chunk) IsSolid(x,y int) bool {
	for _, p := range c.Pixels {
		if p.X == x && p.Y == y {
			return p.G > 180
		}
	}
	return false
}

func coordToChunkKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func getChunk(x,y int) *Chunk {
	x2 := x + ChunkSize
	y2 := y + ChunkSize
	chunkKey := coordToChunkKey(x, y)
	chunk, ok := ActiveChunks[chunkKey]

	if !ok {

		pixels := utils.GetPixels(x,y,x2,y2)

		// override pixels with town data
		for i, p := range pixels {
			for _, town := range Towns {
				if town.Contains(p.X, p.Y) {
					pixels[i] = town.Pixel(p.X, p.Y)
				}
			}
		}

		chunk = Chunk{
			X:    x,
			Y:    y,
			Size: ChunkSize,
			//ServeObjects: []Object{o},
			Pixels: pixels,
		}
		chunk.Monsters = NewMonsters(x, y, x+ChunkSize, y+ChunkSize, 0.1, chunk)
		ActiveChunks[chunkKey] = chunk
	}

	return &chunk
}

func ServeChunks(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	x := utils.GetParam(r, "x")
	y := utils.GetParam(r, "y")
	chunk := getChunk(x,y)
	json.NewEncoder(w).Encode(chunk)
}