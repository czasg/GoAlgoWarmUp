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
