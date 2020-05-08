package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"image/color"
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

func pixels(w http.ResponseWriter, r *http.Request){
	x1 := getParam(r, "x1")
	y1 := getParam(r, "y1")
	x2 := getParam(r, "x2")
	y2 := getParam(r, "y2")
	fmt.Printf("(%d, %d), (%d, %d)\n", x1, y1, x2, y2)
	pixels := []Pixel{}
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			pixels = append(pixels, Pixel{
				X: x,
				Y: y,
				R: getValue(R, x, y, 10),
				G: getValue(G, x, y, 10),
				B: getValue(B, x, y, 10),
			})
		}
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(pixels)
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
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

func createPerlinImage(size int, zoom float64, xOffset, yOffset int){
	f, err := os.Create(fmt.Sprintf("img_%d_%d.jpg", xOffset, yOffset))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	target := image.NewRGBA(image.Rect(0, 0, size, size))

	alpha := 10.0
	beta := 10.0
	n := 5
	seed := int64(0)
	p := perlin.NewPerlin(alpha, beta, n, seed)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			value := p.Noise2D(float64(x+xOffset)/zoom, float64(y+yOffset)/zoom)*255+255/2
			target.Set(x, y, color.RGBA{uint8(value),0,0,255})
		}
	}
	jpeg.Encode(f, target, nil)
}

