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

func (l *ListNode) List() (arr []int) {
	var cur = l
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return
}

// 删除链表指定元素
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

// 删除链表的倒数第N个节点 - 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpNode := &ListNode{Next: head}
	fast, slow := tmpNode, tmpNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return tmpNode.Next
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
	arr := head.List()
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}
	return true
}

// 判断是否属于回文链表 - 双指针法
func IsPalindrome2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	current := slow
	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	firstHalf, secondHalf := head, prev
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return true
}

// 判断是否存在回环，
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

// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA := headA
	pB := headB
	if headA != headB {
		if headA == nil {
			headA = pB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = pA
		} else {
			headB = headB.Next
		}
	}
	return pA
}

// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return dummy.Next
}
