package main

import (
	"fmt"
	"time"
)

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

// Add=1, Subtract=2, Multiply=3

// recommend
func isActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}

// unrecommend
func isActive1(now, start, stop int) bool {
	return start <= now && now < stop
}

func main() {
	now := time.Now()
	start := time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2022, 12, 31, 23, 0, 0, 0, time.UTC)
	fmt.Println(isActive(now, start, end))
}
