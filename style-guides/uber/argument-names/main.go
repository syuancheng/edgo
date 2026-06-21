package main

import "fmt"

func printInfo(name string, isLocal, done bool) {
	fmt.Println(name, isLocal, done)
}

func main() {

	printInfo("foo", true /* isLocal */, true /* done */)

	var s1 []int
	fmt.Println(len(s1))

}
