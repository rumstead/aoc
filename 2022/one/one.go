package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var elves []int
	total := 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += calorie
		} else {
			// new elf
			elves = append(elves, total)
			total = 0
		}
	}

	sort.Ints(elves)
	// most calories
	fmt.Println(elves[len(elves)-1])
	// total top 3 calories
	top3 := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Println(top3)
}
