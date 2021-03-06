package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Object struct {
	ID int `json:"id"`
	X int `json:"x"`
	Y int `json:"y"`
	Type string `json:"type"`
	Images []string `json:"images"`
	Solid bool `json:"solid"`
}

func (o *Object) AsString() string {
	if o.Type == "player" {
		panic("damnit")
	}
	jsonString, err := json.Marshal(o)
	if err != nil {
		log.Println("write:", err)
	}
	return string(jsonString)
}

type ObjectInterface interface {
	UpdateLocation(x,y int)
	UpdateDeltaLocation(x,y int)
	GetLocation() (x,y int)
	GetID() int
	AsString() string
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

func (o *Object) UpdateLocation(x,y int){
	o.X = x
	o.Y = y
	fmt.Printf("moving %s[%d] to (%d, %d)\n", o.Type, o.ID, o.X, o.Y)

	BroadcastLocationChange(o, ObjectClients)
}

func (o *Object) UpdateDeltaLocation(x,y int){
	o.X += x
	o.Y += y
	//fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	BroadcastLocationChange(o, ObjectClients)
}


func (o *Object) GetLocation() (x,y int) {
	return o.X, o.Y
}

func (o *Object) GetID() int {
	return o.ID
}

func (o *Object) Broadcast(){
	for _, client := range ObjectClients {
		client.UpdateObject(o)
	}
}

func (o *Object) BroadcastRemove(){
	for _, client := range ObjectClients {
		client.RemoveObject(o)
	}
}
