package main

import "testing"

func Test_isVisibleCol(t *testing.T) {
	type args struct {
		grid [][]int
		r    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "left_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,3,3,2},{3,3,5,4,9},{3,5,3,9,0}}, r: 1, c: 1}, want: true},
		{name: "right_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,3,3,2},{3,3,5,4,9},{3,5,3,9,0}}, r: 1, c: 2}, want: true},
		{name: "middle_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,3,3,2},{3,3,5,9,9},{3,5,3,8,0}}, r: 3, c: 3}, want: true},
		{name: "left_not_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,5,3,2},{3,3,5,4,9},{3,5,3,9,0}}, r: 2, c: 1}, want: false},
		{name: "right_not_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,3,3,2},{3,3,5,4,9},{3,5,3,9,0}}, r: 3, c: 1}, want: false},
		{name: "middle_not_visible", args: args{grid: [][]int{{3,0,3,7,3},{2,5,5,1,2},{6,5,3,3,2},{3,3,5,9,9},{3,5,3,8,0}}, r: 2, c: 1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVisibleCol(tt.args.grid, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("isVisibleCol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVisibleRow(t *testing.T) {
	type args struct {
		row  []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{
			row:  []int{},
			x:    0,
		}, want: false},
		{name: "left_visible", args: args{
			row:  []int{9 ,7, 8, 7, 7},
			x:    0,
		}, want: true},
		{name: "right_visible", args: args{
			row:  []int{7,7,8,8,9},
			x:    4,
		}, want: true},
		{name: "middle_visible", args: args{
			row:  []int{7,7,9,8,8},
			x:    2,
		}, want: true},
		{name: "equal", args: args{
			row:  []int{9,9,9,9,9},
			x:    2,
		}, want: false},
		{name: "left_not_visible", args: args{
			row:  []int{5 ,7, 8, 7, 7},
			x:    0,
		}, want: false},
		{name: "right_not_visible", args: args{
			row:  []int{7,7,8,8,5},
			x:    4,
		}, want: false},
		{name: "middle_not_visible", args: args{
			row:  []int{7,7,5,8,8},
			x:    2,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVisibleRow(tt.args.row, tt.args.x); got != tt.want {
				t.Errorf("isVisibleRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_populateVisibility(t *testing.T) {
	type args struct {
		grid       [][]int
		visibility [][]int
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			populateVisibility(tt.args.grid, tt.args.visibility)
		})
	}
}
