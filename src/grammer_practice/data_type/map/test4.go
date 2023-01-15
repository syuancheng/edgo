package main

func main() {

	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}

	//map中的value不可以直接修改 如: m2[0].x ++. 正确做法是：
	u := m2[0]
	u.x++
	m2[0] = u
	//或者使用指针
	p := map[int]*user{
		1: &user{"krystal", 28},
	}
	p[1].age--
}
