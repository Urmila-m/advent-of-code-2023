package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	calibrationSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibrationSum += findCalibrationValue(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(calibrationSum)
}

type DigitIndex struct {
	Digit string
	Index int
}

func findAllDigitIndices(line string) []DigitIndex {
	digits := [18]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var digitIndices []DigitIndex
	for _, digit := range digits {
		digitIndices = append(digitIndices, findDigitIndices(line, digit, 0)...)
	}
	return digitIndices
}

func findDigitIndices(line string, digit string, offset int) []DigitIndex {
	var digitIndices []DigitIndex
	subLine := line[offset:]
	if present := strings.Index(subLine, digit); present != -1 {
		digitIndices = append(digitIndices, DigitIndex{digit, offset + present})
		if strings.Count(subLine, digit) > 1 {
			offset += present + len(digit)
			digitIndices = append(digitIndices, findDigitIndices(line, digit, offset)...)
		}
	}
	return digitIndices
}

func sortByValue(digitIndices []DigitIndex) []DigitIndex {
	sort.Slice(digitIndices, func(i int, j int) bool {
		return digitIndices[i].Index < digitIndices[j].Index
	})
	return digitIndices
}

func findCalibrationValue(line string) int {
	digitIndices := findAllDigitIndices(line)
	sortedDigitIndices := sortByValue(digitIndices)
	firstDigit := mapToDigit(sortedDigitIndices[0].Digit)
	lastDigit := mapToDigit(sortedDigitIndices[len(sortedDigitIndices)-1].Digit)
	calibrationValueStr := fmt.Sprint(firstDigit, lastDigit)
	calibrationValue, err := strconv.Atoi(calibrationValueStr)

	if err != nil {
		log.Fatal(err)
	}
	return calibrationValue
}

func mapToDigit(number string) string {
	if len(number) == 1 {
		return number
	}
	var digitWordMap = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	digit, ok := digitWordMap[number]

	if !ok {
		log.Fatal("Not a number")
	}
	return digit
}
