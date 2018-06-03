package leetcode

import "testing"
import "fmt"

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 1, 5},
			[]int{1, 5, 1},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{6, 9, 8, 7, 7, 5, 5, 4, 3, 2, 1},
			[]int{7, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		},
	}
	for i, tc := range testCases {
		s := fmt.Sprintf("%v", tc.nums)
		nextPermutation(tc.nums)
		for k, v := range tc.nums {
			if v != tc.expected[k] {
				t.Errorf("#%d nums=%s, expected %v, got %v ", i, s, tc.expected, tc.nums)
				break
			}
		}
	}
}
