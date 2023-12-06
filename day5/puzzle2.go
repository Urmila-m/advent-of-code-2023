package main

import (
	"day5/lib"
	"fmt"
)

func main() {
	seed, location := FindLowestLocationFromSeedRange("puzzle_input.txt")
	fmt.Printf("Location: %d\nSeed: %d\n", location, seed)
}

func FindLowestLocationFromSeedRange(filePath string) (seed, location int) {
	seedNums, allCategoryMaps := lib.ParseFile(filePath)
	minLocation := -1
	minSeed := -1
	for i := 0; i < len(seedNums); i += 2 {
		for seed := seedNums[i]; seed < (seedNums[i] + seedNums[i+1]); seed++ {
			location := lib.FindLocation(seed, allCategoryMaps)
			if minLocation == -1 || location < minLocation {
				minSeed = seed
				minLocation = location
			}
		}
	}
	return minSeed, minLocation
}
