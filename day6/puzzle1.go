package main

import (
	"day6/lib"
	"fmt"
)

func main() {
	fmt.Println(FindPuzzle1Output("puzzle_input.txt"))
}

func FindPuzzle1Output(filepath string) int {
	output := -1
	for totalTime, distanceRecord := range lib.MapTimeDistance(lib.ParseFile(filepath)) {
		numWaysToBeat := lib.FindNumWaysToBeat(totalTime, distanceRecord)
		if numWaysToBeat == 0 {
			return 0
		}
		if output == -1 {
			output = numWaysToBeat
		} else {
			output *= numWaysToBeat
		}
	}
	return output
}
