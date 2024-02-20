package heap

// MinHeap 是一个最小堆的结构
type MinHeap struct {
	array []int
	size  int
}

// 创建一个新的最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// 上移操作，确保堆属性
func (h *MinHeap) heapifyUp() {
	currentIndex := h.size - 1

	for currentIndex > 0 && h.array[currentIndex] < h.array[parentIndex(currentIndex)] {
		h.array[currentIndex], h.array[parentIndex(currentIndex)] = h.array[parentIndex(currentIndex)], h.array[currentIndex]
		currentIndex = parentIndex(currentIndex)
	}
}

// 插入操作，将元素插入到堆中
func (h *MinHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp()
}

// 下移操作，确保堆属性
func (h *MinHeap) heapifyDown(index int) {
	minimum := index
	left := leftChildIndex(index)
	right := rightChildIndex(index)

	if left < h.size && h.array[left] < h.array[minimum] {
		minimum = left
	}

	if right < h.size && h.array[right] < h.array[minimum] {
		minimum = right
	}

	if index != minimum {
		h.array[index], h.array[minimum] = h.array[minimum], h.array[index]
		h.heapifyDown(minimum)
	}
}

// 弹出操作，从堆中移除并返回最小元素
func (h *MinHeap) Pop() int {
	if h.size == 0 {
		return -1
	}

	root := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--
	h.heapifyDown(0)

	return root
}
