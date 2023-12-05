package lib

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

func FindAllAdjNums(allLines []string) []int {
	overallAdjNums := make([]int, 0)
	for i, line := range allLines {
		for j, character := range []rune(line) {
			if isSpecialCharacter := IsSpecialCharacter(character); isSpecialCharacter {
				overallAdjNums = append(overallAdjNums, FindAdjNumsPerChar(i, j, allLines)...)
			}
		}
	}
	return overallAdjNums
}
func ParseLinesFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var allLines []string
	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return allLines
}

func IsSpecialCharacter(c rune) bool {
	return !(unicode.IsDigit(c) || string(c) == ".")
}

func FindAdjNumsPerChar(lineNum int, specialCharPos int, lines []string) []int {
	allAdjNumbers := FindAdjNumsInSameLine(specialCharPos, lines[lineNum])

	if lineNum > 0 {
		allAdjNumbers = append(allAdjNumbers, FindAdjNumsInAdjLine(specialCharPos, lines[lineNum-1])...)
	}
	if lineNum < (len(lines) - 1) {
		allAdjNumbers = append(allAdjNumbers, FindAdjNumsInAdjLine(specialCharPos, lines[lineNum+1])...)
	}

	return allAdjNumbers
}

func FindImmLeftNum(specialCharPos int, line string) (leftNumber int, present bool) {
	var number string = ""
	for i := specialCharPos - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			number = fmt.Sprintf("%c%s", rune(line[i]), number)
		} else {
			break
		}
	}
	if number != "" {
		leftNumber, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		return leftNumber, true
	}
	return 0, false
}

func FindImmRightNum(specialCharPos int, line string) (rightNumber string, present bool) {
	var number string = ""
	for i := specialCharPos + 1; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			number = fmt.Sprintf("%s%c", number, rune(line[i]))
		} else {
			break
		}
	}
	if number != "" {
		return number, true
	}
	return "", false
}

func FindAdjNumsInSameLine(specialCharPos int, line string) []int {
	adjNums := make([]int, 0)

	if leftNum, present := FindImmLeftNum(specialCharPos, line); present {
		adjNums = append(adjNums, leftNum)
	}

	if rightNum, present := FindImmRightNum(specialCharPos, line); present {
		rightNum, err := strconv.Atoi(rightNum)
		if err != nil {
			log.Fatal(err)
		}
		adjNums = append(adjNums, rightNum)
	}

	return adjNums
}

func FindAdjNumsInAdjLine(specialCharPos int, adjLine string) []int {
	adjNums := make([]int, 0)

	lowerBound := int32(math.Max(0, float64(specialCharPos-1)))
	upperBound := int32(math.Min(float64(len(adjLine)-1), float64(specialCharPos+1)))
	for i := lowerBound; i <= upperBound; i++ {
		currChar := rune(adjLine[i])

		if unicode.IsDigit(currChar) {
			leftNum, _ := FindImmLeftNum(int(i), adjLine)
			adjNumStr := fmt.Sprintf("%d%c", leftNum, currChar)

			rightNum, rPresent := FindImmRightNum(int(i), adjLine)

			if rPresent {
				adjNumStr = fmt.Sprintf("%s%s", adjNumStr, rightNum)
			}

			adjNum, err := strconv.Atoi(adjNumStr)
			if err != nil {
				log.Fatal(err)
			}
			adjNums = append(adjNums, adjNum)

			if rPresent {
				break
			}
		}
	}
	return adjNums

}

func FindGearRatio(lineNum int, specialCharPos int, lines []string) (gearRatio int, isGear bool) {
	allAdjNums := FindAdjNumsPerChar(lineNum, specialCharPos, lines)
	if len(allAdjNums) == 2 {
		gearRatio = allAdjNums[0] * allAdjNums[1]
		return gearRatio, true
	}
	return 0, false
}
