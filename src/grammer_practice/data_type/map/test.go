package main

import (
	"fmt"
	"sync"
)

type user struct {
	name string
	age  byte
}

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2)

	//use ok-idiom 判断key是否存在
	if v, ok := m["d"]; ok {
		println(v) //返回零值
	}
	delete(m, "d") //即使key不存在也不会出错

	//map中的value不可以直接修改 如: m2[0].x ++. 正确做法是：
	u := m2[0]
	u.x++
	m2[0] = u
	//或者使用指针
	p := map[int]*user{
		1: &user{"krystal", 28},
	}
	p[1].age--

	//for
	m3 := map[int]string{
		1: "krystal",
		2: "jessica",
		3: "luna",
		4: "suli",
		5: "lisa",
	}
	//乱序
	for i := 0; i < 4; i++ {
		for k, v := range m3 {
			print(k, ":", v, " ")
		}
		println()
	}
	println(len(m3))
	for k, v := range m3 {
		if k == 5 {
			delete(m3, k)
			m3[6] = "amber"
		}
		print(k, ":", v, " ")

	}
	println(len(m3))

	//fatal error: concurrent map read and map write if don't use RW lock
	var lock sync.RWMutex
	m4 := make(map[string]int)
	go func() {
		for {
			lock.Lock()
			m4["a"] = 1
			lock.Unlock()
			//time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m4["b"]
			lock.RUnlock()
			//time.Sleep(time.Microsecond)
		}
	}()

	select {}

	//test make cap
}
