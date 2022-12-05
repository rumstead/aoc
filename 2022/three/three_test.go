package main

import "testing"

func Test_getPriority(t *testing.T) {
	type args struct {
		first  []string
		second []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "A", args: args{
			first:  []string{"A"},
			second: []string{"A"},
		}, want: 27},
		{name: "a", args: args{
			first:  []string{"a"},
			second: []string{"a"},
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPriority(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("getPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
