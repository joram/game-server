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

func (r *Rat) move() {
	for {
		time.Sleep(time.Second)
		player := r.moveToNearestPlayer(6)

		// started attacking
		if !r.IsAttacking && player != nil {
			fmt.Println("Rat now attacking!")
			r.IsAttacking = true
			r.Images = []string{
				"/images/dc-mon/animals/rat.png",
				"/images/dc-misc/animated_weapon.png",
			}
			r.Broadcast()

		// stopped attacking
		} else if r.IsAttacking && player == nil {
			fmt.Printf("Rat[%d] stopped attacking\n", r.ID)
			r.IsAttacking = false
			r.Images = []string{
				"/images/dc-mon/animals/rat.png",
			}
			r.Broadcast()
		}

		// attack
		if r.IsAttacking {
			damage := rand.Intn(r.MaxDamage - r.MinDamage) + r.MinDamage
			player.TakeDamage(damage, r)
		}

		// die
		if r.IsDead() {
			r.Images = []string{
				"/images/dc-misc/blood_red.png",
			}
			r.Broadcast()
			fmt.Printf("%s[%d] died!\n", r.Type, r.ID)
			return
		}
	}
}
