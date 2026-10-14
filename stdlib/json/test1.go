package main

import (
	"encoding/json"
	"fmt"
)

type MbpResponse struct {
	Version string `json:"version"`
}

func main() {
	data := []byte{}

	mbpResp := &MbpResponse{}
	fmt.Println(string(data))
	if err := json.Unmarshal(data, mbpResp); err != nil {
		fmt.Printf("err=%+v\n", err)
	}
}
