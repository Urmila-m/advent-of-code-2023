package main

import (
	"day14/lib"
	"fmt"
)

func puzzle_one() {
	fmt.Println(lib.CalcNorthLoad(lib.RollToNorth(lib.ParseFile("puzzle_input.txt"))))
}

func puzzle_two() {
	fmt.Println(lib.FindNorthLoadAfterSpin(1000000000, "puzzle_input.txt"))
}

func main() {
	puzzle_one()
	puzzle_two()
}
