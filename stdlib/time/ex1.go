package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Printf("before t: %s, hour: %d, min: %d, sec: %d\n", t, t.Hour(), t.Minute(), t.Second())

	diff := t.Minute() % 10
	if diff != 0 {
		t = t.Add(time.Duration(-diff)*time.Minute + time.Duration(-t.Second())*time.Second)

	}
	fmt.Printf("after t: %s, hour: %d, min: %d, sec: %d\n", t, t.Hour(), t.Minute(), t.Second())

}
