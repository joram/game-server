package monsters

import (
	"fmt"
	"github.com/joram/game-server/utils"
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
