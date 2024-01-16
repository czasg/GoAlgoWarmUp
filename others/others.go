package others

import "sort"

// 回旋镖
func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		stat := map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			dis := planeDistanceSquared(points[i], points[j])
			stat[dis] += 1
		}
		for _, v := range stat {
			// 对于n个点满足距离相等的情况下，相对于每一个点，存在n-1个点满足要求
			ans += v * (v - 1)
		}
	}
	return ans
}

func planeDistanceSquared(a, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return x*x + y*y
}

// 车队
type Car struct {
	Position int
	Speed    int
	Time     float64
}

func carFleet(target int, position []int, speed []int) int {
	if len(position) < 1 {
		return 0
	}
	if len(position) == 1 {
		return 1
	}
	cars := []Car{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{
			Position: position[i],
			Speed:    speed[i],
			Time:     float64((target - position[i])) / float64(speed[i]),
		})
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})
	fleet := 1
	headCar := cars[0]
	for i := 1; i < len(cars); i++ {
		if cars[i].Time > headCar.Time {
			fleet++
			headCar = cars[i]
		}
	}
	return fleet
}

// 爬楼梯
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

// 买股票
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	minPrice := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		curPrice := prices[i]
		if curPrice < minPrice {
			minPrice = curPrice
		}
		curProfit := curPrice - minPrice
		if curProfit > profit {
			profit = curProfit
		}
	}
	return profit
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

// 寻找两个有序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n := len(nums1) + len(nums2)
	midIndex := n / 2
	isZero := n%2 == 0
	i, j := 0, 0
	var cur int
	var pre int
	for i < len(nums1) && j < len(nums2) {
		pre = cur
		if nums1[i] < nums2[j] {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
		if midIndex == 0 {
			if isZero {
				return (float64(pre + cur)) / 2
			} else {
				return float64(cur)
			}
		}
		midIndex--
	}
	for k := j; k < len(nums2); k++ {
		pre = cur
		cur = nums2[k]
		if midIndex == 0 {
			if isZero {
				return (float64(pre + cur)) / 2
			} else {
				return float64(cur)
			}
		}
		midIndex--
	}
	return 0
}
