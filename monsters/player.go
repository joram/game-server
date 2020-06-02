package monsters

import (
	"fmt"
	"github.com/joram/game-server/db"
	"github.com/joram/game-server/utils"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)


type Player struct {
	*BaseMonster
}

var baseImages = listImages("player/base", ".png")
var legsImages = listImages("player/legs", ".png")
var bootsImages = listImages("player/boots", ".png")
var bodyImages = listImages("player/body", ".png")
var headImages = listImages("player/head", ".png")

var PLAYERS []*Player

func listImages(path, ext string) []string {
	paths := []string{}
	root := fmt.Sprintf("static/images/%s", path)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			paths = append(paths, strings.Replace(path, "static", "",1))
		}
		return nil
	})
	return paths
}

func NewPlayer(id,x,y int) Player {
	images := []string{
		baseImages[rand.Intn(len(baseImages))],
		legsImages[rand.Intn(len(legsImages))],
		bootsImages[rand.Intn(len(bootsImages))],
		bodyImages[rand.Intn(len(bodyImages))],
		headImages[rand.Intn(len(headImages))],
	}

	// always a negative number so it doesn't collide with monsters
	if id > 0 {
		id = -id
	}

	//dbPlayer := db.
	p := Player{
		&BaseMonster{
			Object: &utils.Object{
				ID:    id,
				X:     x,
				Y:     y,
				Type:  "player",
				Solid: true,
				Images: images,
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
	return p.BaseMonster.GetLocation()
}

func (p Player) UpdateLocation(x,y int){
	p.X = x
	p.Y = y
	db.UpdatePlayer(p.ID, p.X, p.Y)
	//fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	utils.BroadcastLocationChange(p, utils.ObjectClients)
}

func (p Player) UpdateDeltaLocation(x,y int){
	p.X += x
	p.Y += y
	//fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	db.UpdatePlayer(p.ID, p.X, p.Y)
	utils.BroadcastLocationChange(p, utils.ObjectClients)
}

func (p Player) GetID() int {
	return p.BaseMonster.GetID()
}

//func (p Player) GetType() string {
//	return p.BaseMonster.GetType()
//}
//
//func (p Player) IsDead() bool {
//	return p.BaseMonster.IsDead()
//}
//func (p Player) TakeDamage(damage int, attacker utils.BaseMonsterInterface) {
//	p.BaseMonster.TakeDamage(damage, attacker)
//}
//
//func (p Player) AsString() string {
//	return p.BaseMonster.AsString()
//}

func (p *Player) register(){
	PLAYERS = append(PLAYERS, p)
}

func (p *Player) Unregister(){
	var players []*Player
	for _, pp := range PLAYERS {
		if pp.ID == p.ID { continue }
		players = append(players, pp)
	}
	PLAYERS = players
	//p.Broadcast()
}
