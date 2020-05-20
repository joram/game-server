package utils

import (
	"github.com/aquilax/go-perlin"
)


var R = createPerlin(0)
var G = createPerlin(1)
var B = createPerlin(2)

type Pixel struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}


func GetPixel(x, y int) Pixel {
	return Pixel{
		X: x,
		Y: y,
		R: getValue(R, x, y, 10),
		G: getValue(G, x, y, 10),
		B: getValue(B, x, y, 10),
	}
}

func GetPixels(x1, y1, x2, y2 int) []Pixel {
	pixels := []Pixel{}
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			pixels = append(pixels, GetPixel(x, y))
		}
	}
	return pixels
}


func createPerlin(seed int64) *perlin.Perlin {
	alpha := 10.0
	beta := 10.0
	n := 5
	p := perlin.NewPerlin(alpha, beta, n, seed)
	return p
}

func getValue(p *perlin.Perlin, x, y int, zoom float64) int {
	value := p.Noise2D(float64(x)/zoom, float64(y)/zoom)*255+255/2
	return int(value)
}
