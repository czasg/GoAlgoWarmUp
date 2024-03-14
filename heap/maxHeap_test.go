package heap

import (
	"sort"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap()
	arr := []int{0, 2, 5, 1, 2, 5, 46, 55, 6, 23, 421, 123, 5, 6, 5, 7, 2, 34, 213, 4, 46, 54, 6, 785, 78}
	for _, v := range arr {
		maxHeap.Insert(v)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for _, v := range arr {
		if v1 := maxHeap.Pop(); v1 != v {
			t.Errorf("NewMaxHeap() = %v, want %v", v, v1)
		}
	}
}

func TestNewMinHeap(t *testing.T) {
	minHeap := NewMinHeap()
	arr := []int{0, 2, 5, 1, 2, 5, 46, 55, 6, 23, 421, 123, 5, 6, 5, 7, 2, 34, 213, 4, 46, 54, 6, 785, 78}
	for _, v := range arr {
		minHeap.Insert(v)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for _, v := range arr {
		if v1 := minHeap.Pop(); v1 != v {
			t.Errorf("NewMaxHeap() = %v, want %v", v, v1)
		}
	}
}
