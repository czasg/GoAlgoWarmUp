package stack

import "testing"

func Test_reverseParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "(abcd)",
			},
			want: "dcba",
		},
		{
			name: "",
			args: args{
				s: "(u(love)i)",
			},
			want: "iloveu",
		},
		{
			name: "",
			args: args{
				s: "(i(op(pa))hs)",
			},
			want: "shopapi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
