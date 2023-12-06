package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input-test.txt
var input string

func main() {
	grid := makeGrid(input)
	fmt.Println(grid)
	findEngine(grid)
}

func makeGrid(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for y, line := range lines {
		//dy := len(lines) - y - 1
		chars := strings.Split(line, "")
		grid[y] = make([]string, len(chars))
		for x, c := range chars {
			grid[y][x] = c
		}
	}
	return grid
}

func findEngine(grid [][]string) {
	for y, line := range grid {
		for x, c := range line {
			if utils.isDigit(c) {
				
			}
		}
	}
}
