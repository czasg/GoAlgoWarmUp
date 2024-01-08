package link

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var b strings.Builder
	var ll = l
	for ll != nil {
		b.WriteString(fmt.Sprintf("%d", ll.Val))
		ll = ll.Next
	}
	return b.String()
}

// 移除链表元素
func Remove(head *ListNode, val int) *ListNode {
	node := &ListNode{
		Val:  0,
		Next: head,
	}
	for node.Next != nil {
		if node.Next.Val == val {
			node = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return node.Next
}

// 合并两个有序链表
func Merge(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	curNode := node
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curNode.Next = list1
			list1 = list1.Next
		} else {
			curNode.Next = list2
			list2 = list2.Next
		}
		curNode = curNode.Next
	}
	if list1 != nil {
		curNode.Next = list1
	} else if list2 != nil {
		curNode.Next = list2
	}
	return node.Next
}

// 翻转链表
func Reverse(head *ListNode) *ListNode {
	var prevNode *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = prevNode
		prevNode = head
		head = nextNode
	}
	return prevNode
}

// 有序链表-删除重复元素
func Duplicate(head *ListNode) *ListNode {
	node := &ListNode{Next: head}
	curNode := node
	for curNode.Next != nil && curNode.Next.Next != nil {
		if curNode.Next.Val == curNode.Next.Next.Val {
			curNode = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return node.Next
}

// 判断是否属于回文链表
func IsPalindrome(head *ListNode) bool {
	arr := NodeToArray(head)
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}
	return true
}

func NodeToArray(head *ListNode) (arr []int) {
	if head == nil {
		return
	}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return
}

// 判断是否存在回环
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	oneStep := head
	twoStep := head.Next
	for oneStep != nil && twoStep != nil && twoStep.Next != nil { // 双指针法
		if oneStep == twoStep {
			return true
		}
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}
	return false
}
