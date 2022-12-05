package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const a_ASCII = rune('a')
const A_ASCII = rune('A')

//go:embed input.txt
var input string

func main() {
	totalPriorities := 0
	groupTotalPriorities := 0
	var groupSet [][]string
	count := 0
	for _, line := range strings.Split(input, "\n") {
		chars := strings.Split(line, "")
		count++
		groupSet = append(groupSet, chars)
		if count%3 == 0 {
			groupTotalPriorities += handleGroup(groupSet)
			groupSet = [][]string{}
		}
		first := chars[:len(chars)/2]
		second := chars[len(chars)/2:]
		totalPriorities += getPriority(first, second)
	}
	// part one
	fmt.Println(totalPriorities)
	// part two
	fmt.Println(groupTotalPriorities)
}

func handleGroup(set [][]string) int {
	if len(set) != 3 {
		return 0
	}

	one := set[0]
	two := set[1]
	three := set[2]

	groupSet := make(map[string]struct{})
	for _, val := range one {
		if contains(val, two) && contains(val, three) {
			groupSet[val] = struct{}{}
		}
	}

	total := 0
	for k, _ := range groupSet {
		ascii := rune(k[0])
		total += determinePriority(ascii)
	}
	return total
}

func getPriority(first []string, second []string) int {
	set := make(map[string]struct{})
	for _, valI := range first {
		if contains(valI, second) {
			set[valI] = struct{}{}
		}
	}

	total := 0
	for k, _ := range set {
		ascii := []rune(k)[0]
		total += determinePriority(ascii)
	}
	return total
}

func determinePriority(ascii rune) int {
	// lowercase
	if ascii >= a_ASCII {
		return int(ascii) - int(a_ASCII) + 1 // priority starts at 1
	} else { // uppercase
		return int(ascii) - int(A_ASCII) + 27 // 1 + 26 (a-z)
	}
}

func contains(val string, data []string) bool {
	for _, valData := range data {
		if val == valData {
			return true
		}
	}
	return false
}
