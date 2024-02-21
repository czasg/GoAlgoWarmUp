package others

// 整数反转
func reverseInt(x int) int {
	ans := 0
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

// 最大交换
func maximumSwap(num int) int {
	arr := []int{}
	n := num
	for n > 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	reverse(arr)
	set := map[int]int{}
	for index, v := range arr {
		set[v] = index
	}
	for index, v := range arr {
		for bigV := 9; bigV > v; bigV-- {
			vIndex, ok := set[bigV]
			if ok && vIndex > index {
				arr[index], arr[vIndex] = arr[vIndex], arr[index]
				return arr2num(arr)
			}
		}
	}
	return num
}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return
}

func arr2num(arr []int) int {
	ans := 0
	for _, v := range arr {
		ans = ans*10 + v
	}
	return ans
}
