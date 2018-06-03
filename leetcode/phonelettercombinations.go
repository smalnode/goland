package leetcode

// Given a digitstring, return all possible letter
// combinations that the number could represent.
// The mapping of digit to letters is like normal
// phone keyboard.
// Input: Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
func letterCombinations(digits string) []string {
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}[:] // as a slice
	totalComb := 1
	for i := 0; i < len(digits); i++ {
		digit := int(digits[i] - '0')
		if digit < 0 || digit >= len(m) || m[digit] == "" {
			continue
		}
		totalComb *= len(m[digit])
	}
	res := make([]string, totalComb)[0:0]
	for i := 0; i < len(digits); i++ {
		digit := int(digits[i] - '0')
		if digit < 0 || digit >= len(m) || m[digit] == "" {
			continue
		}
		digitLetters := m[digit]
		if len(res) == 0 {
			for j := 0; j < len(digitLetters); j++ {
				res = append(res, digitLetters[j:j+1])
			}
		} else {
			curLen := len(res)
			for j := 1; j < len(digitLetters); j++ {
				for k := 0; k < curLen; k++ {
					res = append(res, res[k]+digitLetters[j:j+1])
				}
			}
			for k := 0; k < curLen; k++ {
				res[k] = res[k] + digitLetters[0:1]
			}
		}
	}
	return res
}
