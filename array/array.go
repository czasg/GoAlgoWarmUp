package array

import (
	"fmt"
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

// 查找数组中重复的元素 - 最多只有两个重复元素
func findDuplicates(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dupl := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			dupl = append(dupl, nums[i])
		}
	}
	return dupl
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	// 使用双指针，一个指向当前位置，一个指向要替换的位置
	current := 0
	replace := 1
	// 遍历数组
	for replace < len(nums) {
		// 如果当前元素不等于前一个元素，则将其替换到 replace 位置
		if nums[replace] != nums[current] {
			current++
			nums[current] = nums[replace]
		}
		replace++
	}
	// 返回新的数组长度
	return current + 1
}

// 移除数组中指定的的元素
func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// 搜索插入位置 - 二分查找法
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	ans := []int{1}
	ans = append(ans, digits...)
	return ans
}

// 合并两个有序数组 - 数据存放到nums1中
func merge1(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}

// 合并两个有序数组 - 新申请一个数组
func merge2(a, b []int) []int {
	arr := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}
	if i < len(a) {
		arr = append(arr, a[i:]...)
	} else {
		arr = append(arr, b[j:]...)
	}
	return arr
}

// 将有序数组转化为二叉搜索数
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

// 杨辉三角
func generateYHSJ(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{{1}}
	}
	rows := make([][]int, numRows)
	rows[0] = []int{1}
	for i := 1; i < numRows; i++ {
		headRow := rows[i-1]
		curRow := make([]int, i+1)
		curRow[0], curRow[i] = 1, 1
		for j := 1; j < i; j++ {
			curRow[j] = headRow[j-1] + headRow[j]
		}
		rows[i] = curRow
	}
	return rows
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
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

// 买卖股票的最佳时机 - 买入卖出，求最大利润 - 当前价格减去历史最低价格
func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		curProfit := prices[i] - minPrice
		if curProfit > profit {
			profit = curProfit
		}
	}
	return profit
}

// 只出现一次的数字 - 位运算
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 找出出现次数最多的元素
func majorityElement(nums []int) int {
	stat := map[int]int{}
	for _, num := range nums {
		stat[num]++
	}
	maxK := 0
	maxV := 0
	for k, v := range stat {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	return maxK
}

// 有序数组求汇总区间
func summaryRanges(nums []int) []string {
	ans := []string{}
	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
			continue
		}
		if start == end {
			ans = append(ans, fmt.Sprintf("%d", start))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", start, end))
		}
		start = nums[i]
		end = nums[i]
	}
	if start == end {
		ans = append(ans, fmt.Sprintf("%d", start))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", start, end))
	}
	return ans
}

// 丢失的数字 - 找出这个范围内没有出现在数组中的那个数
func missingNumber(nums []int) int {
	n := len(nums)
	ans := n * (n + 1) / 2
	for _, num := range nums {
		ans -= num
	}
	return ans
}

// 移动零 - 将所有 0 移动到数组的末尾
func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[index] = nums[i]
		index++
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
