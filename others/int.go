package others

import "fmt"

// 整数反转
func reverseInt(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

// 最大交换
func maximumSwap(num int) int {
	arr := []byte(fmt.Sprintf("%d", num))
	for i := 0; i < len(arr); i++ {
		maxIndex := i
		for j := maxIndex; j < len(arr); j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
			ans := 0
			for _, v := range arr {
				ans = ans*10 + int(v-'0')
			}
			return ans
		}
	}
	return num
}

// 位1的个数
func hammingWeight(num uint32) uint32 {
	var ans uint32
	for num != 0 {
		ans += (num & 1)
		num >>= 1
	}
	return ans
}
