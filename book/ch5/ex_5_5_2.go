package main

func ex_5_5_2() {
	type user struct {
		name string
		age  int
	}

	p := &user{
		name: "zhao",
		age:  12,
	}

	p.age++
	p.name += " kuangyi"

	p1 := &p // p1 是 **user（二级指针）

	// 原写法：*p1.name = "zhao" //编译错误：多级指针不能直接访问字段
	// 这正是本例要演示的点。正确写法是先解引用回一级指针：
	(*p1).name = "zhao"
}
