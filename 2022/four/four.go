package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	subSetCount, overlapCount := 0, 0
	for _, line := range strings.Split(input, "\n") {
		sections := strings.Split(line, ",")
		sectionOne := sections[0]
		sectionTwo := sections[1]
		sectionSetOne := sectionToArray(sectionOne)
		sectionSetTwo := sectionToArray(sectionTwo)
		if isSubset(sectionSetOne, sectionSetTwo) || isSubset(sectionSetTwo, sectionSetOne) {
			subSetCount++
		}

		if hasOverlap(sectionSetOne, sectionSetTwo) {
			overlapCount++
		}
	}
	// part one
	fmt.Println(subSetCount)
	// part two
	fmt.Println(overlapCount)
}

func isSubset(one, two []int) bool {
	for x := 0; x < len(one) || x < len(two); x++ {
		oneVal, twoVal := 0, 0
		if x < len(one) {
			oneVal = one[x]
		}

		if x < len(two) {
			twoVal = two[x]
		}

		if oneVal-twoVal > 0 {
			return false
		}
	}
	return true
}

func hasOverlap(one, two []int) bool {
	for x := 0; x < len(one) || x < len(two); x++ {
		oneVal, twoVal := 0, 0
		if x < len(one) {
			oneVal = one[x]
		}

		if x < len(two) {
			twoVal = two[x]
		}

		if oneVal+twoVal > 1 {
			return true
		}
	}
	return false
}

func sectionToArray(section string) []int {
	sectionRange := strings.Split(section, "-")
	lengthStr := sectionRange[len(sectionRange)-1]
	end, err := strconv.Atoi(lengthStr)
	if err != nil {
		log.Fatalln(err)
	}
	start, err := strconv.Atoi(sectionRange[0])
	sectionSlice := make([]int, end+1)
	for i := start; i < len(sectionSlice); i++ {
		sectionSlice[i] = 1
	}

	return sectionSlice
}
