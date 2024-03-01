package array

import (
	"reflect"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 0, 0},
				n:         3,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{0, 0, 0, 0, 0, 0},
				n:         4,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{0, 1, 0, 0, 0, 0},
				n:         2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{0, 1, 0, 0, 0, 0},
				n:         3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				nums: []int{3, 4, 5, 4, 5, 1, 2},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateYHSJ(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				numRows: 5,
			},
			want: [][]int{
				{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateYHSJ(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateYHSJ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge1(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums1: []int{1, 4, 6, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 7},
				n:     3,
			},
			want: []int{1, 2, 4, 5, 6, 7},
		},
		{
			name: "",
			args: args{
				nums1: []int{1, 2, 3, 4, 0, 0, 0},
				m:     4,
				nums2: []int{4, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge1(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge2() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}

func Test_merge2(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				a: []int{1, 3, 5, 7},
				b: []int{2, 4, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "",
			args: args{
				digits: []int{1, 2, 9},
			},
			want: []int{1, 3, 0},
		},
		{
			name: "",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantNum []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 4, 4, 5},
			},
			want:    5,
			wantNum: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6, 6, 6, 7},
			},
			want:    7,
			wantNum: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want || !reflect.DeepEqual(tt.args.nums[:got], tt.wantNum) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[:got], tt.wantNum)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantNum []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 3, 4},
				val:  3,
			},
			want:    4,
			wantNum: []int{1, 2, 2, 4},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 3, 5},
				val:  3,
			},
			want:    5,
			wantNum: []int{1, 2, 4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want || !reflect.DeepEqual(tt.args.nums[:got], tt.wantNum) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[:got], tt.wantNum)
			}
		})
	}
}

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setArr(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				a: []int{1, 2, 3, 4, 5},
				b: []int{4, 5, 6, 7, 8},
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setArr(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				nums: []int{-10, -5, -3, 0, 5, 7, 9},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -5,
					Left: &TreeNode{
						Val: -10,
					},
					Right: &TreeNode{
						Val: -3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(preorderTraversal(got), preorderTraversal(tt.want)) {
				t.Errorf("sortedArrayToBST() = %v, want %v", preorderTraversal(got), preorderTraversal(tt.want))
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 2, 3},
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			want: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 5, 6, 6, 6, 6, 6},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 3, 4, 5, 6, 7, 6, 8, 9},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
