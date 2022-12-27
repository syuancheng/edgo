package main

func main() {
	type attr struct {
		perm int
	}
	type mymap map[string]int
	type file struct {
		attr
		mymap
		name string
	}

	f := file{
		name: "test.dat",
		attr: attr{
			perm: 94949,
		},
		mymap: map[string]int{},
	}
	f.attr.perm = 14433
	f.perm = 4567
	f.mymap["siyuan"] = 23
	println(f.mymap)
}
