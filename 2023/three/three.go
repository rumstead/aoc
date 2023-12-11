package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/rumstead/aoc/2023/utils"
)

//go:embed input.txt
var input string

// Not proud of this one but got a little bored with it and just wanted to get it done
func main() {
	nums, specialPoints := parseGrid(input)
	//partOne(nums, specialPoints)
	partTwo(nums, specialPoints)
}

func partTwo(nums []map[int][]utils.Point[string], points map[string]utils.Point[string]) {
	var gears = map[string]utils.Point[string]{}
	for _, sp := range points {
		if sp.Value == "*" {
			gears[getSpecialPointKey(sp)] = sp
		}
	}

	var strPoints = map[string]utils.Point[string]{}
	var pointsToNum = map[string]int{}
	for _, numToPoints := range nums {
		for num, p := range numToPoints {
			for _, point := range p {
				strPoints[getSpecialPointKey(point)] = point
				pointsToNum[getSpecialPointKey(point)] = num
			}
		}
	}
	total := 0
	for _, g := range gears {
		cp := adjacentPoints(g, strPoints, getSpecialPointKey)
		if len(cp) >= 2 {
			var numSet []int
			testPoint := utils.Point[string]{
				Y: -1,
				X: -1,
			}
			for i, p := range cp {
				if num, ok := pointsToNum[getSpecialPointKey(p)]; ok {
					if testPoint.Y != p.Y {
						numSet = append(numSet, num)
					} else if testPoint.X != p.X-1 {
						numSet = append(numSet, num)
					}
					testPoint = cp[i]
				}
			}
			if len(numSet) != 2 && len(numSet) > 0 {
				fmt.Printf("%v: numSet not 2 %v\n", g, numSet)
			} else {
				fmt.Printf("%v: numSet 2 %v\n", g, numSet)
				total += numSet[0] * numSet[1]
			}
		}
	}
	fmt.Println(total)
}

func partOne(nums []map[int][]utils.Point[string], points map[string]utils.Point[string]) {
	engineNums := findEngine(nums, points)
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
			if isEngineAdjacent(points, specialPoints, getSpecialPointKey) {
				engineNums = append(engineNums, num)
			}
		}
	}
	return engineNums
}

func isEngineAdjacent(points []utils.Point[string], searchPoints map[string]utils.Point[string], key func(p utils.Point[string]) string) bool {
	for _, p := range points {
		if len(adjacentPoints(p, searchPoints, key)) > 0 {
			return true
		}
	}
	return false
}

func adjacentPoints(p utils.Point[string], searchPoints map[string]utils.Point[string], key func(p utils.Point[string]) string) []utils.Point[string] {
	var points []utils.Point[string]
	// test top left
	testPoint := p
	testPoint.X--
	testPoint.Y++
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test top
	testPoint = p
	testPoint.Y++
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test top right
	testPoint = p
	testPoint.X++
	testPoint.Y++
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test left
	testPoint = p
	testPoint.X--
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test right
	testPoint = p
	testPoint.X++
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test bottom left
	testPoint = p
	testPoint.X--
	testPoint.Y--
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test bottom
	testPoint = p
	testPoint.Y--
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	// test bottom right
	testPoint = p
	testPoint.X++
	testPoint.Y--
	if pointExists(testPoint, searchPoints, key) {
		points = append(points, testPoint)
	}
	return points
}

func pointExists(p utils.Point[string], points map[string]utils.Point[string], key func(p utils.Point[string]) string) bool {
	_, ok := points[key(p)]
	return ok
}
