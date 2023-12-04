package main

import (
	"advent-of-code-2023/lib"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	referenceCubeSet := lib.CubeSet{Red: 12, Green: 13, Blue: 14}
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

	gameIdSum := 0

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game := lib.ParseGameFromLine(line)
		if lib.IsGameValid(game, referenceCubeSet) {
			gameIdSum += game.GameId
			// fmt.Printf("%s \n %v \n\n", line, game)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(gameIdSum)
}
