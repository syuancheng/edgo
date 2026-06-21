package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Fruit struct {
	Name  string `json:"name" validate:"required"`
	Color string `json:"color,omitempty"`
}

func main() {
	fruit := Fruit{Name: "apple", Color: "red"}
	field, _ := reflect.TypeOf(fruit).FieldByName("Name")
	fmt.Println("json tag:", field.Tag.Get("json"))
	fmt.Println("validate tag:", field.Tag.Get("validate"))

	data, err := json.Marshal(fruit)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
