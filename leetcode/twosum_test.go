package leetcode

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		solution []int
	}{
		{
			[]int{1, 3, 5, 4, 8, 0, 5},
			10,
			[]int{2, 6},
		},
		{
			[]int{1, 3, 5, 7, 9},
			-1,
			nil,
		},
	}

	for i, testCase := range testCases {
		s := twoSum(testCase.nums, testCase.target)
		if s == nil && testCase.solution != nil {
			t.Errorf("#%d solution is nil, expected %v ", i, testCase.solution)
		} else if s != nil && testCase.solution == nil {
			t.Errorf("#%d solution is %v, expected nil ", i, s)
		} else if s != nil && testCase.solution != nil {
			if len(testCase.solution) != 2 {
				t.Errorf("#%d test case solution illegal %v ", i, testCase.solution)
			}
			if len(s) != 2 {
				t.Errorf("#%d len(s) is %d, expected 2 ", i, len(s))
			}
			sort.Ints(s)
			for k, v := range s {
				if v != testCase.solution[k] {
					t.Errorf("#%d got %v, expected %v ", i, s, testCase.solution)
				}
			}
		}
	}
}
