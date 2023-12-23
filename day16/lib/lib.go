package lib

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func ParseFile(filePath string) []string {
	file, fErr := os.Open(filePath)
	if fErr != nil {
		log.Fatal(fErr)
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	matrix := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}
	return matrix
}

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	Up    Direction = "up"
	Down  Direction = "down"
)

type Point struct {
	x int
	y int
}

type State struct {
	GridElem  rune
	Position  Point
	Direction Direction
}

func NextState(s1 State, matrix []string) []State {
	newStates := make([]State, 0)
	nextElemPosition := Point{-1, -1}
	directions := make([]Direction, 0)
	var nextElem rune
	switch s1.Direction {
	case Right:
		if (s1.Position.y + 1) < len(matrix[0]) {
			nextElemPosition = Point{s1.Position.x, s1.Position.y + 1}
			nextElem = rune(matrix[nextElemPosition.x][nextElemPosition.y])
			switch nextElem {
			case '|':
				directions = append(directions, Up, Down)
			case '\\':
				directions = append(directions, Down)
			case '/':
				directions = append(directions, Up)
			default:
				directions = append(directions, Right)
			}
		}
	case Left:
		if (s1.Position.y - 1) >= 0 {
			nextElemPosition = Point{s1.Position.x, s1.Position.y - 1}
			nextElem = rune(matrix[nextElemPosition.x][nextElemPosition.y])
			switch nextElem {
			case '|':
				directions = append(directions, Up, Down)
			case '\\':
				directions = append(directions, Up)
			case '/':
				directions = append(directions, Down)
			default:
				directions = append(directions, Left)
			}
		}
	case Up:
		if (s1.Position.x - 1) >= 0 {
			nextElemPosition = Point{s1.Position.x - 1, s1.Position.y}
			nextElem = rune(matrix[nextElemPosition.x][nextElemPosition.y])
			switch nextElem {
			case '-':
				directions = append(directions, Left, Right)
			case '/':
				directions = append(directions, Right)
			case '\\':
				directions = append(directions, Left)
			default:
				directions = append(directions, Up)
			}
		}
	case Down:
		if (s1.Position.x + 1) < len(matrix) {
			nextElemPosition = Point{s1.Position.x + 1, s1.Position.y}
			nextElem = rune(matrix[nextElemPosition.x][nextElemPosition.y])
			switch nextElem {
			case '-':
				directions = append(directions, Left, Right)
			case '/':
				directions = append(directions, Left)
			case '\\':
				directions = append(directions, Right)
			default:
				directions = append(directions, Down)
			}
		}
	}
	for _, dir := range directions {
		newStates = append(newStates, State{
			Direction: dir,
			Position:  nextElemPosition,
			GridElem:  rune(nextElem),
		})
	}
	return newStates
}

func Traverse(startState State, matrix []string, path []State) []State {
	traversalPath := make([]State, 0)
	if newStates := NextState(startState, matrix); newStates != nil {
		for _, state := range newStates {
			if !slices.Contains(path, state) {
				traversalPath = append(traversalPath, state)
				traversalPath = append(traversalPath, Traverse(state, matrix, path)...)
			}
		}
	}
	return traversalPath
}

func TraverseLoop(startState State, matrix []string, path []State) []State {
	remainingStates := make([]State, 0)
	for {
		if len(path) > 0 && slices.Contains(path[:len(path)-1], startState) {
			break
		}
		newStates := NextState(startState, matrix)
		if len(newStates) == 0 {
			break
		} else {
			startState = newStates[0]
			path = append(path, startState)
			if len(newStates) > 1 {
				remainingStates = append(remainingStates, newStates[1:]...)
			}
		}
	}
	if len(remainingStates) != 0 {
		for _, state := range remainingStates {
			if !slices.Contains(path, state) {
				path = append(path, TraverseLoop(state, matrix, path)...)
			}
		}
	}
	return path
}
