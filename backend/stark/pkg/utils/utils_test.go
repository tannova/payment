package utils

import (
	"fmt"
	"testing"
)

func TestConvertIntToStr(t *testing.T) {
	type args struct {
		i      int64
		digits int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1 digit",
			args: args{i: 1, digits: 1},
			want: "1",
		},
		{
			name: "2 digits",
			args: args{i: 1, digits: 2},
			want: "01",
		},
		{
			name: "3 digits",
			args: args{i: 1, digits: 3},
			want: "001",
		},
		{
			name: "10 digits",
			args: args{i: 1, digits: 10},
			want: "0000000001",
		},
		{
			name: "max digits",
			args: args{i: 2147483647, digits: 8},
			want: "21474836",
		},

		{
			name: "random digits",
			args: args{i: 123456, digits: 2},
			want: "12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertIntToStr(tt.args.i, tt.args.digits)
			if got != tt.want {
				t.Errorf("ConvertIntToStr() = %v, want %v", got, tt.want)
			}
			fmt.Println("got =", got, "want =", tt.want)
		})
	}
}

func TestRandomDigitStr(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "8 digits",
			args: args{digits: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomDigitStr(tt.args.digits)
			fmt.Println("got =", got)
		})
	}
}
