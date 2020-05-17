package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

type ObjectClient struct {
	c *websocket.Conn
	X float64 `json:"x"`
	Y float64 `json:"y"`
}


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


func (o *Object)updateLocation(x,y int){
	o.X = x
	o.Y = y
	fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	broadcastLocationChange(*o)
}

var objectClients = []ObjectClient{}

func (cw *ObjectClient) readMessage() (map[string]string, error) {
	_, message, err := cw.c.ReadMessage()
	if err != nil {
		return nil, err
	}
	var result map[string]string
	json.Unmarshal([]byte(message), &result)
	return result, nil
}

func (cw *ObjectClient) writeMessage(object Object) {
	jsonString, err := json.Marshal(object)
	if err != nil {
		log.Println("write:", err)
	}
	cw.c.WriteMessage(1, []byte(jsonString))
}

func broadcastLocationChange(object Object){
	for _, client := range objectClients {
		client.writeMessage(object)
	}
}

func objects(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := ObjectClient{c, 0,0}
	objectClients = append(objectClients, client)

	go func(client ObjectClient){
		defer client.c.Close()
		for {
			msg, err := client.readMessage()
			if err != nil {
				break
			}
			fmt.Println("objects:", msg)
		}
	}(client)
}

