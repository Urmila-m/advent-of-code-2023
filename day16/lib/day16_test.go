package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	currentState := State{
		Position:  Point{7, 7},
		GridElem:  '|',
		Direction: Up,
	}
	expectedOp := []State{
		{Position: Point{6, 7}, GridElem: '\\', Direction: Left},
	}
	assert.Equal(t, expectedOp, NextState(currentState, ParseFile("test_input.txt")))
}

func TestTraverse(t *testing.T) {
	matrix := ParseFile("test_input.txt")
	startState := State{
		Position:  Point{0, 0},
		Direction: Right,
		GridElem:  rune(matrix[0][0]),
	}
	for _, state := range Traverse(startState, matrix, make([]State, 0)) {
		state.Display()
	}
}

func TestFindUniqGridElems(t *testing.T) {
	matrix := ParseFile("test_input.txt")
	startPosition := Point{0, -1}
	startState := State{
		Position:  startPosition,
		Direction: Right,
		GridElem:  '.',
	}
	path := Traverse(startState, matrix, make([]State, 0))
	assert.Equal(t, 46, len(FindUniqGridElems(path)))
}

func TestFindInitialStates(t *testing.T) {
	matrix := ParseFile("test_input2.txt")
	initialStates := FindInitialStates(matrix)
	for _, state := range initialStates {
		state.Display()
	}
	assert.Equal(t, len(matrix)*2+len(matrix[0])*2, len(initialStates))
}

func TestFindMostEfficientConfig(t *testing.T) {
	_, maxNumOfGrids := FindMostEfficientConfig(ParseFile("test_input.txt"))
	assert.Equal(t, 51, maxNumOfGrids)
}
