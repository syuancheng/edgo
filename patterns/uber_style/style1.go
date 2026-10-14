package main

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func style1() {

	// f1.f() 无法修改底层数据
	// f2.f() 可以修改底层数据，给接口变量 f2 赋值时使用的是对象指针
	var f1 F = S1{}
	var f2 F = &S2{}

	// 修复编译：原文件 f1/f2 声明后未使用，实际调用一下
	f1.f() // 值接收器：无法修改底层数据
	f2.f() // 指针接收器：可以修改底层数据

	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	i.f() // 修复编译：i 声明后未使用

	_ = s2Val // 修复编译：s2Val 只在下方注释示例中被提及
	//  下面代码无法通过编译。因为 s2Val 是一个值，而 S2 的 f 方法中没有使用值接收器
	// i = s2Val
}
