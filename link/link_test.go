package link

import (
	"reflect"
	"testing"
)

func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    3,
			},
			want: createLinkedList([]int{1, 2, 4, 5}),
		},
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    1,
			},
			want: createLinkedList([]int{1, 2, 3, 4}),
		},
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 4, 5}),
				n:    5,
			},
			want: createLinkedList([]int{2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got.List(), tt.want.List()) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 2, 3, 2, 1}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome2(tt.args.head); got != tt.want {
				t.Errorf("IsPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuplicate(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 1, 2, 2, 3, 3, 4, 4}),
			},
			want: createLinkedList([]int{1, 2, 3, 4}),
		},
		{
			name: "",
			args: args{
				head: createLinkedList([]int{1, 1, 2, 2, 3, 3, 4, 5}),
			},
			want: createLinkedList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate(tt.args.head); !reflect.DeepEqual(got.List(), tt.want.List()) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
