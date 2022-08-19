package practice

import "testing"

func TestReplaceBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name:  "s包含数字",
			args:  args{s: "aBfdsigfuf17301"},
			want:  "aBfdsigfuf17301",
			want1: false,
		},
		{
			name:  "s替换空格",
			args:  args{s: "Ab Cd ef HaHa HaHa"},
			want:  "Ab%20Cd%20ef%20HaHa%20HaHa",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ReplaceBlank(tt.args.s)
			if got != tt.want {
				t.Errorf("ReplaceBlank() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReplaceBlank() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
