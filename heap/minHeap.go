package heap

// MinHeap 是一个最小堆的结构
type MinHeap []int

// 创建一个新的最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// 插入操作，将元素插入到堆中
func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.heapifyUp()
}

// 上移操作，确保堆属性
func (h *MinHeap) heapifyUp() {
	arr := *h
	curIndex := len(arr) - 1
	parentIndex := getParentIndex(curIndex)
	for curIndex >= 0 && parentIndex >= 0 && arr[curIndex] < arr[parentIndex] {
		arr[curIndex], arr[parentIndex] = arr[parentIndex], arr[curIndex]
		curIndex = parentIndex
		parentIndex = getParentIndex(curIndex)
	}
}

// 弹出操作，从堆中移除并返回最小元素
func (h *MinHeap) Pop() int {
	arr := *h
	n := len(arr)
	pop := arr[0]
	arr[0] = arr[n-1]
	arr = arr[:n-1]
	*h = arr
	h.heapifyDown()
	return pop
}

// 下移操作，确保堆属性
func (h *MinHeap) heapifyDown() {
	arr := *h
	n := len(arr)
	curIndex := 0
	leftChildIndex := getLeftChildIndex(curIndex)
	rightChildIndex := getRightChildIndex(curIndex)
	for curIndex < n {
		minIndex := curIndex
		if leftChildIndex < n && arr[minIndex] > arr[leftChildIndex] {
			minIndex = leftChildIndex
		}
		if rightChildIndex < n && arr[minIndex] > arr[rightChildIndex] {
			minIndex = rightChildIndex
		}
		if minIndex == curIndex {
			return
		}
		arr[curIndex], arr[minIndex] = arr[minIndex], arr[curIndex]
		curIndex = minIndex
		leftChildIndex = getLeftChildIndex(curIndex)
		rightChildIndex = getRightChildIndex(curIndex)
	}
}
