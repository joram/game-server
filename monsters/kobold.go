package monsters

import (
	"fmt"
	"github.com/joram/game-server/ids"
	"github.com/joram/game-server/items"
	"github.com/joram/game-server/utils"
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

	options := map[int]items.ItemType{
		10: items.DullSword,
		5: items.SharpSword,
		4: items.LeatherHelmet,
		1: items.LeatherArmour,
	}
	k.InitialItems(1,2, options)

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
			k.Attack(player)
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

