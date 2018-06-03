package leetcode

// Given n non-negative integers a1, a2, ..., an,
// where each represents a point at coordinate (i, ai).
// n vertical lines are dran such that the two endpoints
// of line i is at (i, ai) and (i, 0).
// Find tow lines, which together with x-axis forms a
// container, such that the container contains the most
// water.
// Note: You may not slant the container and n is at least 2.

func maxArea(height []int) int {
	max := 0
	for lo, hi := 0, len(height)-1; lo < hi; {
		x, y := hi-lo, 0
		// every time move the smaller index toward arry inner
		if height[lo] < height[hi] {
			y = height[lo]
			lo++
		} else {
			y = height[hi]
			hi--
		}
		if cur := x * y; max < cur {
			max = cur
		}
	}
	return max
}
