package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CategoryMap struct {
	Src     string
	Dst     string
	Mapping []SrcDstMapping
}

type SrcDstMapping struct {
	SrcStart    int
	DstStart    int
	RangeLength int
}

func ToCategoryMapping(mapChunk []string) CategoryMap {
	mapKeys := strings.Split(strings.Split(mapChunk[0], " ")[0], "-to-")
	src := mapKeys[0]
	dest := mapKeys[1]

	allMappings := make([]SrcDstMapping, 0)
	for _, line := range mapChunk[1:] {
		numsInLine := strings.Split(line, " ")

		dstStart, err := strconv.Atoi(numsInLine[0])
		if err != nil {
			log.Fatal(err)
		}

		srcStart, err := strconv.Atoi(numsInLine[1])
		if err != nil {
			log.Fatal(err)
		}
		rangeLength, err := strconv.Atoi(numsInLine[2])
		if err != nil {
			log.Fatal(err)
		}
		allMappings = append(allMappings, SrcDstMapping{SrcStart: srcStart, DstStart: dstStart, RangeLength: rangeLength})
	}

	return CategoryMap{
		Src:     src,
		Dst:     dest,
		Mapping: allMappings,
	}
}

func FindInitialSeeds(line string) []int {
	allSeeds := make([]int, 0)
	for _, seedStr := range strings.Split(strings.Split(line, ": ")[1], " ") {
		if seed, convertErr := strconv.Atoi(seedStr); convertErr != nil {
			log.Fatal(convertErr)
		} else {
			allSeeds = append(allSeeds, seed)
		}
	}
	return allSeeds
}

func ParseFile(filePath string) ([]int, []CategoryMap) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if err != nil {
			log.Fatal(closeErr)
		}
	}(file)

	var initialSeeds []int
	scanner := bufio.NewScanner(file)

	// Parse 1st line to find seed input
	if scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds:") {
			initialSeeds = FindInitialSeeds(line)
		}
	}

	//Scan the next empty line
	if scanner.Scan() {
		scanner.Text()
	}

	mapChunks := make([][]string, 0)
	mapChunk := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mapChunks = append(mapChunks, mapChunk)
			mapChunk = make([]string, 0)
		} else {
			mapChunk = append(mapChunk, line)
		}
	}
	mapChunks = append(mapChunks, mapChunk)
	allCategoryMaps := make([]CategoryMap, 0)
	for _, mapChunk := range mapChunks {
		allCategoryMaps = append(allCategoryMaps, ToCategoryMapping(mapChunk))
	}
	return initialSeeds, allCategoryMaps
}

func MapCategory(srcValue int, categoryMap CategoryMap) int {
	for _, mapping := range categoryMap.Mapping {
		if mapping.SrcStart <= srcValue && srcValue < (mapping.SrcStart+mapping.RangeLength) {
			return mapping.DstStart + (srcValue - mapping.SrcStart)
		}
	}
	return srcValue
}

func FindLocation(seed int, allCategoryMap []CategoryMap) int {
	mappedValue := seed
	for _, categoryMap := range allCategoryMap {
		mappedValue = MapCategory(mappedValue, categoryMap)
	}
	return mappedValue
}
