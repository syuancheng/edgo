package main

func main() {
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

	p1 := &p
	//*p1.name = "zhao" //不支持
}
