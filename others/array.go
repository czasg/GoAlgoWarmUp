package others

import (
	"sort"
)

// 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{}
}

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

// 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	i := 0
	for i < len(flowerbed) {
		// 当前位置为0，且前一个和后一个位置也为0，表示可以种花
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1 // 种花
			count++
			i += 2 // 跳到下一个可种植位置的后一个位置
		} else {
			i++
		}
	}
	return count >= n
}
