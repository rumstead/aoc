package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/rumstead/aoc/2023/utils"
)

//go:embed input-test.txt
var input string

func main() {
	nums, specialPoints := parseGrid(input)
	//fmt.Println(grid)
	engineNums := findEngine(nums, specialPoints)
	total := 0
	for _, num := range engineNums {
		total += num
	}
	fmt.Println(total)
}

func parseGrid(input string) ([]map[int][]utils.Point[string], map[string]utils.Point[string]) {
	var specialMap = map[string]utils.Point[string]{}
	var nums []map[int][]utils.Point[string]
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		strNum := ""
		var numPoints []utils.Point[string]
		for x, c := range chars {
			if c == "." {
				continue
			}
			point := utils.Point[string]{Y: y, X: x, Value: c}
			if !utils.IsDigit(c) {
				key := getSpecialPointKey(point)
				specialMap[key] = point
			} else {
				strNum += c
				numPoints = append(numPoints, point)
			}
			// look ahead to see if we should stop processing this number
			if utils.IsDigit(c) && (x+1 < len(chars) && !utils.IsDigit(chars[x+1]) || x+1 == len(chars)) {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					panic(err)
				}
				var numToPoints = map[int][]utils.Point[string]{}
				numToPoints[num] = numPoints
				nums = append(nums, numToPoints)
				strNum = ""
				numPoints = []utils.Point[string]{}

			}
		}
	}
	return nums, specialMap
}

func getSpecialPointKey(point utils.Point[string]) string {
	return fmt.Sprintf("%d,%d", point.Y, point.X)
}

func findEngine(nums []map[int][]utils.Point[string], specialPoints map[string]utils.Point[string]) []int {
	var engineNums []int
	for _, numToPoints := range nums {
		for num, points := range numToPoints {
			if isAdjacent(points, specialPoints) {
				engineNums = append(engineNums, num)
			}
		}
	}
	return engineNums
}

func isAdjacent(points []utils.Point[string], searchPoints map[string]utils.Point[string]) bool {
	for _, p := range points {
		// test top left
		if pointExists(p.X-1, p.Y+1, searchPoints) {
			return true
		}
		// test top
		if pointExists(p.X, p.Y+1, searchPoints) {
			return true
		}
		// test top right
		if pointExists(p.X+1, p.Y+1, searchPoints) {
			return true
		}
		// test left
		if pointExists(p.X-1, p.Y, searchPoints) {
			return true
		}
		// test right
		if pointExists(p.X+1, p.Y, searchPoints) {
			return true
		}
		// test bottom left
		if pointExists(p.X-1, p.Y-1, searchPoints) {
			return true
		}
		// test bottom
		if pointExists(p.X, p.Y-1, searchPoints) {
			return true
		}
		// test bottom right
		if pointExists(p.X+1, p.Y-1, searchPoints) {
			return true
		}
	}
	return false
}

func pointExists(x, y int, points map[string]utils.Point[string]) bool {
	testPoint := utils.Point[string]{X: x, Y: y}
	_, ok := points[getSpecialPointKey(testPoint)]
	return ok
}
