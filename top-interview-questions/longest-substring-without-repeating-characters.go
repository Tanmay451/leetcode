// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package topinterviewquestions

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	maxLen := 0
	chars := make(map[byte]int)
	for right < len(s) {
		if _, ok := chars[s[right]]; !ok {
			chars[s[right]] = right
			right++
			maxLen = max(maxLen, right-left)
		} else {
			delete(chars, s[left])
			left++
		}
	}
	return maxLen
}
