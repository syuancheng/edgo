package main

func ex_5_2_1() {
	b := [...]int{1, 2}
	var cb [2]int

	// 修复编译：原文件 b/cb 声明后未使用（会编译失败），补充打印
	println("len(b) =", len(b), "len(cb) =", len(cb))
}
