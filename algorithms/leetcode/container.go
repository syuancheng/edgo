package leetcode

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		shorter := height[left]
		if height[right] < shorter {
			shorter = height[right]
		}
		if area := width * shorter; area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
