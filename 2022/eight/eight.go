package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	grid := parseGrid(input)
	populateVisibility(grid)
}

func parseGrid(input string) [][]int {
	rows := strings.Split(input, "\n")
	grid := make([][]int, len(rows))
	for i, row := range rows {
		chars := strings.Split(row, "")
		for j, char := range chars {
			if j == 0 {
				grid[i] = make([]int, len(chars))
			}
			grid[i][j], _ = strconv.Atoi(char)
		}
	}
	return grid
}

func populateVisibility(grid [][]int) {
	dy := len(grid) - 1
	dx := len(grid[0]) - 1
	totalVisible := 0
	largestViewingDistance := 0
	for r, row := range grid {
		for c, tree := range row {
			cols := getCol(grid, r, c)
			tv := getTotalView(row, cols, r, c, tree)
			if tv > largestViewingDistance {
				largestViewingDistance = tv
			}

			// all outside trees can be seen
			if r == 0 || c == 0 || r == dy || c == dx {
				totalVisible++
				continue
			}
			isVisible := isVisibleRow(row, c) || isVisibleRow(cols, r)
			if isVisible {
				totalVisible++
			}

		}
	}
	fmt.Println(totalVisible)
	fmt.Println(largestViewingDistance)
}

func getTotalView(row []int, cols []int, r, c, tree int) int {
	left := row[0:c]
	right := row[c+1:]
	top := cols[0:r]
	bottom := cols[r+1:]

	totalViewLeft := 0
	for i := len(left) - 1; i >= 0; i-- {
		val := left[i]
		totalViewLeft++
		if tree <= val {
			break
		}
	}

	totalViewRight := 0
	for _, val := range right {
		totalViewRight++
		if tree <= val {
			break
		}
	}

	totalViewTop := 0
	for i := len(top) - 1; i >= 0; i-- {
		val := top[i]
		totalViewTop++
		if tree <= val {
			break
		}
	}

	totalViewBottom := 0
	for _, val := range bottom {
		totalViewBottom++
		if tree <= val {
			break
		}
	}

	total := totalViewLeft * totalViewRight * totalViewTop * totalViewBottom

	return total
}

func getCol(grid [][]int, r, c int) []int {
	var colToRow []int
	for i := 0; i < len(grid[r]); i++ {
		colToRow = append(colToRow, grid[i][c])
	}
	return colToRow
}

func isVisibleRow(row []int, x int) bool {
	if len(row) == 0 {
		return false
	}
	tree := row[x]
	var copyRow []int
	copyRow = append(copyRow, row...)
	left := copyRow[0:x]
	right := copyRow[x+1:]
	sort.Ints(left)
	sort.Ints(right)
	if len(left) > 0 && left[len(left)-1] < tree {
		return true
	}

	if len(right) > 0 && right[len(right)-1] < tree {
		return true
	}

	return false
}
