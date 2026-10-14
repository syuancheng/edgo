package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)

	nums := make([]int, 0, size)

	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	if size <= 0 {
		return 0
	} else if size == 1 {
		return float64(nums[0])
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var pos1, pos2 int
	if size%2 == 0 {
		pos1 = size/2 - 1
		pos2 = size / 2

		// 修复：原来是 float64((nums[pos1]+nums[pos2])/2)，整数除法会丢失 .5
		// 例如 [1,2] 的中位数是 1.5，整数除法会得到 1
		return float64(nums[pos1]+nums[pos2]) / 2
	} else {
		pos1 = size / 2

		return float64(nums[pos1])
	}

	// return 0 // 修复（vet: unreachable code）：if/else 两个分支都有 return，此行永远执行不到
}

// n4 本地验证用例（原文件只有解法函数，补一个运行入口）
func n4() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))       // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))    // 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5})) // 3
}
