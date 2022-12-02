package main

import "testing"

func Test_getAction(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAction(tt.args.a); got != tt.want {
				t.Errorf("getAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getExpectedAction(t *testing.T) {
	type args struct {
		op int
		s  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Draw", args: args{
			op: Rock,
			s:  Draw,
		}, want: Rock},
		{name: "Lose", args: args{
			op: Rock,
			s:  Lose,
		}, want: Scissors},
		{name: "Win", args: args{
			op: Scissors,
			s:  Win,
		}, want: Rock},
		{name: "Lose2", args: args{
			op: Paper,
			s:  Lose,
		}, want: Rock},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExpectedAction(tt.args.op, tt.args.s); got != tt.want {
				t.Errorf("getExpectedAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPoints(t *testing.T) {
	type args struct {
		op int
		me int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPoints(tt.args.op, tt.args.me); got != tt.want {
				t.Errorf("getPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
