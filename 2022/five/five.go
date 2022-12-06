package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var numbersRegex = regexp.MustCompile("[0-9]+")

func main() {
	// i dont feel like trying to dynamically parse the file...
	stacks := make([]list, 10)
	for i := range stacks {
		stacks[i] = newList()
	}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "move") {
			handleMove(stacks, line)
		} else if !strings.Contains(line, " 1   2   3") {
			updateStacks(stacks, line)
		}
	}
	var builder strings.Builder
	for _, s := range stacks {
		builder.WriteString(s.Pop())
	}
	fmt.Println(builder.String())
}

func handleMove(stacks []list, line string) {
	moves := numbersRegex.FindAllString(line, -1)
	if len(moves) != 3 {
		log.Fatalf("invalid moves: %s", moves)
	}
	quantity, err := strconv.Atoi(moves[0])
	if err != nil {
		log.Fatal(err)
	}
	from, err := strconv.Atoi(moves[1])
	if err != nil {
		log.Fatal(err)
	}
	to, err := strconv.Atoi(moves[2])
	if err != nil {
		log.Fatal(err)
	}
	doMovePart2(quantity, &stacks[from], &stacks[to])
}

func doMove(quantity int, from *list, to *list) {
	for i := 0; i < quantity; i++ {
		fromVal := from.Pop()
		to.Push(fromVal)
	}
}

func doMovePart2(quantity int, from *list, to *list) {
	elements := from.PopN(quantity)
	to.PushN(elements)
}

func updateStacks(stacks []list, line string) {
	lines := strings.Split(line, "")
	for i, s := 1, 1; i < len(lines) && s < len(stacks); i, s = i+4, s+1 {
		val := lines[i]
		if val != " " {
			stacks[s].Enqueue(val)
		}
	}
}

type list struct {
	data []string
}

// Pop removes and returns the head element
func (q *list) Pop() string {
	first := ""
	if len(q.data) > 0 {
		self := *q
		first = self.data[0]
		q.data = self.data[1:len(self.data)]
	}

	return first
}

// PopN removes and returns the first N elements or nothing if N > number of elements
func (q *list) PopN(n int) []string {
	var elements []string
	if len(q.data) >= n {
		self := *q
		elements = append(elements, self.data[0:n]...)
		q.data = self.data[n:len(self.data)]
	}
	return elements
}

// Enqueue appends an element to the tail
func (q *list) Enqueue(element string) {
	q.data = append(q.data, element)
}

// Push appends an element to the head
func (q *list) Push(element string) {
	q.data = append([]string{element}, q.data...)
}

// PushN appends N elements to the head
func (q *list) PushN(elements []string) {
	q.data = append(elements, q.data...)
}

func newList() list {
	s := list{}
	s.data = []string{}
	return s
}
