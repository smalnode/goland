package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		solution [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4, 5},
			[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
				{-1, -4, 5},
			},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[][]int{},
		},
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{-1, -2, -3, -4, -5},
			[][]int{},
		},
		{
			[]int{-1, -2, 3, -4, -5, 9, 5},
			[][]int{
				{-2, -1, 3},
				{-5, -4, 9},
				{-4, -1, 5},
			},
		},
		{
			[]int{-5, -4, -2, -1, -1, -1, 3, 5, 9, 9, 9},
			[][]int{
				{-2, -1, 3},
				{-5, -4, 9},
				{-4, -1, 5},
			},
		},
	}

	for i, testCase := range testCases {
		s := threeSum2(testCase.nums)
		if len(s) != len(testCase.solution) {
			t.Errorf("#%d len(s) got %d, expected %d ", i, len(s), len(testCase.solution))
		}
		for k, e := range s {
			if len(e) != 3 {
				t.Errorf("#%d.%d solution len got %d, expected 3 ", i, k, len(e))
			}
			if e[0]+e[1]+e[2] != 0 {
				t.Errorf("#%d.%d solution %v illegal ", i, k, e)
			}
		}
	}
}
