package stack

// 判断是否是有效的括号
func IsValidKH(s string) bool {
	kh := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		_, ok := kh[b]
		if ok {
			stack = append(stack, b)
			continue
		}
		if len(stack) < 1 || stack[len(stack)-1] != b {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

// 括号内字符串反转
func reverseParentheses(s string) string {
	stack := make([]string, 0)
	current := ""
	for _, char := range s {
		if char == '(' {
			stack = append(stack, current)
			current = ""
		} else if char == ')' {
			current = stack[len(stack)-1] + reverseString(current)
			stack = stack[:len(stack)-1]
		} else {
			current += string(char)
		}
	}
	return current
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}
