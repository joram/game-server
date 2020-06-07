package monsters

import (
	"fmt"
	"github.com/joram/game-server/ids"
	"github.com/joram/game-server/items"
	"github.com/joram/game-server/utils"
	"math/rand"
	"time"
)


type Kobold struct {
	*BaseMonster
}

func NewKobold(x, y int) Kobold {
	k := Kobold{
		&BaseMonster{
			Object: &utils.Object{
				ID:     ids.NextID("monster"),
				X:      x,
				Y:      y,
				Type:   "kobold",
				Solid:  true,
				Images: []string{"/images/dc-mon/kobold.png"},
			},
			MaxHealth: 20,
			Health:      20,
			MinDamage:   1,
			MaxDamage:   3,
			IsAttacking: false,
		},
	}
	sword := items.SWORD.NewInstance(0,0,false,true, k.ID, -1)
	k.PickUpItem(&sword)
	go k.move()
	return k
}

func (k *Kobold) move() {
	for {
		time.Sleep(time.Second)

		player := k.moveToNearestPlayer(6)

		// started attacking
		if !k.IsAttacking && player != nil {
			fmt.Println("kobold now attacking!")
			k.IsAttacking = true
			k.Images = []string{
				"/images/dc-mon/kobold.png",
				"/images/dc-misc/animated_weapon.png",
			}
			k.UpdateDeltaLocation(0, 0)

		// stopped attacking
		} else if k.IsAttacking && player == nil {
			fmt.Printf("kobold[%d] stopped attacking\n", k.ID)
			k.IsAttacking = false
			k.Images = []string{
				"/images/dc-mon/kobold.png",
			}
			k.UpdateDeltaLocation(0,0)
		}

		// attack
		if k.IsAttacking {
			damage := rand.Intn(k.MaxDamage - k.MinDamage) + k.MinDamage
			player.TakeDamage(damage, k)
		}

		// die
		if k.IsDead() {
			k.Images = []string{
				"/images/dc-misc/blood_red.png",
			}
			k.Broadcast()
			fmt.Printf("%s[%d] died!\n", k.Type, k.ID)
			return
		}

	}
}

