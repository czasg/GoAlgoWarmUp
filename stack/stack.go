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
