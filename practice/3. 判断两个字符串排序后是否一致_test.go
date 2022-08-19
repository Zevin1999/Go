package practice

import "testing"

func TestSameRegroupString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "长度不一致",
			args: args{
				s1: "3701vcscgiwlcb",
				s2: "eue02-ei-20rei-ie1",
			},
			want: false,
		},
		{
			name: "排序后不一致",
			args: args{
				s1: "321123guguda",
				s2: "guguda32113567",
			},
			want: false,
		},
		{
			name: "排序后一致",
			args: args{
				s1: "3214567123123guguda",
				s2: "guguda1233211234567",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameRegroupString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("SameRegroupString() = %v, want %v", got, tt.want)
			}
		})
	}
}
