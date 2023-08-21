package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []int{1, 2, 3}
	index := 10 // This index is out of range for the 'data' slice

	// Attempt to access an element using an out-of-range index
	value := data[index]

	fmt.Println("Value:", value)

	// Attempt to marshal the data to JSON
	_, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
	}
}