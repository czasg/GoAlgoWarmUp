package sort

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

type testCase struct {
	name string
	args args
	want []int
}

var testCases = []testCase{
	{
		name: "",
		args: args{
			arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		name: "",
		args: args{
			arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		name: "",
		args: args{
			arr: []int{2, 3, 7, 5, 6, 4, 8, 9, 0, 1},
		},
		want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
