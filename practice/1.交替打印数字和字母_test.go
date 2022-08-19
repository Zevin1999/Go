package practice

import "testing"

func TestPrintNumberAndLetter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "交替打印数字和字母"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNumberAndLetter()
		})
	}
}
