package leetcode

// Given an integer array nums, find the contiguous
// subarray (containing at least one number) which
// has the largest sum and return its sum.
// Follow up:
// If you have figured out the O(n) solution, try
// coding another solution using the divide and
// conquer approach, which is more subtle.
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for lo, hi := 0, 1; lo < len(nums); {

	}
	return max
}
