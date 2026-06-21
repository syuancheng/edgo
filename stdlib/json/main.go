package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Response struct {
	Version string `json:"version"`
	Ready   bool   `json:"ready"`
}

func main() {
	decoder := json.NewDecoder(strings.NewReader(`{"version":"1.0","ready":true}`))
	decoder.DisallowUnknownFields()

	var response Response
	if err := decoder.Decode(&response); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)

	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
