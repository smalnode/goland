package leetcode

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	testcases := []struct {
		s string
		l int
	}{
		{
			"(()",
			2,
		},
		{
			")()())",
			4,
		},
		{
			")()())((())()))))",
			8,
		},
		{
			"()(()",
			2,
		},
		{
			"(()(((()",
			2,
		},
		{
			"(())",
			4,
		},
	}

	for i, c := range testcases {
		got := longestValidParentheses(c.s)
		if got != c.l {
			t.Errorf("#%d s=%s, expected %d, got %d ", i, c.s, c.l, got)
		}
	}
}
