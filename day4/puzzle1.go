package main

import (
	"bufio"
	"day4/lib"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(FindOverallPoints("puzzle_input.txt"))
}

func FindOverallPoints(filePath string) int {
	file, fErr := os.Open(filePath)
	if fErr != nil {
		log.Fatal(fErr)
	}

	defer func(f *os.File) {
		closeErr := f.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	overallPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		overallPoints += lib.FindCardPoints(lib.ParseLine(line))
	}
	return overallPoints
}
