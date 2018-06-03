package leetcode

// Given n pairs of parentheses, wirte a function
// to generate all combinations of well-formed parentheses.
// For example, given n = 3, a solution set is:
// [ "((()))", "(()())", "(())()", "()(())" , "()()()" ]
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	s := "("
	return genParenthesis(s, 2*n-1, n-1, 1)
}

func genParenthesis(s string, restlen, restleft, unpairleft int) []string {
	if restlen == 1 {
		return []string{s + ")"}
	}

	if restleft > 0 && unpairleft > 0 {
		return append(genParenthesis(s+"(", restlen-1, restleft-1, unpairleft+1), genParenthesis(s+")", restlen-1, restleft, unpairleft-1)...)
	} else if unpairleft == 0 {
		return genParenthesis(s+"(", restlen-1, restleft-1, unpairleft+1)
	}
	return genParenthesis(s+")", restlen-1, restleft, unpairleft-1)
}
