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

		for _, otherClient := range utils.ObjectClients {
			client.UpdateObject(otherClient.Player)
		}
		for _, o := range allMonsters() {
			client.UpdateObject(o)
		}

		for {
			msg, err := client.ReadMessage()
			if err != nil {
				break
			}

			// login
			if accessToken, ok := msg["accessToken"]; ok {
				googleId, _ := msg["googleId"]
				client.GoogleId = googleId.(string)
				client.AccessToken = accessToken.(string)

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
			newOjectClients = append(newOjectClients, otherCLient)
		}
		utils.ObjectClients = newOjectClients

		player := client.Player.(monsters.Player)
		player.Unregister()

	}(client)
}


