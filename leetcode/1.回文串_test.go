package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "空字符串",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "长度为1的字符串",
			args: args{
				s: "Z",
			},
			want: true,
		},
		{
			name: "长度为2的非回文字符串",
			args: args{
				s: "zS",
			},
			want: false,
		},
		{
			name: "非回文字符串",
			args: args{
				s: "svsevwfh928f0239bcnks3091vo022dvvj",
			},
			want: false,
		},
		{
			name: "长度为2的有特殊字符的字符串",
			args: args{
				s: ":(",
			},
			want: true,
		},
		{
			name: "全是特殊字符的字符串",
			args: args{
				s: "《：：：：：：》",
			},
			want: true,
		},
		{
			name: "长度为2的回文字符串",
			args: args{
				s: "Zz",
			},
			want: true,
		},
		{
			name: "有特殊字符的回文字符串",
			args: args{
				s: "1234Abc (,^cBa;43:)><21",
			},
			want: true,
		},
		{
			name: "无特殊字符的回文字符串",
			args: args{
				s: "1234abccba4321",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
