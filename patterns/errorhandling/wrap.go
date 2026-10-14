package main

import (
	"errors"
	"fmt"
)

// ValidationError 自定义错误类型：实现 error 接口的结构体，
// 可以携带结构化的上下文信息，调用方用 errors.As 提取。
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: field=%s reason=%s", e.Field, e.Reason)
}

func validateAge(age int) error {
	if age < 0 {
		// 用 %w 包装自定义错误类型
		return fmt.Errorf("validateAge: %w", &ValidationError{Field: "age", Reason: "不能为负数"})
	}
	return nil
}

// wrapDemo 演示错误包装（%w）、errors.As 类型断言、以及 %w 与 %v 的区别。
func wrapDemo() {
	err := validateAge(-1)
	fmt.Println("err:", err)

	// errors.As：沿着错误链查找指定类型的错误，找到后提取出来
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("As 提取到类型化错误: field=%s, reason=%s\n", ve.Field, ve.Reason)
	}

	// %w（Go 1.13+）：包装并保留错误链，Is/As 可以穿透
	wrapped := fmt.Errorf("handler: %w", err)
	fmt.Println("wrapped:", wrapped)
	fmt.Println("errors.Is 沿链查找:", errors.Is(wrapped, err))

	// 对比 %v：只拼接错误文本，链条丢失，Is/As 从此查不到 —— 工作中常见的坑
	broken := fmt.Errorf("handler: %v", err)
	fmt.Println("broken(用 v 动词包装):", broken)
	var ve2 *ValidationError
	fmt.Println("errors.As 对 v 包装结果失效:", !errors.As(broken, &ve2))
}
