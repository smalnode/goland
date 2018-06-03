package leetcode

// Implement next permutation, which rearranges numbers
// into the lexicographically next greater permutation
// of numbers. If such arrangement is not possible, it
// must rearrange it as the lowest possible order (ie,
// sorted in ascending order).
// The replacement must be in-place and use only constant
// extra memory.
// Examples: 1, 2, 3 -> 1, 3, 2
//           3, 2, 1 -> 1, 2, 3
//           1, 1, 5 -> 1, 5, 1
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	var i, j int
outloop:
	for i = len(nums) - 2; i >= 0; i-- {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break outloop
			}
		}
	}

	var lo, hi int
	if i >= 0 {
		nums[i], nums[j] = nums[j], nums[i]
		lo, hi = i+1, len(nums)-1
	} else {
		lo, hi = 0, len(nums)-1
	}

	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}

}
