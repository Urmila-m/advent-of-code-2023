package main

import (
	"day15/lib"
	"fmt"
)

func puzzleOne() {
	fmt.Println(lib.FindHashSum(lib.ParseFile("puzzle_input.txt")))
}

func puzzleTwo() {
	fmt.Println(lib.SumAllFocusingPower(lib.FindResultingConfiguration(lib.ParseFile("puzzle_input.txt"))))
}

func main() {
	puzzleOne()
	puzzleTwo()
}
