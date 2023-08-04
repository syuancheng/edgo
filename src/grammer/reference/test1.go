package main


type Person struct {
	name string
	school string
}

func main() {
	var p *Person
	p.name = "tom"
	p.school = "UU"
}