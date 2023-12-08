package main

import (
	"day7/lib"
	"fmt"
)

func main() {
	fmt.Println(CalcTotalWinnings("puzzle_input.txt"))
}

func CalcTotalWinnings(filePath string) int {
	totalWinnings := 0
	for i, hand := range lib.SortHands(lib.ParseFile(filePath, false), false) {
		totalWinnings += hand.Bid * (i + 1)
	}
	return totalWinnings
}
