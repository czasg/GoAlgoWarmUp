package others

import "math"

// 整数反转
func reverseInt(x int) int {
	ans := 0
	for x != 0 {
		if ans > math.MaxInt32/10 || ans < math.MinInt32/10 {
			return 0
		}
		pop := x % 10
		x /= 10
		ans = ans*10 + pop
	}
	return ans
}
