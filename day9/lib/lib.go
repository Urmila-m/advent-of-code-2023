package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func PredictNextValue(numbers []int) int {
	if AreAllElementsSame(numbers) {
		return numbers[0]
	}
	diff := make([]int, 0)

	for i := 0; i < len(numbers)-1; i++ {
		diff = append(diff, numbers[i+1]-numbers[i])
	}
	return numbers[len(numbers)-1] + PredictNextValue(diff)
}

func PredictPreviousValue(numbers []int) int {
	if AreAllElementsSame(numbers) {
		return numbers[0]
	}
	diff := make([]int, 0)

	for i := 0; i < len(numbers)-1; i++ {
		diff = append(diff, numbers[i+1]-numbers[i])
	}
	return numbers[0] - PredictPreviousValue(diff)
}

func AreAllElementsSame(numbers []int) bool {
	num := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != num {
			return false
		}
	}
	return true
}

func ParseFile(filePath string) [][]int {
	file, fError := os.Open(filePath)
	if fError != nil {
		log.Fatal(fError)
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	lines := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		mLine := make([]int, 0)
		for _, numStr := range strings.Fields(line) {
			num, convertErr := strconv.Atoi(numStr)
			if convertErr != nil {
				log.Fatal(convertErr)
			}
			mLine = append(mLine, num)
		}
		lines = append(lines, mLine)
	}
	return lines
}

func SumExtrapolatedValues(filePath string, extrapolateType string) int {
	sum := 0
	for _, line := range ParseFile(filePath) {
		if extrapolateType == "next" {
			sum += PredictNextValue(line)
		} else if extrapolateType == "previous" {
			sum += PredictPreviousValue(line)
		} else {
			log.Fatal("Extrapolation Type incorrect!")
		}
	}
	return sum
}
