package main

import (
	"errors"
	"fmt"
)

// 哨兵错误（sentinel error）：包级预定义错误变量，调用方用 errors.Is 精确判断。
// 这是工作中最常用的错误约定方式之一，命名惯例以 Err 开头。
var (
	ErrNotFound  = errors.New("record not found")
	ErrForbidden = errors.New("permission denied")
)

func findUser(id int) error {
	// 模拟：只有 id=1 存在
	if id != 1 {
		// 用 %w 包装哨兵错误：保留错误链，调用方才能用 errors.Is 判断
		return fmt.Errorf("findUser id=%d: %w", id, ErrNotFound)
	}
	return nil
}

// sentinelDemo 演示哨兵错误的标准用法。
func sentinelDemo() {
	err := findUser(42)
	if err != nil {
		// 判断错误类型要用 errors.Is，永远不要拿字符串去比较错误
		if errors.Is(err, ErrNotFound) {
			fmt.Println("业务处理：按「未找到」处理，比如返回 404")
			return
		}
		fmt.Println("其他错误:", err)
	}
}
