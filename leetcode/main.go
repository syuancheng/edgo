package main

import (
	"fmt"
	"os"
	"sort"
)

// demos 注册本目录所有题目。
// 用法: go run . <题号>   例如 go run . n2；不带参数则列出全部。
var demos = map[string]func(){
	"n2":  n2,
	"n4":  n4,
	"n11": n11,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("可用题目（go run . <题号>）:")
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
		fmt.Printf("未找到题目 %q，不带参数运行可查看全部题目\n", os.Args[1])
		os.Exit(1)
	}
	demo()
}
