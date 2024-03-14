package heap

// 获取父节点的索引
func getParentIndex(i int) int {
	return (i - 1) / 2
}

// 获取左子节点的索引
func getLeftChildIndex(i int) int {
	return 2*i + 1
}

// 获取右子节点的索引
func getRightChildIndex(i int) int {
	return 2*i + 2
}
