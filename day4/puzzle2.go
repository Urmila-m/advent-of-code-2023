package main

import (
	"day4/lib"
	"fmt"
)

func main() {
	fmt.Println(lib.CalcTotalScratchCards(lib.CalcAllCardsCount("puzzle_input.txt")))
}
