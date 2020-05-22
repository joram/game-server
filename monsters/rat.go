package monsters

import (
	"fmt"
	"github.com/joram/game-server/utils"
	"math"
	"math/rand"
	"time"
)


type Rat struct {
	*BaseMonster
}

func NewRat(x, y int) Rat {
	k := Rat{
		&BaseMonster{
			Object: &utils.Object{
				ID:    utils.NextID(),
				X:     x,
				Y:     y,
				Type:  "Rat",
				Solid: true,
				Images: []string{"/images/dc-mon/animals/rat.png"},
			},
			MaxHealth: 5,
			Health:      5,
			MinDamage:   1,
			MaxDamage:   3,
			IsAttacking: false,
		},
	}
	go k.move()
	return k
}

func (k *Rat) nearestPlayer() (*Player, float64) {
	var nearest *Player
	nearestDistance := -1.0
	for _, p := range PLAYERS {
		if p.IsDead() {
			continue
		}

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


func (k *Rat) isSolid(x,y int) bool {
	return utils.GetPixel(x,y).G > 180
}

func (k *Rat) move() {
	for {
		time.Sleep(time.Second)
		player := k.moveToNearestPlayer(6)

		// started attacking
		if !k.IsAttacking && player != nil {
			fmt.Println("Rat now attacking!")
			k.IsAttacking = true
			k.Images = []string{
				"/images/dc-mon/animals/rat.png",
				"/images/dc-misc/animated_weapon.png",
			}
			k.UpdateDeltaLocation(0, 0)

		// stopped attacking
		} else if k.IsAttacking && player == nil {
			fmt.Printf("Rat[%d] stopped attacking\n", k.ID)
			k.IsAttacking = false
			k.Images = []string{
				"/images/dc-mon/animals/rat.png",
			}
			k.UpdateDeltaLocation(0,0)
		}

		if k.IsAttacking {
			damage := rand.Intn(k.MaxDamage - k.MinDamage) + k.MinDamage
			player.TakeDamage(damage, k)
		}

	}
}

func (k *Rat) moveToNearestPlayer(maxDistance float64) *Player {
	player, distance := k.nearestPlayer()

	x := k.X
	y := k.Y
	if math.Round(distance) == 1 {
		return player
	}

	if player != nil && distance <= maxDistance {
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
			return nil
		}

		k.UpdateLocation(x,y)
	}
	return nil
}
