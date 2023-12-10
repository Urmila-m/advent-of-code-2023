package main

import (
	"day8/lib"
	"fmt"
)

func main() {
	instructions, network := lib.ParseFile("puzzle_input.txt")
	fmt.Println(lib.FindNumOfSteps(instructions, network, "AAA", "ZZZ"))

}
