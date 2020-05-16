package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatClient struct {
	c *websocket.Conn
	ID string `json:"id"`
}

func (cw *ChatClient) readMessage() (map[string]string, error) {
	_, message, err := cw.c.ReadMessage()
	if err != nil {
		return nil, err
	}
	var result map[string]string
	json.Unmarshal([]byte(message), &result)
	fmt.Println(result)
	broadcast(result["message"], result["id"])
	return result, nil
}

func broadcast(msg, from string){
	for _, client := range clients {
		client.writeMessage(msg, from)
	}
}

func (cw *ChatClient) writeMessage(message string, from string) {
	data := map[string]string{"message":message, "from": from}
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Println("write:", err)
	}
	cw.c.WriteMessage(1, []byte(jsonString))
}

var clients = []ChatClient{}

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := ChatClient{c, "id"}
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

