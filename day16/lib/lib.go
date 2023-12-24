package lib

import (
	"bufio"
	"fmt"
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
	X int
	Y int
}

type State struct {
	GridElem  rune
	Position  Point
	Direction Direction
}

func (s State) Display() {
	fmt.Printf("{'%c', {%d, %d}, %s}\n", s.GridElem, s.Position.X, s.Position.Y, s.Direction)
}

func NextState(s1 State, matrix []string) []State {
	newStates := make([]State, 0)
	nextElemPosition := Point{-1, -1}
	directions := make([]Direction, 0)
	var nextElem rune
	switch s1.Direction {
	case Right:
		if (s1.Position.Y + 1) < len(matrix[0]) {
			nextElemPosition = Point{s1.Position.X, s1.Position.Y + 1}
			nextElem = rune(matrix[nextElemPosition.X][nextElemPosition.Y])
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
		if (s1.Position.Y - 1) >= 0 {
			nextElemPosition = Point{s1.Position.X, s1.Position.Y - 1}
			nextElem = rune(matrix[nextElemPosition.X][nextElemPosition.Y])
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
		if (s1.Position.X - 1) >= 0 {
			nextElemPosition = Point{s1.Position.X - 1, s1.Position.Y}
			nextElem = rune(matrix[nextElemPosition.X][nextElemPosition.Y])
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
		if (s1.Position.X + 1) < len(matrix) {
			nextElemPosition = Point{s1.Position.X + 1, s1.Position.Y}
			nextElem = rune(matrix[nextElemPosition.X][nextElemPosition.Y])
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
	pathLen := len(path)
	if newStates := NextState(startState, matrix); len(newStates) != 0 {
		for _, state := range newStates {
			if !slices.Contains(path, state) {
				path = append(path, state)
				path = append(path, Traverse(state, matrix, path)...)
			}
		}
	}
	return path[pathLen:]
}

func FindUniqGridElems(path []State) map[Point]struct{} {
	uniq := make(map[Point]struct{})
	for _, s := range path {
		if _, ok := uniq[s.Position]; !ok {
			uniq[s.Position] = struct{}{}
		}
	}
	return uniq
}

func FindInitialStates(matrix []string) []State {
	initialStates := make([]State, 0)
	for i := 0; i < len(matrix); i++ {
		initialStates = append(initialStates, State{
			GridElem:  '.',
			Position:  Point{i, -1},
			Direction: Right,
		})
		initialStates = append(initialStates, State{
			GridElem:  '.',
			Position:  Point{i, len(matrix[0])},
			Direction: Left,
		})
	}
	for j := 0; j < len(matrix[0]); j++ {
		initialStates = append(initialStates, State{
			GridElem:  '.',
			Position:  Point{-1, j},
			Direction: Down,
		})
		initialStates = append(initialStates, State{
			GridElem:  '.',
			Position:  Point{len(matrix), j},
			Direction: Up,
		})
	}
	return initialStates
}

func FindMostEfficientConfig(matrix []string) (Point, int) {
	mostEfficientStartPoint := Point{-1, -1}
	highestEnergizedGridNums := -1
	for _, s := range FindInitialStates(matrix) {
		numOfEnergizedGrids := len(FindUniqGridElems(Traverse(s, matrix, make([]State, 0))))
		if (mostEfficientStartPoint.X == -1 && mostEfficientStartPoint.Y == -1) || (highestEnergizedGridNums < numOfEnergizedGrids) {
			mostEfficientStartPoint = s.Position
			highestEnergizedGridNums = numOfEnergizedGrids
		}
	}
	return mostEfficientStartPoint, highestEnergizedGridNums
}
