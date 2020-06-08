package game

import (
	"fmt"
	"github.com/joram/game-server/db"
	"github.com/joram/game-server/monsters"
	"github.com/joram/game-server/utils"
	"log"
	"net/http"
	"sync"
)


func fullStateUpdate(client utils.ObjectClient){
	client.SendPlayerID()
	for _, otherClient := range utils.ObjectClients {
		client.UpdateMonster(otherClient.Player)
	}
	for _, monster := range allMonsters() {
		client.UpdateMonster(monster)
	}
	for _, item := range monsters.ITEMS {
		client.UpdateItem(item)
	}
}

func ServeObjects(w http.ResponseWriter, r *http.Request) {
	c, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := utils.ObjectClient{C: c, Player:monsters.Player{}, Mux:&sync.Mutex{}}

	// login
	msg, err := client.ReadMessage()
	if err != nil {
		panic(err)
	}
	accessToken, _ := msg["accessToken"]
	googleId, _ := msg["googleId"]
	email, _ := msg["email"]
	firstName, _ := msg["firstName"]
	lastName, _ := msg["lastName"]
	client.GoogleId = googleId.(string)
	client.AccessToken = accessToken.(string)
	dbPlayer := db.GetOrCreatePlayer(email.(string), firstName.(string), lastName.(string))
	client.Player = monsters.NewPlayer(dbPlayer.Id, dbPlayer.X, dbPlayer.Y)

	utils.ObjectClients = append(utils.ObjectClients, client)


	go func(client utils.ObjectClient){
		defer client.C.Close()

		fullStateUpdate(client)
		for {
			msg, err := client.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}


			// login
			if accessToken, ok := msg["accessToken"]; ok {
				googleId, _ := msg["googleId"]
				client.GoogleId = googleId.(string)
				client.AccessToken = accessToken.(string)

			// move
			} else if direction, ok := msg["direction"]; ok {
				x,y := client.Player.GetLocation()
				if direction == "left" { x -= 1}
				if direction == "right" { x += 1}
				if direction == "up" { y -= 1}
				if direction == "down" { y += 1}
				if !client.Player.IsDead() {
					m := monsterAt(x, y)
					if m != nil && !m.IsDead(){
						client.Player.Attack(m)
						//fmt.Printf("player[%d] attacks %s[%d]\n", client.Player.GetID(), m.GetType(), m.GetID())
						//m.TakeDamage(5, client.Player)
					} else {
						if !utils.IsSolid(x,y) {
							client.Player.UpdateLocation(x, y)
						}
					}
				}


			// list everything
			} else if _, ok := msg["full_state"]; ok {
				fullStateUpdate(client)

			// list items in backpack
			} else if _, ok := msg["backpack"]; ok {
				for _, item := range client.Player.GetBackpackItems() {
					client.SendBackpackItem(item)
				}
				for _, item := range monsters.ITEMS {
					x,y := client.Player.GetLocation()
					if item.X == x && item.Y == y {
						client.SendBackpackItem(item)
						fmt.Printf("%s[%d] on the ground\n", item.Name, item.ID)
					}
				}

			// equip item
			} else if id, ok := msg["equip_item"]; ok {
				item :=  client.Player.EquipItem(int(id.(float64)))
				client.SendBackpackItem(item)

			// unequip item
			} else if id, ok := msg["unequip_item"]; ok {
				item :=  client.Player.UnequipItem(int(id.(float64)))
				client.SendBackpackItem(item)

			// drop item
			} else if id, ok := msg["drop_item"]; ok {
				item :=  client.Player.DropItem(int(id.(float64)))
				client.SendBackpackItem(item)
			}
	}

		newOjectClients := []utils.ObjectClient{}
		for _, otherCLient := range utils.ObjectClients {
			otherCLient.RemoveObject(client.Player)
			if otherCLient.Player.GetID() != client.Player.GetID() {
				newOjectClients = append(newOjectClients, otherCLient)
			}
		}
		utils.ObjectClients = newOjectClients

		player := client.Player.(monsters.Player)
		player.Unregister()

	}(client)
}


