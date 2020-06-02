package utils

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"math"
	"os"
)


type Town struct {
	Pixels []Pixel
	X      int
	Y      int
	Width int
	Height int
}

var Towns = []Town{LoadTown()}

func IsTown(x,y int) bool {
	for _, t := range Towns {
		if t.Contains(x, y) {
			return true
		}
	}
	return false
}



func (t *Town) Contains(x,y int) bool {
	if x < t.X { return false }
	if y < t.Y { return false }
	if x > t.X+t.Width-1 { return false }
	if y > t.Y+t.Height-1 { return false }
	return true
}

func (t *Town) Pixel(x,y int) Pixel {
	for _, p := range t.Pixels {
		if p.X+t.X == x && p.Y+t.Y == y {
			pixel := Pixel{
				X:x,
				Y:y,
				R:p.R,
				G:p.G,
				B:p.B,
			}
			return pixel
		}
	}
	fmt.Println("unknown pixel")
	return Pixel{}
}

func LoadTown() Town {
	t := Town{
		X: -10,
		Y: -10,
		Width: 20,
		Height: 20,
	}
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("./static/towns/town1.png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer file.Close()

	t.Pixels, t.Width, t.Height = getPixels(file)
	t.Width += 1
	t.Height += 1
	return t
}

func getPixels(file io.Reader) ([]Pixel, int, int) {
	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := []Pixel{}
	maxX := 0
	maxY := 0
	for y := 0; y < height; y++ {
		maxY = int(math.Max(float64(maxY), float64(y)))
		for x := 0; x < width; x++ {
			maxX = int(math.Max(float64(maxX), float64(x)))
			r,g,b, _ := img.At(x, y).RGBA()
			pixels = append(pixels, Pixel{x,y,int(r),int(g),int(b)})
		}
	}

	return pixels, maxX, maxY
}