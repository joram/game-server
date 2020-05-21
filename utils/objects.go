package utils

import (
	"encoding/json"
	"fmt"
	//"github.com/joram/game-server/game"
	"io/ioutil"
)

type Object struct {
	ID int `json:"id"`
	X int `json:"x"`
	Y int `json:"y"`
	Type string `json:"type"`
	Images []string `json:"images"`
	Solid bool `json:"solid"`
}

type ObjectInterface interface {
	UpdateLocation(x,y int)
	UpdateDeltaLocation(x,y int)
	GetLocation() (x,y int)
	GetID() int
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
	fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	BroadcastLocationChange(o, ObjectClients)
}

func (o *Object) UpdateDeltaLocation(x,y int){
	o.X += x
	o.Y += y
	fmt.Printf("moving %d to (%d, %d)\n", o.ID, o.X, o.Y)
	BroadcastLocationChange(o, ObjectClients)
}


func (o *Object) GetLocation() (x,y int) {
	return o.X, o.Y
}

func (o *Object) GetID() int {
	return o.ID
}