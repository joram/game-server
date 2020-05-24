package game

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joram/game-server/utils"
	"log"
	"net/http"
	"sync"
)



type ChatClient struct {
	c *websocket.Conn
	ID string `json:"id"`
	Mux    *sync.Mutex
}

func (cw *ChatClient) readMessage() (map[string]string, error) {
	cw.Mux.Lock()
	defer cw.Mux.Unlock()

	_, message, err := cw.c.ReadMessage()
	if err != nil {
		return nil, err
	}
	var result map[string]string
	json.Unmarshal([]byte(message), &result)
	broadcastChatMessage(result["message"], result["id"])
	return result, nil
}

func broadcastChatMessage(msg, from string){
	for _, client := range clients {
		client.writeMessage(msg, from)
	}
}

func (cw *ChatClient) writeMessage(message string, from string) {
	cw.Mux.Lock()
	defer cw.Mux.Unlock()

	data := map[string]string{"message":message, "from": from}
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Println("write:", err)
	}
	cw.c.WriteMessage(1, []byte(jsonString))
}

var clients = []ChatClient{}

func ServeChat(w http.ResponseWriter, r *http.Request) {
	c, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := ChatClient{c:c, ID:"id",Mux: &sync.Mutex{}}
	clients = append(clients, client)

	go func(client ChatClient){
		defer client.c.Close()
		for {
			msg, err := client.readMessage()
			if err != nil {
				break
			}
			fmt.Println(msg)
		}
	}(client)
}

