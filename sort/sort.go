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
	panic("TODO")
}

// 希尔排序
func ShellSort(arr []int) []int {
	panic("TODO")
}

// 快速排序
func QuickSort(arr []int) []int {
	panic("TODO")
}

// 归并排序
func MergeSort(arr []int) []int {
	panic("TODO")
}
