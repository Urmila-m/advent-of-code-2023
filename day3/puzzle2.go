package main

import (
	"day3/lib"
	"fmt"
)

func main() {
	allLines := lib.ParseLinesFromFile("puzzle_input.txt")

	gearRatioSum := 0
	for i, line := range allLines {
		for j, character := range []rune(line) {
			if string(character) == "*" {
				if gearRatio, isGear := lib.FindGearRatio(i, j, allLines); isGear {
					gearRatioSum += gearRatio
				}
			}
		}
	}
	fmt.Println(gearRatioSum)
}
