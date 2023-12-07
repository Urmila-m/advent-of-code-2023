package main

import (
	"day6/lib"
	"fmt"
)

func main() {
	fmt.Println(FindPuzzle2Output("puzzle_input.txt"))

}

func FindPuzzle2Output(fileName string) int {
	totalTime, distanceRecord := lib.ParseFile2(fileName)
	return lib.FindNumWaysToBeat(totalTime, distanceRecord)
}
