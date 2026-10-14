package main

import (
	"fmt"
	"os"
	"sort"
)

// demos 注册本目录所有示例。
// 用法: go run . <示例名>   运行指定示例；不带参数则列出全部可用示例。
var demos = map[string]func(){
	"test1": test1,
	"test2": test2,
	"test3": test3,
	"test4": test4,
	"test5": test5,
	"test6": test6,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("可用示例（go run . <名称>）:")
		names := make([]string, 0, len(demos))
		for name := range demos {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Println("  -", name)
		}
		return
	}
	demo, ok := demos[os.Args[1]]
	if !ok {
		fmt.Printf("未找到示例 %q，不带参数运行可查看全部示例\n", os.Args[1])
		os.Exit(1)
	}
	demo()
}
