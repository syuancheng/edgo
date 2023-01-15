package main

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func main() {

	// f1.f() 无法修改底层数据
	// f2.f() 可以修改底层数据，给接口变量 f2 赋值时使用的是对象指针
	var f1 F = S1{}
	var f2 F = &S2{}

	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	//  下面代码无法通过编译。因为 s2Val 是一个值，而 S2 的 f 方法中没有使用值接收器
  	// i = s2Val
}
