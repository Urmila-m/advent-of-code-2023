package lib

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func ParseFile(filePath string) []string {
	file, fErr := os.Open(filePath)
	if fErr != nil {
		log.Fatal(fErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

type Point struct {
	X int
	Y int
}

type Pipe struct {
	Letter   rune
	Position Point
}

func (pipe Pipe) FindAdjacentPipes(lines []string) map[Point]rune {

	adjPipes := make(map[Point]rune)
	switch pipe.Letter {
	case 'L':
		pos1 := Point{pipe.Position.X - 1, pipe.Position.Y}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X, pipe.Position.Y + 1}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	case 'J':
		pos1 := Point{pipe.Position.X - 1, pipe.Position.Y}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X, pipe.Position.Y - 1}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	case '7':
		pos1 := Point{pipe.Position.X + 1, pipe.Position.Y}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X, pipe.Position.Y - 1}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	case 'F':
		pos1 := Point{pipe.Position.X + 1, pipe.Position.Y}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X, pipe.Position.Y + 1}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	case 'S':
		allPositionValidChars := map[Point][]rune{
			Point{X: pipe.Position.X + 1, Y: pipe.Position.Y}: []rune{'|', 'J', 'L'},
			Point{X: pipe.Position.X, Y: pipe.Position.Y - 1}: []rune{'-', 'F', 'L'},
			Point{X: pipe.Position.X, Y: pipe.Position.Y + 1}: []rune{'-', '7', 'J'},
			Point{X: pipe.Position.X - 1, Y: pipe.Position.Y}: []rune{'|', '7', 'F'},
		}

		for position, validChars := range allPositionValidChars {
			if position.X >= 0 && position.X < len(lines) && position.Y >= 0 && position.Y < len(lines[0]) {
				actualChar := rune(lines[position.X][position.Y])
				if slices.Contains(validChars, actualChar) {
					adjPipes[position] = actualChar
				}
			}
		}
	case '-':
		pos1 := Point{pipe.Position.X, pipe.Position.Y - 1}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X, pipe.Position.Y + 1}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	case '|':
		pos1 := Point{pipe.Position.X - 1, pipe.Position.Y}
		adjPipes[pos1] = rune(lines[pos1.X][pos1.Y])

		pos2 := Point{pipe.Position.X + 1, pipe.Position.Y}
		adjPipes[pos2] = rune(lines[pos2.X][pos2.Y])
	}
	return adjPipes
}

func FindMainLoop(filePath string) []Pipe {
	mainLoop := make([]Pipe, 0)
	startingPoint := Point{-1, -1}
	lines := ParseFile(filePath)
	for i, line := range lines {
		if yPos := strings.Index(line, "S"); yPos != -1 {
			startingPoint = Point{i, yPos}
			mainLoop = append(mainLoop, Pipe{Position: startingPoint, Letter: 'S'})
			break
		}
	}
	currPipe := mainLoop[0]
	var prevPipe Pipe
	isMainLoopComplete := false
	for !isMainLoopComplete {
		for nextPosition, nextPipeLetter := range currPipe.FindAdjacentPipes(lines) {
			if prevPipe == (Pipe{}) || nextPosition != prevPipe.Position {
				prevPipe = currPipe
				currPipe = Pipe{Position: nextPosition, Letter: nextPipeLetter}
				mainLoop = append(mainLoop, currPipe)
				if nextPipeLetter == 'S' {
					isMainLoopComplete = true
				}
				break
			}
		}
	}
	return mainLoop
}

func FindNumOfStepsForFarthestPoint(filePath string) int {
	mainLoop := FindMainLoop(filePath)
	return int(math.Floor(float64(len(mainLoop)-1) / 2))
}

func FindAreaUsingShoeLace(coordsArr []Point) float64 {
	positiveArea := 0
	negativeArea := 0
	for i, coords := range coordsArr[:len(coordsArr)-1] {
		positiveArea += coords.X * coordsArr[i+1].Y
		negativeArea += coords.Y * coordsArr[i+1].X
	}
	positiveArea += coordsArr[len(coordsArr)-1].X * coordsArr[0].Y
	negativeArea += coordsArr[len(coordsArr)-1].Y * coordsArr[0].X
	netArea := math.Abs(float64(positiveArea-negativeArea)) / 2
	return netArea
}

// FindNumOfInsidePoints Uses Pick's Theorem
func FindNumOfInsidePoints(area float64, numOfBorderPoints int) int {
	/*
		Area = NumOfPointsStrictlyLyingInsideLoop + (NumberOfPointsLyingOnBorder/2) - 1
	*/
	return int(area + 1 - float64(numOfBorderPoints/2))
}
