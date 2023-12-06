package main

import (
	"day5/lib"
	"fmt"
)

func main() {
	seed, location := FindLowestLocation("puzzle_input.txt")
	fmt.Printf("Location: %d\nSeed:%d\n", location, seed)
}

func FindLowestLocation(filePath string) (seed, location int) {
	initialSeeds, allCategoryMaps := lib.ParseFile(filePath)

	seedToLocation := make(map[int]int)
	for _, seed := range initialSeeds {
		seedToLocation[seed] = lib.FindLocation(seed, allCategoryMaps)
	}

	minLocation := -1
	minSeed := -1
	for seed, location := range seedToLocation {
		if minLocation == -1 || location < minLocation {
			minSeed = seed
			minLocation = location
		}
	}
	return minSeed, minLocation
}
