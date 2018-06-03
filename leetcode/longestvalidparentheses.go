package leetcode

// Given a string containing just the characters '(' and ')',
// find the length of the longest valid(well-formed) parentheses substring.
// For "(()", the longest valid parentheses substring is "()", which has length =2.
// Another example is ")()())", where the longest valid parentheses substring is
// "()()", which has length = 4.
func longestValidParentheses(s string) int {
	stack := make([]int, len(s))[0:0]
	max, left := 0, -1
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				left = i
			} else {
				stack = stack[:len(stack)-1] // pop
				if len(stack) == 0 {
					if t := i - left; t > max {
						max = t
					}
				} else {
					if t := i - stack[len(stack)-1]; /* peek */ t > max {
						max = t
					}
				}
			}
		}
	}
	return max
}
