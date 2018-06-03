package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"abcdef",
			"a",
		},
		{
			"",
			"",
		},
		{
			"x",
			"x",
		},
		{
			"abccbaxefghhhgfexxx",
			"xefghhhgfex",
		},
	}
	for i, tc := range testcases {
		got := longestPalindrome(tc.s)
		if got != tc.expected {
			t.Errorf("#%d s=%s, got %s, expected %s ", i, tc.s, got, tc.expected)
		}
	}
}
