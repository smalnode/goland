package leetcode

import "testing"

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		digits string
		comb   []string
	}{
		{
			"0",
			[]string{},
		},
		{
			"1",
			[]string{},
		},
		{
			"01100*##*",
			[]string{},
		},
		{
			"23",
			[]string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
		},
		{
			"2003",
			[]string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
		},
		{
			"#*200300",
			[]string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
		},
		{
			"97",
			[]string{"wp", "xp", "yp", "zp", "wq", "xq", "yq", "zq", "wr", "xr", "yr", "zr", "ws", "xs", "ys", "zs"},
		},
		{
			"22",
			[]string{"aa", "ba", "ca", "ab", "bb", "cb", "ac", "bc", "cc"},
		},
	}
	for i, testCase := range testCases {
		got := letterCombinations(testCase.digits)
		if len(got) != len(testCase.comb) {
			t.Errorf("#%d digits=%s, got %s, expected %s ", i, testCase.digits, got, testCase.comb)
		}
		for j := 0; j < len(got); j++ {
			if got[j] != testCase.comb[j] {
				t.Errorf("#%d digits=%s, got %s, expected %s, mismatch at pos %d ", i, testCase.digits, got, testCase.comb, j)
			}
		}
	}
}

func TestStringSlice(t *testing.T) {
	s := "abcd"
	t.Log(s[0:1])
}
