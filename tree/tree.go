package tree

import "proj/math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := []int{}
	arr = append(arr, root.Val)
	arr = append(arr, PreorderTraversal(root.Left)...)
	arr = append(arr, PreorderTraversal(root.Right)...)
	return arr
}

// 中序遍历
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := []int{}
	arr = append(arr, InorderTraversal(root.Left)...)
	arr = append(arr, root.Val)
	arr = append(arr, InorderTraversal(root.Right)...)
	return arr
}

// 后序遍历
func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := []int{}
	arr = append(arr, InorderTraversal(root.Left)...)
	arr = append(arr, InorderTraversal(root.Right)...)
	arr = append(arr, root.Val)
	return arr
}

// 获取树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + math.Max(
		MaxDepth(root.Left),
		MaxDepth(root.Right),
	)
}

// 获取树的最小深度
func MinDepth(root *TreeNode) int {
	panic("TODO")
}

// 判断是否属于对称二叉树
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return (left.Val == right.Val) && isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)
}

// 判断是否属于平衡二叉树
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := MaxDepth(root.Left)
	rightHeight := MaxDepth(root.Right)
	return math.Abs(leftHeight-rightHeight) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

// 判断两个树是否相同
func IsSame(p *TreeNode, q *TreeNode) bool {
	panic("TODO")
}

// 翻转二叉树
func Mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = Mirror(root.Right)
	root.Right = Mirror(root.Left)
	return root
}

// 将数组转化为二叉搜索树
func NewTreeWithArray(arr []int) *TreeNode {
	panic("TODO")
}

// 搜索二叉树
func Search(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return Search(root.Left, val)
	} else {
		return Search(root.Right, val)
	}
}

// 插入值
func Insert(root *TreeNode, val int) *TreeNode {
	panic("TODO")
}

// 删除值
func Delete(root *TreeNode, val int) *TreeNode {
	panic("TODO")
}

// 计算左叶子树之和
func SumLeftTree(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	// 判断当前节点的左叶子是否属于叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	// 递归处理
	sum += SumLeftTree(root.Left)
	sum += SumLeftTree(root.Right)
	return
}
