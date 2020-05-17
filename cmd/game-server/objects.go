package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

type Object struct {
	ID int `json:"id"`
	X int `json:"x"`
	Y int `json:"y"`
	Type string `json:"type"`
	Image string `json:"image"`
	Solid bool `json:"solid"`
}

type ObjectType struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Solid bool `json:"solid"`
}

func LoadObjectTypes() []ObjectType {
	var objectTypes []ObjectType
	file, _ := ioutil.ReadFile("object_types.json")
	_ = json.Unmarshal([]byte(file), &objectTypes)
	return objectTypes
}

func useTypeData(object *Object, objectTypes []ObjectType){
	for _, ot := range objectTypes {
		if ot.Name == object.Type {
			object.Image = ot.Image
			object.Solid = ot.Solid
		}
	}
}

func LoadObjects() []Object {
	objectTypes := LoadObjectTypes()

	var objects []Object
	file, _ := ioutil.ReadFile("objects.json")
	_ = json.Unmarshal([]byte(file), &objects)

	i := 0
	for i, _ = range objects {
		objects[i].ID = i
		useTypeData(&objects[i], objectTypes)
	}

	goblin := Object{
		ID: i+1,
		X:     -5,
		Y:     -5,
		Type:  "goblin",
	}
	useTypeData(&goblin, objectTypes)
	objects = append(objects, goblin)
	go func(goblin Object){
		for {
			time.Sleep(time.Second)
			goblin.updateLocation(-5,-5)
			time.Sleep(time.Second)
			goblin.updateLocation(-5,-4)
		}
	}(goblin)

	return objects
}

func (o *Object) updateLocation(x,y int){
	o.X = x
	o.Y = y
	fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	broadcastLocationChange(*o)
}


type ObjectClient struct {
	c *websocket.Conn
	Character *Object
}

var objectClients = []ObjectClient{}

func (cw *ObjectClient) readMessage() (map[string]interface{}, error) {
	_, message, err := cw.c.ReadMessage()
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(message), &result)
	return result, nil
}

func (cw *ObjectClient) removeObject(object Object) {
	type removeMessage struct {
		Action string `json:"action"`
		ObjectID int `json:"id"`
	}

	rm := removeMessage{"remove", object.ID}
	jsonString, err := json.Marshal(rm)
	if err != nil {
		log.Println("write:", err)
	}
	cw.c.WriteMessage(1, []byte(jsonString))
}

func (cw *ObjectClient) updateObject(object Object) {
	jsonString, err := json.Marshal(object)
	if err != nil {
		log.Println("write:", err)
	}
	cw.c.WriteMessage(1, []byte(jsonString))
}

func broadcastLocationChange(object Object){
	for _, client := range objectClients {
		client.updateObject(object)
	}
}

func objects(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	character := Object{rand.Int(),0,0,"character", "http://localhost:2303/images/character.png", true}
	client := ObjectClient{c, &character}
	objectClients = append(objectClients, client)

	go func(client ObjectClient){
		defer client.c.Close()

		fmt.Printf("telling %d about %d other clients \n", client.Character.ID, len(objectClients)-1)
		for _, otherClient := range objectClients {
			if otherClient == client { continue }
			fmt.Printf("telling %d about %d at (%d,%d)\n", client.Character.ID, otherClient.Character.ID, otherClient.Character.X, otherClient.Character.Y)
			client.updateObject(*otherClient.Character)
		}

		for {
			msg, err := client.readMessage()
			if err != nil {
				break
			}
			//fmt.Println("objects:", msg)
			x := int(msg["x"].(float64))
			y := int(msg["y"].(float64))
			client.Character.updateLocation(x, y)
		}

		newOjectClients := []ObjectClient{}
		for _, otherCLient := range objectClients {
			if otherCLient == client { continue }
			otherCLient.removeObject(*client.Character)
			newOjectClients = append(newOjectClients, otherCLient)
		}
		objectClients = newOjectClients

	}(client)
}

