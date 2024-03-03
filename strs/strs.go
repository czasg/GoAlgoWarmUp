package strs

import (
	"math/rand"
	"proj/math"
	"time"
)

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

// 生成随机字符串
func genRandomStr(length int, need_special_char bool) string {
	baseChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialChars := "!@#$%^&*~"
	if need_special_char {
		baseChars += specialChars
	}
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = baseChars[rand.Intn(len(baseChars))]
	}

	return string(result)
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		if i > len(strs[0]) {
			return strs[0]
		}
		pub := strs[0][:i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j]) {
				return strs[j]
			}
			if pub != strs[j][:i] {
				if i == 0 {
					return ""
				}
				return strs[0][:i-1]
			}
		}
	}
}
