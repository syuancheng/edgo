package main

import (
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
		pos2 = size/2

		return float64((nums[pos1] + nums[pos2]) / 2)
	} else {
		pos1 = size/2

		return float64(nums[pos1])
	}

	return 0
}