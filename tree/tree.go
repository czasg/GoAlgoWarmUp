package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreorderTraversal(root *TreeNode) []int {
	panic("TODO")
}

// 中序遍历
func InorderTraversal(root *TreeNode) []int {
	panic("TODO")
}

// 后序遍历
func PostorderTraversal(root *TreeNode) []int {
	panic("TODO")
}

// 获取树的最大深度
func MaxDepth(root *TreeNode) int {
	panic("TODO")
}

// 获取树的最小深度
func MinDepth(root *TreeNode) int {
	panic("TODO")
}

// 判断是否属于对称二叉树
func IsSymmetric(root *TreeNode) bool {
	panic("TODO")
}

// 判断是否属于平衡二叉树
func IsBalanced(root *TreeNode) bool {
	panic("TODO")
}

// 判断两个树是否相同
func IsSame(p *TreeNode, q *TreeNode) bool {
	panic("TODO")
}

// 翻转二叉树
func Mirror(root *TreeNode) *TreeNode {
	panic("TODO")
}

// 将数组转化为二叉搜索树
func NewTreeWithArray(arr []int) *TreeNode {
	panic("TODO")
}

// 搜索二叉树
func Search(root *TreeNode, val int) *TreeNode {
	list.New()
	panic("TODO")
}

// 插入值
func Insert(root *TreeNode, val int) *TreeNode {
	panic("TODO")
}

// 删除值
func Delete(root *TreeNode, val int) *TreeNode {
	panic("TODO")
}
