package utils

import "fmt"

type Point[T any] struct {
	Y, X  int
	Value T
}

func (p Point[T]) String() string {
	return fmt.Sprintf("(Y: %d,X: %d)", p.Y, p.X)
}

func (p Point[T]) SamePoint(oPoint Point[T]) bool {
	return p.Y == oPoint.Y && p.X == oPoint.X
}

func (p Point[T]) IsAdjacent(oPoint Point[T]) bool {
	testPoint := p
	testPoint.X--
	testPoint.Y++
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test top
	testPoint = p
	testPoint.Y++
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test top right
	testPoint = p
	testPoint.X++
	testPoint.Y++
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test left
	testPoint = p
	testPoint.X--
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test right
	testPoint = p
	testPoint.X++
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test bottom left
	testPoint = p
	testPoint.X--
	testPoint.Y--
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test bottom
	testPoint = p
	testPoint.Y--
	if testPoint.SamePoint(oPoint) {
		return true
	}
	// test bottom right
	testPoint = p
	testPoint.X++
	testPoint.Y--
	if testPoint.SamePoint(oPoint) {
		return true
	}
	return false
}
