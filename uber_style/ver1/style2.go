package main

import "fmt"

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	sVals := map[int]S{1: S{"ddd"}}

	sVals[1].Read()
	// sVals[1].Write("eee")//

	pVals := map[int]*S{1: &S{"ppp"}}

	pVals[1].Read()
	pVals[1].Write("ptptp")

	t := S{"zhao"}
	fmt.Println(t)
	t.Read()
	t.Write("test")//update success
	fmt.Println(t)

}
