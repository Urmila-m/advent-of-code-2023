package main

import (
	"day8/lib"
	"fmt"
)

func main() {
	instructions, network := lib.ParseFile("puzzle_input.txt")
	allPathNoOfSteps := make([]int, 0)
	for _, path := range lib.FindAllPaths(instructions, network) {
		allPathNoOfSteps = append(allPathNoOfSteps, path.NoOfSteps)
	}

	/* test `TestEqualityOfNoOfSteps` in `lib/day8_test.go` passes.
	So, the LCM of number of steps of all the paths should give the total number of steps
	required to traverse all the paths simultaneously and have `Z` at last in all of them.
	*/
	fmt.Println(lib.FindLCM(allPathNoOfSteps))
}
