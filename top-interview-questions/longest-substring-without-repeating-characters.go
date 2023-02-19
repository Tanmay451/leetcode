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

// This is a Go programming language implementation of a function called lengthOfLongestSubstring. This function takes a string s as input and returns an integer that represents the length of the longest substring without repeating characters in the input string.

// Here's how the function works:

// The function initializes three variables:
// left represents the left boundary of the current substring.
// right represents the right boundary of the current substring.
// maxLen represents the maximum length of a substring without repeating characters found so far.
// The function initializes a chars map to keep track of the characters in the current substring and their last seen positions.
// The function enters a loop that continues until the right boundary right reaches the end of the input string s.
// Inside the loop, the function checks if the character s[right] is already in the chars map by using the _ , ok := chars[s[right]] assignment.
// If the character s[right] is not in the chars map, it is added along with its position right. The right index is incremented, and the current length of the substring right-left is compared with the maximum length maxLen so far. If the current length is greater than maxLen, maxLen is updated.
// If the character s[right] is already in the chars map, the character at the left boundary s[left] is deleted from the map, and the left index is incremented. The loop continues from step 4.
// Once the loop completes, the function returns the maximum length of a substring without repeating characters found so far.
// The max function is a helper function used to calculate the maximum of two integers. It takes two integers a and b as input and returns the greater of the two integers. It is used in step 5 to update the maxLen variable.
