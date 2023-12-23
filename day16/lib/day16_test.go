package lib

import (
	"fmt"
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
	matrix := ParseFile("test_input3.txt")
	startState := State{
		Position:  Point{1, 0},
		Direction: Right,
		GridElem:  rune(matrix[1][0]),
	}
	fmt.Println(TraverseLoop(startState, matrix, make([]State, 0)))
}
