package main

import (
	"day3/lib"
	"fmt"
)

func main() {
	allLines := lib.ParseLinesFromFile("puzzle_input.txt")
	allAdjNum := lib.FindAllAdjNums(allLines)
	//fmt.Println(allAdjNum)
	// fmt.Println(len(allAdjNum))
	partNumberSum := 0
	for _, num := range allAdjNum {
		partNumberSum += num
	}
	fmt.Println(partNumberSum)
}
