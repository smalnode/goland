package leetcode

import "testing"

func TestLengthOfLonegestSubstring(t *testing.T) {
	testcases := []struct {
		s string
		l int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"",
			0,
		},
		{
			"b",
			1,
		},
	}
	for i, c := range testcases {
		got := lengthOfLongestSubstring(c.s)
		if got != c.l {
			t.Errorf("#%d s=%s got %d, expected %d ", i, c.s, got, c.l)
		}
	}
}
