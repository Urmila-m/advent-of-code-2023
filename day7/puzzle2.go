package main

import (
	"day7/lib"
	"fmt"
)

func main() {
	fmt.Println(CalcTotalWinningsWJoker("puzzle_input.txt"))
}

func CalcTotalWinningsWJoker(filePath string) int {
	totalWinnings := 0
	for i, hand := range lib.SortHands(lib.ParseFile(filePath, true), true) {
		totalWinnings += hand.Bid * (i + 1)
	}
	return totalWinnings
}
