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

// 合并两个有序链表
func Merge(list1 *ListNode, list2 *ListNode) *ListNode {
	panic("TODO")
}

// 翻转链表
func Reverse(head *ListNode) *ListNode {
	panic("TODO")
}

// 判断是否属于回文链表
func IsPalindrome(head *ListNode) bool {
	panic("TODO")
}
