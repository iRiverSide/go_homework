package goTaskOne

func mySqrt(x int) int {
	left, right, result := 0, x, -1
	for left <= right {
		middle := (left + right) / 2
		if middle*middle <= x {
			result = middle
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return result
}
