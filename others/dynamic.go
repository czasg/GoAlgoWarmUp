package others

// 最大子数组和 - 动态规划
func maxSubArray(nums []int) int {
	curSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], nums[i]+curSum)
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

// 爬楼梯-动态规划
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}

// 打家劫舍 - 不能偷取相邻的数值
func rob(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	arr := make([]int, n)
	arr[0] = nums[0]
	arr[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		arr[i] = max(arr[i-1], nums[i]+arr[i-2])
	}
	return arr[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
