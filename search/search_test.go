package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
				val: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
				val: 8,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
