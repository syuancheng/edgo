package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	roundedDown := now.Truncate(10 * time.Minute)

	fmt.Println("before:", now.Format(time.RFC3339Nano))
	fmt.Println("after: ", roundedDown.Format(time.RFC3339Nano))
}
