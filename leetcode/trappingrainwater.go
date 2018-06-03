package leetcode

// Given n non-negative integers representing an elevation map
// where the width of each bar is 1, compute how much water it
// is able to trap after raining.
// For example, given [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1],
// return 6.

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	w := 0
	maxLevel := 0
	lo, hi := 0, len(height)-1
	var curLevel int
	for lo < hi {
		x := hi - lo - 1
		// move the smaller height index toward array inner
		if height[lo] < height[hi] {
			curLevel = height[lo]
			lo++
		} else {
			curLevel = height[hi]
			hi--
		}
		if curLevel > maxLevel {
			// water level up, add level delta and substract current bar occupation
			w += (curLevel-maxLevel)*x - maxLevel*1
			maxLevel = curLevel
		} else {
			// water level down, it must contian in max level range,
			// just substract current bar occupation
			w -= curLevel * 1
		}
	}
	return w
}
