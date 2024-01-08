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
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + MaxDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + MaxDepth(root.Left)
	}
	return 1 + math.Max(MaxDepth(root.Left), MaxDepth(root.Right))
}

// 获取树的最小深度
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + MinDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + MinDepth(root.Left)
	}
	return 1 + math.Min(MinDepth(root.Left), MinDepth(root.Right))
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

// 判断两个树是否相同
func IsSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && IsSame(p.Left, q.Left) && IsSame(q.Right, q.Right)
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
func SortedArrayToBST(arr []int) *TreeNode {
	if len(arr) < 1 {
		return nil
	}
	mid := len(arr) / 2
	return &TreeNode{
		Val:   mid,
		Left:  SortedArrayToBST(arr[:mid]),
		Right: SortedArrayToBST(arr[mid+1:]),
	}
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
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}
	return root
}

// 删除值
func Delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = Delete(root.Left, val)
	} else if val > root.Val {
		root.Right = Delete(root.Right, val)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		rightMinVal := MinValueBST(root.Right).Val
		root.Val = rightMinVal
		root.Right = Delete(root.Right, rightMinVal)
	}
	return root
}

func MinValueBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 计算左叶子树之和
func SumLeftTree(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += SumLeftTree(root.Left)
	sum += SumLeftTree(root.Right)
	return
}
