package leetcode

// Given an array of integers, return indices of two numbers
// such that they add up to a specific target. You may assume
// that each input would have exactly one solution, and you
// may not use the same element twice

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if ii, ok := m[target-v]; ok {
			return []int{ii, i}
		}
		m[v] = i
	}
	return nil
}
