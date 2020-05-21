package game

import (
	"github.com/joram/game-server/monsters"
	"github.com/joram/game-server/utils"
	"log"
	"net/http"
)

func ServeObjects(w http.ResponseWriter, r *http.Request) {
	c, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	character := monsters.NewPlayer(0,0)
	client := utils.ObjectClient{c, character}
	utils.ObjectClients = append(utils.ObjectClients, client)

	go func(client utils.ObjectClient){
		defer client.C.Close()

		for _, otherClient := range utils.ObjectClients {
			client.UpdateObject(otherClient.Character)
		}
		for _, o := range MONSTERS {
			client.UpdateObject(o)
		}

		for {
			msg, err := client.ReadMessage()
			if err != nil {
				break
			}
			x := int(msg["x"].(float64))
			y := int(msg["y"].(float64))
			client.Character.UpdateLocation(x, y)

		}

		newOjectClients := []utils.ObjectClient{}
		for _, otherCLient := range utils.ObjectClients {
			otherCLient.RemoveObject(client.Character)
			newOjectClients = append(newOjectClients, otherCLient)
		}
		utils.ObjectClients = newOjectClients

		player := client.Character.(monsters.Player)
		player.Unregister()

	}(client)
}


