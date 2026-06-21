package main

import (
	"fmt"
	"sort"
	"sync"
)

type user struct {
	name string
	age  int
}

func main() {
	var nilMap map[string]int
	fmt.Println("nil map read:", nilMap["missing"])
	// nilMap["answer"] = 42 // panic: assignment to entry in nil map

	scores := map[string]int{"Go": 90, "C": 80}
	value, ok := scores["Rust"]
	fmt.Printf("lookup: value=%d exists=%t\n", value, ok)

	users := map[int]user{1: {name: "Gopher", age: 10}}
	u := users[1] // map 元素不可寻址，先取出、修改，再写回。
	u.age++
	users[1] = u

	keys := make([]string, 0, len(scores))
	for key := range scores {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 需要稳定输出时显式排序。
	for _, key := range keys {
		fmt.Printf("%s=%d ", key, scores[key])
	}
	fmt.Println()

	var mu sync.RWMutex
	mu.Lock()
	scores["Go"]++
	mu.Unlock()
	mu.RLock()
	fmt.Println("Go:", scores["Go"])
	mu.RUnlock()
}
