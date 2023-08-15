package main

/*
https://leetcode.com/problems/container-with-most-water/description/
*/

import "fmt"

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	size := len(height)

	left, right := 0, size-1

	maxWidth := size - 1

	area := 0

	for width := maxWidth; width > 0; width-- {
		if height[left] < height[right] {
			area = Max(area, width*height[left])
			left++
		} else {
			area = Max(area, width*height[right])
			right--
		}
	}

	return area
}

func main() {

	list := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(list)
	fmt.Println(r)

}
