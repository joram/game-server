package monsters

import (
	"github.com/joram/game-server/utils"
)


type Player struct {
	*BaseMonster
}

var PLAYERS []*Player


func NewPlayer(x, y int) Player {
	p := Player{
		&BaseMonster{
			Object: &utils.Object{
				ID:    utils.NextID(),
				X:     x,
				Y:     y,
				Type:  "player",
				Solid: true,
				Images: []string{
					"/images/player/base/gnome_m.png",
					"/images/player/legs/leg_armor01.png",
					"/images/player/boots/short_brown.png",
					"/images/player/body/aragorn.png",
					"/images/player/head/hood_ybrown.png",
				},
			},
			MaxHealth: 20,
			Health: 20,
			MinDamage: 1,
			MaxDamage: 3,
		},
	}
	p.register()
	return p
}

func (p Player) GetLocation() (x,y int){
	return p.BaseMonster.Object.GetLocation()
}

func (p Player) UpdateLocation(x,y int){
	p.BaseMonster.Object.UpdateLocation(x,y)
}

func (p Player) GetID() int {
	return p.BaseMonster.Object.GetID()
}

func (p Player) AsString() string {
	return p.BaseMonster.AsString()
}

func (p *Player) register(){
	PLAYERS = append(PLAYERS, p)
}

func (p *Player) Unregister(){
	var players []*Player
	for _, pp := range PLAYERS {
		if pp == p { continue }
		players = append(players, pp)
	}
	PLAYERS = players
	p.Broadcast()
}
