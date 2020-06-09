package monsters

import (
	"encoding/json"
	"fmt"
	"github.com/joram/game-server/db"
	"github.com/joram/game-server/utils"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)


type Player struct {
	*BaseMonster
	//*utils.Object
	//MaxHealth   int
	//Health      int `json:"health"`
	//MinDamage   int `json:"min_damage"`
	//MaxDamage   int `json:"max_damage"`
	//IsAttacking bool

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
		BaseMonster: &BaseMonster{

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

func (p Player) Attack(target utils.BaseMonsterInterface) {
	minDamage := 0
	maxDamage := 0
	for _, item := range p.GetBackpackItems() {
		if item.IsEquipped {
			minDamage += item.MinDamage
			maxDamage += item.MaxDamage
		}
	}

	damage := 1
	if !(minDamage == 0 && maxDamage == 0) {
		damage = rand.Intn(maxDamage-minDamage) + minDamage
	}

	target.TakeDamage(damage, p)
}

func (p Player) TakeDamage(damage int, attacker utils.BaseMonsterInterface) {
	ac := 0
	for _, item := range p.GetBackpackItems() {
		if item.IsEquipped {
			ac += item.AC
		}
	}

	d := float64(damage) - float64(damage)*float64(ac)/10.0
	dd := int(math.Floor(d))
	if dd == 0 && d > 0 {
		dd = 1
	}
	p.Health -= dd
	p.Solid = false
	p.Broadcast()
	fmt.Printf("%s[%d] took %d damage from %s[%d]\n", p.Type, p.ID, damage, attacker.GetType(), attacker.GetID())
	if p.IsDead() {
		fmt.Printf("%s[%d] died\n", p.Type, p.ID)
		p.DropAllItems()
	}
}


func (p *Player) Broadcast(){
	for _, client := range utils.ObjectClients {
		client.UpdateMonster(p)
	}
}

func (p Player) AsString() string {
	originalImages := p.Images
	p.Images = p.GetImages()
	jsonString, err := json.Marshal(p)
	p.Images = originalImages
	if err != nil {
		log.Println("write:", err)
	}
	return string(jsonString)
}

func (p Player) GetImages() []string {
	if p.IsDead(){
		return []string{"/images/dc-misc/blood_red.png"}
	}
	images := []string{"/images/player/base/human_m.png"}
	for _, item := range p.GetBackpackItems(){
		if item.IsEquipped {
			images = append(images, item.EquippedImage)
		}
	}

	hb := p.HealthBar()
	if hb != nil {
		images = append(images, *hb)
	}

	return images
}

func (p Player) UpdateLocation(x,y int){
	p.X = x
	p.Y = y
	db.UpdatePlayer(p.ID, p.X, p.Y)
	utils.BroadcastLocationChange(p, utils.ObjectClients)
}

func (p Player) UpdateDeltaLocation(x,y int){
	p.X += x
	p.Y += y
	db.UpdatePlayer(p.ID, p.X, p.Y)
	utils.BroadcastLocationChange(p, utils.ObjectClients)
}

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
}
