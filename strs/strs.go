package strs

import "proj/math"

// 最长回文子串-中心扩展法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := math.Max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// 无重复字符的最长子串-滑动窗口
func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	start := 0
	maxLength := 0
	for end, char := range s {
		// 如果字符已经在窗口中，更新窗口的起始位置
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		// 更新字符在窗口中的位置
		charIndexMap[char] = end
		// 计算当前子串的长度
		currentLength := end - start + 1
		// 更新最大长度
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
