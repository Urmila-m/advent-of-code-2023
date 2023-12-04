package main

import (
	"advent-of-code-2023/lib"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("test_input.txt")
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	fewestCubeSetPowerSum := 0

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := lib.ParseGameFromLine(line)
		fewestCubeSet := lib.FindFewestCubesPossible(game)
		fewestCubeSetPower := lib.FindCubeSetPower(fewestCubeSet)
		fewestCubeSetPowerSum += fewestCubeSetPower
		fmt.Printf("%v \n %v \n\n", game, fewestCubeSet)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fewestCubeSetPowerSum)
}
