package game

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/monsters"
	"github.com/joram/game-server/utils"
	"math"
	"net/http"
)

var CHUNK_SIZE = 10
var CHUNKS = loadChunks()

type Chunk struct {
	X int `json:"x"`
	Y int `json:"y"`
	Size int `json:"size"`
	//ServeObjects []Object `json:"objects"`
	Pixels []utils.Pixel `json:"pixels"`
}

func coordToChunkCoord(x, y int) (int, int) {
	x = int(math.Floor(float64(x)/float64(CHUNK_SIZE)))* CHUNK_SIZE
	y = int(math.Floor(float64(y)/float64(CHUNK_SIZE)))* CHUNK_SIZE
	return x, y
}

func coordToChunkKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func randomAt(x,y, seed, chance int) bool {
	n := (x*y+y)*seed
	return n % chance == 0
}

func loadChunks() map[string]Chunk {
	chunks := map[string]Chunk{}
	objects := monsters.LoadMonsters()

	for _, o := range objects {
		x, y := coordToChunkCoord(o.GetLocation())
		x2 := x + CHUNK_SIZE
		y2 := y + CHUNK_SIZE
		chunkKey := coordToChunkKey(x, y)
		chunk, ok := chunks[chunkKey]

		o.UpdateLocation(x%CHUNK_SIZE, y%CHUNK_SIZE)

		// new chunk
		if !ok {
			chunks[chunkKey] = Chunk{
				X:    x,
				Y:    y,
				Size: CHUNK_SIZE,
				//ServeObjects: []Object{o},
				Pixels: utils.GetPixels(x,y,x2,y2),
			}
			continue
		}

		// existing chunk
		//chunk.ServeObjects = append(chunks[chunkKey].ServeObjects, o)
		chunks[chunkKey] = chunk
	}

	return chunks
}

func ServeChunks(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	x := utils.GetParam(r, "x")
	y := utils.GetParam(r, "y")
	fmt.Printf("chunk: (%d, %d)\n", x, y)

	chunkKey := coordToChunkKey(x,y)
	chunk, ok := CHUNKS[chunkKey]
	if ok {
		json.NewEncoder(w).Encode(chunk)
	} else {
		chunk = Chunk{
			X:      x,
			Y:      y,
			Size:   CHUNK_SIZE,
			Pixels: utils.GetPixels(x,y,x+CHUNK_SIZE,y+CHUNK_SIZE),
		}
		json.NewEncoder(w).Encode(chunk)
	}

}