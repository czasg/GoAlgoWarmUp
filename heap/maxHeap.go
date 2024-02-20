package heap

// MaxHeap 是一个最大堆的结构
type MaxHeap struct {
	array []int
	size  int
}

// 创建一个新的最大堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// 获取父节点的索引
func parentIndex(i int) int {
	return (i - 1) / 2
}

// 获取左子节点的索引
func leftChildIndex(i int) int {
	return 2*i + 1
}

// 获取右子节点的索引
func rightChildIndex(i int) int {
	return 2*i + 2
}

// 上移操作，确保堆属性
func (h *MaxHeap) heapifyUp() {
	currentIndex := h.size - 1

	for currentIndex > 0 && h.array[currentIndex] > h.array[parentIndex(currentIndex)] {
		h.array[currentIndex], h.array[parentIndex(currentIndex)] = h.array[parentIndex(currentIndex)], h.array[currentIndex]
		currentIndex = parentIndex(currentIndex)
	}
}

// 插入操作，将元素插入到堆中
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.size++
	h.heapifyUp()
}

// 下移操作，确保堆属性
func (h *MaxHeap) heapifyDown(index int) {
	maximum := index
	left := leftChildIndex(index)
	right := rightChildIndex(index)

	if left < h.size && h.array[left] > h.array[maximum] {
		maximum = left
	}

	if right < h.size && h.array[right] > h.array[maximum] {
		maximum = right
	}

	if index != maximum {
		h.array[index], h.array[maximum] = h.array[maximum], h.array[index]
		h.heapifyDown(maximum)
	}
}

// 弹出操作，从堆中移除并返回最大元素
func (h *MaxHeap) Pop() int {
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
