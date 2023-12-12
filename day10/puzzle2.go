package main

import (
	"day10/lib"
	"fmt"
)

func main() {
	borderPipes := lib.FindMainLoop("puzzle_input.txt")

	borderPoints := make([]lib.Point, 0)
	for _, pipe := range borderPipes[:len(borderPipes)-1] {
		borderPoints = append(borderPoints, pipe.Position)
	}
	fmt.Println(lib.FindNumOfInsidePoints(lib.FindAreaUsingShoeLace(borderPoints), len(borderPoints)))
}
