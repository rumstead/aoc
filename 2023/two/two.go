package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	// +1 since index is the game number
	gameConfigs := make([]gameConfig, len(lines)+1)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		gcResult := createGameConfig(lines[i])
		gameConfigs[gcResult.num] = gcResult
	}
	// 12 red cubes, 13 green cubes, and 14 blue cubes
	total := 0
	powerSum := 0
	for _, gc := range gameConfigs {
		if gc.CanPlay(12, 13, 14) {
			total += gc.num
		}
		powerSum += gc.Power()
	}
	fmt.Println(total)
	fmt.Println(powerSum)
}

func createGameConfig(game string) gameConfig {
	words := strings.Split(game, ":")
	gameNum, err := strconv.Atoi(strings.Split(words[0], " ")[1])
	if err != nil {
		panic(err)
	}
	gc := gameConfig{num: gameNum}
	games := strings.Split(words[1], ";")
	gc.games = paresGames(games)
	return gc
}

func paresGames(games []string) []game {
	var gs []game
	for i := 0; i < len(games); i++ {
		rolls := strings.Split(games[i], ",")
		g := game{}
		for j := 0; j < len(rolls); j++ {
			colorCount := strings.Split(rolls[j], " ")
			// 0th element is empty
			count, err := strconv.Atoi(colorCount[1])
			if err != nil {
				panic(err)
			}
			g.addCount(count, colorCount[2])
		}
		gs = append(gs, g)
	}
	return gs
}

type gameConfig struct {
	games []game
	num   int
}

func (gc *gameConfig) Power() int {
	power := 1
	for _, c := range gc.minCount() {
		if c != 0 {
			power *= c
		}
	}
	if power == 1 {
		return 0
	}
	return power
}

// 0 = r, 1 = g, 2 = b
func (gc *gameConfig) minCount() []int {
	if len(gc.games) == 0 {
		return []int{}
	}
	mincube := []int{gc.games[0].r, gc.games[0].g, gc.games[0].b}
	for i := 0; i < len(gc.games); i++ {
		mincube[0] = setMinCube(mincube[0], gc.games[i].r)
		mincube[1] = setMinCube(mincube[1], gc.games[i].g)
		mincube[2] = setMinCube(mincube[2], gc.games[i].b)
	}
	return mincube
}

func setMinCube(c, n int) int {
	if c == 0 {
		return n
	}

	if c < n {
		return n
	}

	return c
}

func (gc *gameConfig) CanPlay(r, g, b int) bool {
	for i := 0; i < len(gc.games); i++ {
		if !gc.games[i].canPlay(r, g, b) {
			return false
		}
	}
	return true
}

type game struct {
	r, g, b int
}

func (g *game) canPlay(r, green, b int) bool {
	if g.r <= r && g.g <= green && g.b <= b {
		return true
	}
	return false
}

func (g *game) addCount(num int, color string) {
	switch color {
	case "red":
		g.r += num
	case "green":
		g.g += num
	case "blue":
		g.b += num
	}
}
