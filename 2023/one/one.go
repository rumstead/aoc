package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	total := 0
	inputArr := strings.Split(input, "\n")
	for _, line := range inputArr {
		if line == "" {
			continue
		}
		first, last, digit := -1, -1, 0
		for i := 0; i < len(line); i++ {
			char := rune(line[i])
			if unicode.IsDigit(char) {
				digit = int(char - '0')
			} else {
				// part 2
				digit = wordNumbers(line[i:])
			}
			if first < 0 {
				first = digit
			}
			if first > 0 && digit > 0 {
				last = digit
			}
			digit = 0
		}
		if last < 0 {
			last = first
		}
		total += first*10 + last
	}
	fmt.Println(total)
}

func wordNumbers(line string) int {
	if strings.HasPrefix(line, "zero") {
		return 0
	} else if strings.HasPrefix(line, "one") {
		return 1
	} else if strings.HasPrefix(line, "two") {
		return 2
	} else if strings.HasPrefix(line, "three") {
		return 3
	} else if strings.HasPrefix(line, "four") {
		return 4
	} else if strings.HasPrefix(line, "five") {
		return 5
	} else if strings.HasPrefix(line, "six") {
		return 6
	} else if strings.HasPrefix(line, "seven") {
		return 7
	} else if strings.HasPrefix(line, "eight") {
		return 8
	} else if strings.HasPrefix(line, "nine") {
		return 9
	}
	return -1
}
