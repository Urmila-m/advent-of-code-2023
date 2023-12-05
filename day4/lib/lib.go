package lib

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func FindCardPoints(winNums, ownNums []int) int {
	points := 0
	for _, num := range ownNums {
		if slices.Contains(winNums, num) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func ParseLine(line string) (winNums []int, ownNums []int) {
	allNumsArr := strings.Split(strings.Split(line, ":")[1], "|")

	winNums = ParseNumsArrFromStr(allNumsArr[0])
	ownNums = ParseNumsArrFromStr(allNumsArr[1])
	return winNums, ownNums
}

func ParseNumsArrFromStr(str string) []int {
	numsArr := make([]int, 0)
	for _, numStr := range strings.Split(str, " ") {
		if num, err := strconv.Atoi(numStr); err == nil {
			numsArr = append(numsArr, num)
		}
	}
	return numsArr
}

func FindMatchingCardNum(winNums, ownNums []int) int {
	matchingCardNum := 0
	for _, num := range ownNums {
		if slices.Contains(winNums, num) {
			matchingCardNum += 1
		}
	}
	return matchingCardNum
}

func FindAllWonCards(cardNum int, line string) []int {
	allWonCards := make([]int, 0)

	for i := 1; i <= FindMatchingCardNum(ParseLine(line)); i++ {
		allWonCards = append(allWonCards, cardNum+i)
	}
	return allWonCards
}

func AddCardToMap(totalScratchCards *map[int]int, cardNum int, count int) {
	_, exists := (*totalScratchCards)[cardNum]
	if !exists {
		(*totalScratchCards)[cardNum] = count
	} else {
		(*totalScratchCards)[cardNum] += count
	}
}

func CalcAllCardsCount(filePath string) map[int]int {
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

	allCardsCount := make(map[int]int)
	scanner := bufio.NewScanner(file)

	cardNum := 0
	for scanner.Scan() {
		cardNum += 1
		AddCardToMap(&allCardsCount, cardNum, 1)
		line := scanner.Text()
		if allWonCards := FindAllWonCards(cardNum, line); allWonCards != nil {
			for _, wonCard := range allWonCards {
				AddCardToMap(&allCardsCount, wonCard, allCardsCount[cardNum])
			}
		}
	}
	return allCardsCount
}

func CalcTotalScratchCards(allCardsCount map[int]int) int {
	totalScratchCards := 0
	for _, count := range allCardsCount {
		totalScratchCards += count
	}
	return totalScratchCards
}
