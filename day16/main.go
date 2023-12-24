package main

import (
	"day16/lib"
	"fmt"
)

func puzzleOne(matrix []string) int {
	startPosition := lib.Point{X: 0, Y: -1}
	startState := lib.State{
		GridElem:  '.',
		Direction: lib.Right,
		Position:  startPosition,
	}
	path := lib.Traverse(startState, matrix, make([]lib.State, 0))
	for _, s := range path {
		s.Display()
	}
	return len(lib.FindUniqGridElems(path))
}

func puzzleTwo(matrix []string) int {
	_, maxNumOfGrids := lib.FindMostEfficientConfig(matrix)
	return maxNumOfGrids
}

func main() {
	matrix := lib.ParseFile("puzzle_input.txt")
	fmt.Println(puzzleOne(matrix))
	fmt.Println(puzzleTwo(matrix))
}
