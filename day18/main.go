package main

import (
	d10 "day10/lib"
	"day18/lib"
	"fmt"
)

func puzzleOne() {
	allBorderPoints := make([]d10.Point, 0)
	for _, cube := range lib.FindTrench(lib.ParseFile("puzzle_input.txt")) {
		allBorderPoints = append(allBorderPoints, cube.Position)
	}
	fmt.Println(lib.CalcCubeHoldCapacity(allBorderPoints))
}

func puzzleTwo() {
	fmt.Println(lib.CalcCubeHoldCapacity(lib.FindTrench2(lib.ParseFile("puzzle_input.txt"))))
}

func main() {
	puzzleOne()
	puzzleTwo()
}
