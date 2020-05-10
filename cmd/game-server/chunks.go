package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

var CHUNK_SIZE = 10
var CHUNKS = loadChunks()

type Chunk struct {
	X int `json:"x"`
	Y int `json:"y"`
	Size int `json:"size"`
	Objects []Object `json:"objects"`
}

func coordToChunkCoord(x, y int) (int, int) {
	x = int(math.Floor(float64(x/CHUNK_SIZE)))
	y = int(math.Floor(float64(y/CHUNK_SIZE)))
	return x, y
}

func coordToChunkKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func loadChunks() map[string]Chunk {
	chunks := map[string]Chunk{}
	objects := LoadObjects()

	for _, o := range objects {
		x, y := coordToChunkCoord(o.X, o.Y)
		chunkKey := coordToChunkKey(x, y)
		chunk, ok := chunks[chunkKey]

		o.X = o.X%CHUNK_SIZE
		o.Y = o.Y%CHUNK_SIZE

		// new chunk
		if !ok {
			chunks[chunkKey] = Chunk{
				X:       x,
				Y:       y,
				Size:    CHUNK_SIZE,
				Objects: []Object{o},
			}
			continue
		}

		// existing chunk
		chunk.Objects = append(chunks[chunkKey].Objects, o)
		chunks[chunkKey] = chunk
	}

	return chunks
}

func serveChunks(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	x := getParam(r, "x")
	y := getParam(r, "y")
	fmt.Printf("chunk: (%d, %d)\n", x, y)

	chunkKey := coordToChunkKey(x,y)
	chunk, ok := CHUNKS[chunkKey]
	if ok {
		json.NewEncoder(w).Encode(chunk)
	} else {
		json.NewEncoder(w).Encode(Chunk{X:x, Y:y, Size: CHUNK_SIZE, Objects: []Object{}})
	}

}