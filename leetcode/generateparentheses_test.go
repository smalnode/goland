package leetcode

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	for _, s := range generateParenthesis(10) {
		t.Log(s)
	}
}
