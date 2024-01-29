package sort

// 冒泡排序
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 选择排序
func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		selectIndex := i
		for j := selectIndex; j < n; j++ {
			if arr[j] < arr[selectIndex] {
				selectIndex = j
			}
		}
		arr[i], arr[selectIndex] = arr[selectIndex], arr[i]
	}
	return arr
}

// 插入排序
func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j >= 1; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				continue
			}
			break
		}
	}
	return arr
}

// 希尔排序
func ShellSort(arr []int) []int {
	n := len(arr)
	step := n / 2
	for step > 0 {
		for i := step; i < n; i++ {
			for j := i; j >= step; j-- {
				if arr[j-step] > arr[j] {
					arr[j-step], arr[j] = arr[j], arr[j-step]
					continue
				}
				break
			}
		}
		step /= 2
	}
	return arr
}

// 快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	base := arr[0]
	left, right := []int{}, []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < base {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	ans := []int{}
	ans = append(ans, QuickSort(left)...)
	ans = append(ans, base)
	ans = append(ans, QuickSort(right)...)
	return ans
}

// 归并排序
func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := QuickSort(arr[:mid])
	right := QuickSort(arr[mid:])
	return Merge(left, right)
}

func Merge(a, b []int) []int {
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
