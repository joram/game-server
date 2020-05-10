package main

import (
	"encoding/json"
	"io/ioutil"
)

type Object struct {
	ID int `json:"id"`
	X int `json:"x"`
	Y int `json:"y"`
	Type string `json:"type"`
	Image string `json:"image"`
	Solid string `json:"solid"`
}

type ObjectType struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Solid string `json:"solid"`
}

func LoadObjectTypes() []ObjectType {
	var objectTypes []ObjectType
	file, _ := ioutil.ReadFile("object_types.json")
	_ = json.Unmarshal([]byte(file), &objectTypes)
	return objectTypes
}

func LoadObjects() []Object {
	objectTypes := LoadObjectTypes()

	var objects []Object
	file, _ := ioutil.ReadFile("objects.json")
	_ = json.Unmarshal([]byte(file), &objects)

	for i, o := range objects {
		for _, ot := range objectTypes {
			if ot.Name == o.Type {
				objects[i].Image = ot.Image
				objects[i].ID = i
				objects[i].Solid = ot.Solid
			}
		}
	}
	return objects
}
