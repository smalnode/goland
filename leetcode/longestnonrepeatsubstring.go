package leetcode

// Given a string, find the length of the longest substring
// without repeating characters.
// Exapmles: "abcabcbb" -> "abc" 3
//           "bbbbb" -> "b" 1
//           "pwwkew" -> "wke" 3

// keep two pointer lo, hi, hi move forward,
// if s[hi] in s[lo~hi] at k, then move lo to k+1
func lengthOfLongestSubstring(s string) int {
	if l := len(s); l <= 1 {
		return l
	}
	m := make(map[byte]int, len(s))
	max := 0
	for lo, hi := 0, 0; hi < len(s); hi++ {
		if _, existed := m[s[hi]]; existed {
			if t := m[s[hi]] + 1; lo < t {
				lo = t
			}
		}
		m[s[hi]] = hi
		if cur := hi - lo + 1; max < cur {
			max = cur
		}
	}
	return max
}
