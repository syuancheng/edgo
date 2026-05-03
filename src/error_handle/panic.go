package main

import "log"

func main() {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				log.Printf("recover: %v", e)
			}
		}()
		panic("煎鱼焦了")
	}()

	log.Println("Go编程之旅：一起用Go做项目")
}
