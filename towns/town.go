package towns

import (
	"fmt"
	"github.com/joram/game-server/utils"
	"image"
	"image/png"
	"io"
	"os"
)


type Town struct {
	Pixels []utils.Pixel
	X      int
	Y      int
	Width int
	Height int
}

func (t *Town) Contains(x,y int) bool {
	if x < t.X { return false }
	if y < t.Y { return false }
	if x > t.X+t.Width-1 { return false }
	if y > t.Y+t.Height-1 { return false }
	return true
}

func (t *Town) Pixel(x,y int) utils.Pixel {
	for _, p := range t.Pixels {
		if p.X+t.X == x && p.Y+t.Y == y {
			pixel := utils.Pixel{
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
	return utils.Pixel{}
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

	file, err := os.Open("./towns/images/town1.png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}
	defer file.Close()

	pixels := getPixels(file)
	t.Pixels = pixels
	return t
}


func getPixels(file io.Reader) []utils.Pixel {
	img, _, _ := image.Decode(file)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := []utils.Pixel{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r,g,b, _ := img.At(x, y).RGBA()
			pixels = append(pixels, utils.Pixel{x,y,int(r),int(g),int(b)})
		}
	}

	return pixels
}