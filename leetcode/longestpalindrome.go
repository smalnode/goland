package leetcode

// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
// Example: "babad" -> "bab" / "aba"
//          "cbbd" -> "bb"
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	mlo, mhi := 0, 0
	// s[mid] or s[mid]+s[mid+1] as the symmetry point
	for mid := 0; mid < len(s)-1; mid++ {
		possibleMaxLen := mid*2 + 1
		if pm2 := (len(s)-mid-1)*2 + 1; possibleMaxLen > pm2 {
			possibleMaxLen = pm2
		}
		if maxLen := mhi - mlo + 1; possibleMaxLen < maxLen {
			break
		}

		lo, hi := mid, mid
		for ; lo > 0 && hi < len(s)-1 && s[lo-1] == s[hi+1]; lo, hi = lo-1, hi+1 {
		}
		if newLen, maxLen := hi-lo+1, mhi-mlo+1; maxLen < newLen {
			mlo, mhi = lo, hi
		}

		lo, hi = mid, mid+1
		if s[lo] == s[hi] {
			for ; lo > 0 && hi < len(s)-1 && s[lo-1] == s[hi+1]; lo, hi = lo-1, hi+1 {
			}
			if newLen, maxLen := hi-lo+1, mhi-mlo+1; maxLen < newLen {
				mlo, mhi = lo, hi
			}
		}
	}
	return s[mlo : mhi+1]
}
