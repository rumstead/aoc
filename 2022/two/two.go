package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
)

const (
	// Rock starts at 1 for points
	Rock int = iota + 1
	Paper
	Scissors
	Draw string = "Y"
	Lose string = "X"
	Win  string = "Z"
)

//go:embed input.txt
var input string

func main() {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		words := strings.Fields(line)
		if len(words) != 2 {
			log.Fatalf("%s is invalid!\n", line)
		}
		op := getAction(words[0])
		// part one
		//me := getAction(words[1])
		// part two
		me := getExpectedAction(op, words[1])
		total += getPoints(op, me)
	}

	fmt.Println(total)
}

func getExpectedAction(op int, s string) int {
	switch s {
	case Draw:
		return op
	case Lose:
		if op == Rock {
			return Scissors
		}
		return op - 1
	case Win:
		if op == Scissors {
			return Rock
		}
		return op + 1
	}
	return 0
}

func getAction(a string) int {
	a = strings.ToLower(a)
	switch a {
	case "a", "x":
		return Rock
	case "b", "y":
		return Paper
	case "c", "z":
		return Scissors
	}
	return Scissors
}

func getPoints(op, me int) int {
	val := op - me
	// draw
	if val == 0 {
		return 3 + me
	}

	// lose
	if val > 0 && val < 2 || val == -2 {
		return 0 + me
	}

	// win
	return 6 + me
}
