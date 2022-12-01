package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elves []int
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elves)
	// most calories
	fmt.Println(elves[len(elves)-1])
	// total top 3 calories
	top3 := elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Println(top3)
}
