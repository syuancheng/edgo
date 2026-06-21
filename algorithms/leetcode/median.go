package leetcode

// MedianSortedArrays returns the median and false when both inputs are empty.
// It performs a binary search over the shorter input in O(log(min(m, n))).
func MedianSortedArrays(left, right []int) (float64, bool) {
	if len(left)+len(right) == 0 {
		return 0, false
	}
	if len(left) > len(right) {
		left, right = right, left
	}

	low, high := 0, len(left)
	leftSize := (len(left) + len(right) + 1) / 2
	for low <= high {
		leftCut := (low + high) / 2
		rightCut := leftSize - leftCut

		if leftCut < len(left) && rightCut > 0 && right[rightCut-1] > left[leftCut] {
			low = leftCut + 1
			continue
		}
		if leftCut > 0 && rightCut < len(right) && left[leftCut-1] > right[rightCut] {
			high = leftCut - 1
			continue
		}

		var maxLeft int
		switch {
		case leftCut == 0:
			maxLeft = right[rightCut-1]
		case rightCut == 0:
			maxLeft = left[leftCut-1]
		case left[leftCut-1] > right[rightCut-1]:
			maxLeft = left[leftCut-1]
		default:
			maxLeft = right[rightCut-1]
		}

		if (len(left)+len(right))%2 == 1 {
			return float64(maxLeft), true
		}

		var minRight int
		switch {
		case leftCut == len(left):
			minRight = right[rightCut]
		case rightCut == len(right):
			minRight = left[leftCut]
		case left[leftCut] < right[rightCut]:
			minRight = left[leftCut]
		default:
			minRight = right[rightCut]
		}
		return (float64(maxLeft) + float64(minRight)) / 2, true
	}

	return 0, false
}
