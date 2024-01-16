package others

import (
	"sort"
)

// 无序数组求交集
func setArr(a, b []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	ans := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			ans = append(ans, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}
