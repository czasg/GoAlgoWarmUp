package heap

// MaxHeap 是一个最大堆的结构
type MaxHeap []int

// 创建一个新的最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 插入操作，将元素插入到堆中
func (h *MaxHeap) Insert(value int) {
	*h = append(*h, value)
	h.heapifyUp()
}

// 上移操作，确保堆属性
func (h *MaxHeap) heapifyUp() {
	arr := *h
	curIndex := len(arr) - 1
	parentIndex := getParentIndex(curIndex)
	for parentIndex >= 0 && arr[curIndex] > arr[parentIndex] {
		arr[curIndex], arr[parentIndex] = arr[parentIndex], arr[curIndex]
		curIndex = parentIndex
		parentIndex = getParentIndex(curIndex)
	}
}

// 弹出操作，从堆中移除并返回最大元素
func (h *MaxHeap) Pop() int {
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
func (h *MaxHeap) heapifyDown() {
	arr := *h
	n := len(arr)
	curIndex := 0
	leftChildIndex := getLeftChildIndex(curIndex)
	rightChildIndex := getRightChildIndex(curIndex)
	for curIndex < n {
		maxIndex := curIndex
		if leftChildIndex < n && arr[leftChildIndex] > arr[maxIndex] {
			maxIndex = leftChildIndex
		}
		if rightChildIndex < n && arr[rightChildIndex] > arr[maxIndex] {
			maxIndex = rightChildIndex
		}
		if maxIndex == curIndex {
			return
		}
		arr[curIndex], arr[maxIndex] = arr[maxIndex], arr[curIndex]
		curIndex = maxIndex
		leftChildIndex = getLeftChildIndex(curIndex)
		rightChildIndex = getRightChildIndex(curIndex)
	}
}
