package main

import "testing"

func Test_rectangleArea(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"Area of rectangle is 63", args{10.5, 6}, 63},
		{"Area of rectangle is 10", args{5, 2}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rectangleArea(tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("rectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
