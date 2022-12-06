package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var void = struct{}{}

func main() {
	fmt.Println(getMarkerPos(4))
	fmt.Println(getMarkerPos(14))
}

func getMarkerPos(markerLength int) int {
	set := make(map[string]struct{})
	chars := strings.Split(input, "")
	i := 0
	for len(set) < markerLength {
		// reset our set
		set = make(map[string]struct{})
		possibleMarkers := chars[:markerLength]
		// pop each char off 1 at a time
		chars = chars[1:]
		// push the possible markers to a set
		for _, char := range possibleMarkers {
			set[char] = void
		}
		i++
	}
	return i + markerLength - 1
}
