package main

import (
	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(2, 5)
}
