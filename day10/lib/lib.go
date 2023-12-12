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
	x int
	y int
}

type Pipe struct {
	Letter   rune
	Position Point
}

func (pipe Pipe) FindAdjacentPipes(lines []string) map[Point]rune {

	adjPipes := make(map[Point]rune)
	switch pipe.Letter {
	case 'L':
		pos1 := Point{pipe.Position.x - 1, pipe.Position.y}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x, pipe.Position.y + 1}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
	case 'J':
		pos1 := Point{pipe.Position.x - 1, pipe.Position.y}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x, pipe.Position.y - 1}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
	case '7':
		pos1 := Point{pipe.Position.x + 1, pipe.Position.y}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x, pipe.Position.y - 1}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
	case 'F':
		pos1 := Point{pipe.Position.x + 1, pipe.Position.y}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x, pipe.Position.y + 1}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
	case 'S':
		allPositionValidChars := map[Point][]rune{
			Point{x: pipe.Position.x + 1, y: pipe.Position.y}: []rune{'|', 'J', 'L'},
			Point{x: pipe.Position.x, y: pipe.Position.y - 1}: []rune{'-', 'F', 'L'},
			Point{x: pipe.Position.x, y: pipe.Position.y + 1}: []rune{'-', '7', 'J'},
			Point{x: pipe.Position.x - 1, y: pipe.Position.y}: []rune{'|', '7', 'F'},
		}

		for position, validChars := range allPositionValidChars {
			if position.x >= 0 && position.x < len(lines) && position.y >= 0 && position.y < len(lines[0]) {
				actualChar := rune(lines[position.x][position.y])
				if slices.Contains(validChars, actualChar) {
					adjPipes[position] = actualChar
				}
			}
		}
	case '-':
		pos1 := Point{pipe.Position.x, pipe.Position.y - 1}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x, pipe.Position.y + 1}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
	case '|':
		pos1 := Point{pipe.Position.x - 1, pipe.Position.y}
		adjPipes[pos1] = rune(lines[pos1.x][pos1.y])

		pos2 := Point{pipe.Position.x + 1, pipe.Position.y}
		adjPipes[pos2] = rune(lines[pos2.x][pos2.y])
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
		positiveArea += coords.x * coordsArr[i+1].y
		negativeArea += coords.y * coordsArr[i+1].x
	}
	positiveArea += coordsArr[len(coordsArr)-1].x * coordsArr[0].y
	negativeArea += coordsArr[len(coordsArr)-1].y * coordsArr[0].x
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
