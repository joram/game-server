package game

import (
	"fmt"
	"github.com/joram/game-server/monsters"
	"github.com/joram/game-server/utils"
	"log"
	"net/http"
	"sync"
)

func ServeObjects(w http.ResponseWriter, r *http.Request) {
	c, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	player := monsters.NewPlayer(0,0)
	client := utils.ObjectClient{C:c, Player:player, Mux:&sync.Mutex{}}
	utils.ObjectClients = append(utils.ObjectClients, client)

	go func(client utils.ObjectClient){
		defer client.C.Close()

		client.SendPlayerID()
		for _, otherClient := range utils.ObjectClients {
			client.UpdateObject(otherClient.Player)
		}
		for _, o := range allMonsters() {
			client.UpdateObject(o)
		}

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

			} else if direction, ok := msg["direction"]; ok {
				x,y := client.Player.GetLocation()
				if direction == "left" { x -= 1}
				if direction == "right" { x += 1}
				if direction == "up" { y -= 1}
				if direction == "down" { y += 1}
				if !client.Player.IsDead() {
					m := monsterAt(x, y)
					if m != nil && !m.IsDead(){
						fmt.Printf("player[%d] attacks %s[%d]\n", client.Player.GetID(), m.GetType(), m.GetID())
						m.TakeDamage(5, client.Player)
					} else {
						if !utils.IsSolid(x,y) {
							client.Player.UpdateLocation(x, y)
						}
					}
				}

			// position update
			} else {
				x := int(msg["x"].(float64))
				y := int(msg["y"].(float64))
				if !client.Player.IsDead() {
					m := monsterAt(x, y)
					if m != nil && !m.IsDead(){
						fmt.Printf("player[%d] attacks %s[%d]\n", client.Player.GetID(), m.GetType(), m.GetID())
						m.TakeDamage(5, client.Player)
					} else {
						client.Player.UpdateLocation(x, y)
					}
				}
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


