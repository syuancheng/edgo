package main

import "fmt"

type FashionItemManager struct {
	itemList []int64
}

var manger = &FashionItemManager{}

func main() {
	l := len(manger.itemList)
	fmt.Printf("list length is %s", l)
}
