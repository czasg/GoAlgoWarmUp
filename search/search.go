package search

// 二分查找算法
func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
