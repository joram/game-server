package server

import (
	"context"
	"github.com/joram/game-server/pb"
	"github.com/aquilax/go-perlin"
)

func (s *GameServer) GetTiles(ctx context.Context, request *pb.GetTilesRequest, ) (*pb.GetTilesResponse, error) {
	size := 20
	alpha := 10.0
	beta := 10.0
	n := 5
	zoom := 10.
	seed := int64(0)
	p := perlin.NewPerlin(alpha, beta, n, seed)

	response := pb.GetTilesResponse{}
	for x := 0; x < size; x++ {
		response.Tiles[int64(x)] = &pb.RGBARow{}
		for y := 0; y < size; y++ {
			value := p.Noise2D(float64(int64(x)+request.X)/zoom, float64(int64(y)+request.Y)/zoom)*255+255/2
			if value < 0 { value = 0 }
			if value > 255 { value = 255 }

			response.Tiles[int64(x)].Row[int64(y)] = &pb.RGBA{
				R: int64(value),
				A: 255,
			}
		}
	}

	return &response, nil
}

