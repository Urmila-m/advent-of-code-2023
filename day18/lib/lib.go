package lib

import (
	"bufio"
	d10 "day10/lib"
	d16 "day16/lib"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	Position d10.Point
	Color    string
}

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
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func FindTrench(lines []string) []Cube {
	directionMap := map[string]d16.Direction{
		"R": d16.Right,
		"L": d16.Left,
		"U": d16.Up,
		"D": d16.Down,
	}
	startingPoint := d10.Point{X: 0, Y: 0}
	allCubes := make([]Cube, 0)

	for _, line := range lines {
		splitBySpace := strings.Fields(line)
		direction := splitBySpace[0]
		steps, convertErr := strconv.Atoi(splitBySpace[1])
		if convertErr != nil {
			log.Fatal(convertErr)
		}
		colorCode := strings.Trim(splitBySpace[2], "()")
		nextPoints := Next(directionMap[direction], steps, startingPoint)
		for _, p := range nextPoints {
			allCubes = append(allCubes, Cube{
				Position: p,
				Color:    colorCode,
			})
		}
		startingPoint = nextPoints[len(nextPoints)-1]
	}
	return allCubes
}

func Next(direction d16.Direction, numOfSteps int, currentPos d10.Point) []d10.Point {
	directionMap := map[d16.Direction]d10.Point{
		d16.Right: {X: 0, Y: 1},
		d16.Left:  {X: 0, Y: -1},
		d16.Up:    {X: -1, Y: 0},
		d16.Down:  {X: 1, Y: 0},
	}
	nextPoints := make([]d10.Point, 0)
	for i := 1; i <= numOfSteps; i++ {
		nextPoints = append(nextPoints, d10.Point{X: currentPos.X + directionMap[direction].X*i, Y: currentPos.Y + directionMap[direction].Y*i})
	}
	return nextPoints
}

func CalcCubeHoldCapacity(mainLoop []d10.Point) int {
	return d10.FindNumOfInsidePoints(d10.FindAreaUsingShoeLace(mainLoop), len(mainLoop)) + len(mainLoop)
}

func CalcDistanceDirectionFromColor(colorCode string) (d16.Direction, int) {
	directionMap := map[string]d16.Direction{
		"0": d16.Right,
		"1": d16.Down,
		"2": d16.Left,
		"3": d16.Up,
	}
	numOfSteps, convertErr := strconv.ParseInt(colorCode[:5], 16, 0)
	if convertErr != nil {
		log.Fatal(convertErr)
	}
	direction := directionMap[string(colorCode[5])]
	return direction, int(numOfSteps)
}

func FindTrench2(lines []string) []d10.Point {
	startingPoint := d10.Point{X: 0, Y: 0}
	allPoints := make([]d10.Point, 0)

	for _, line := range lines {
		splitBySpace := strings.Fields(line)
		colorCode := strings.Trim(splitBySpace[2], "#()")
		direction, steps := CalcDistanceDirectionFromColor(colorCode)
		allPoints = append(allPoints, Next(direction, steps, startingPoint)...)
		startingPoint = allPoints[len(allPoints)-1]
	}
	return allPoints
}
