package main

import (
	"fmt"
	"sync"
)

func main() {
	const workers = 4
	results := make([]int, workers)

	var wg sync.WaitGroup
	for id := 0; id < workers; id++ {
		id := id // 在所有受支持的 Go 版本中都为当前迭代创建独立值。
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[id] = id * id
		}()
	}

	wg.Wait()
	fmt.Println(results)
}
