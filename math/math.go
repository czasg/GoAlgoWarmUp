package math

// 最大值
func Max(a, b int, others ...int) (ans int) {
	if a > b {
		ans = a
	} else {
		ans = b
	}
	for i := 0; i < len(others); i++ {
		if others[i] > ans {
			ans = others[i]
		}
	}
	return
}

// 最小值
func Min(a, b int, others ...int) (ans int) {
	if a < b {
		ans = a
	} else {
		ans = b
	}
	for i := 0; i < len(others); i++ {
		if others[i] < ans {
			ans = others[i]
		}
	}
	return
}

// 绝对值
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
