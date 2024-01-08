package search

// 二分查找算法
func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if val < arr[mid] {
			right = mid
		} else if val > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
