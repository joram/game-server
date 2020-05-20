package monsters

import (
	"fmt"
	"github.com/joram/game-server/utils"
	"math"
	"time"
)


type Kobold struct {
	*BaseMonster
}

func NewKobold(id, x,y int) Kobold {
	k := Kobold{
		&BaseMonster{
			Object: &utils.Object{
				ID:    id,
				X:     x,
				Y:     y,
				Type:  "kobold",
				Solid: true,
				Images: []string{"/images/dc-mon/kobold.png"},
			},
			Health: 20,
			MinDamage: 1,
			MaxDamage: 3,
		},
	}
	go k.move()
	return k
}

func (k *Kobold) nearestPlayer() (*Player, float64) {
	var nearest *Player
	nearestDistance := -1.0
	for _, p := range PLAYERS {
		x1,y1 := k.GetLocation()
		x2,y2 := p.GetLocation()
		a := math.Abs(float64(x1-x2))
		b := math.Abs(float64(y1-y2))
		distance := math.Sqrt(a*a + b*b)
		if nearest == nil || distance < nearestDistance {
				nearest = p
				nearestDistance = distance
		}
	}
	return nearest, nearestDistance
}


func (k *Kobold) isSolid(x,y int) bool {
	return utils.GetPixel(x,y).G > 180
}

func (k *Kobold) move() {
	for {
		time.Sleep(time.Second)
		k.moveToNearestPlayer(6)
	}
}
func (k *Kobold) moveToNearestPlayer(maxDistance float64) bool {
	player, distance := k.nearestPlayer()

	x := k.X
	y := k.Y
	if player != nil && math.Round(distance) != 1 && distance <= maxDistance {
		fmt.Printf("nearest player is %v at %v\n", player.ID, distance)

		if player.X < k.X {
			x -= 1
		} else if player.X > k.X {
			x += 1
		} else if player.Y < k.Y {
			y -= 1
		} else if player.Y > k.Y {
			y += 1
		}

		if k.isSolid(x,y) {
			k.UpdateDeltaLocation(0, 0)
			return false
		}

		k.UpdateLocation(x,y)
		return true
	}

	k.UpdateDeltaLocation(0, 0)
	return false
}
