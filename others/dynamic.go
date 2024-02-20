package others

import "proj/math"

// 最大子数组和 - 动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		// 在当前元素加入之前的子数组和加上当前元素的值进行比较
		currentSum = int(math.Max(nums[i], currentSum+nums[i]))
		// 更新全局最大和
		maxSum = int(math.Max(maxSum, currentSum))
	}
	return maxSum
}

// 爬楼梯-动态规划
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0], arr[1] = 1, 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
