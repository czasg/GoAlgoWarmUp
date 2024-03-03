package others

// 最大子数组和 - 动态规划
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum := nums[i]
		if maxSum+curSum > curSum {
			curSum = maxSum + curSum
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
	//if len(nums) == 0 {
	//	return 0
	//}
	//maxSum := nums[0]
	//currentSum := nums[0]
	//for i := 1; i < len(nums); i++ {
	//	// 在当前元素加入之前的子数组和加上当前元素的值进行比较
	//	currentSum = int(math.Max(nums[i], currentSum+nums[i]))
	//	// 更新全局最大和
	//	maxSum = int(math.Max(maxSum, currentSum))
	//}
	//return maxSum
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

// 打家劫舍 - 不能偷取相邻的数值
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}
