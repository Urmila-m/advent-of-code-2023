package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLine(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	winNums, ownNums := ParseLine(line)
	assert.Equal(t, []int{41, 48, 83, 86, 17}, winNums)
	assert.Equal(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, ownNums)
}

func TestFindCardPoints(t *testing.T) {
	winNums := []int{41, 48, 83, 86, 17}
	ownNums := []int{83, 86, 6, 31, 17, 9, 48, 53}

	assert.Equal(t, 8, FindCardPoints(winNums, ownNums))
}

func TestFindMatchingCardNum(t *testing.T) {
	winNums := []int{41, 48, 83, 86, 17}
	ownNums := []int{83, 86, 6, 31, 17, 9, 48, 53}
	assert.Equal(t, 4, FindMatchingCardNum(winNums, ownNums))
}

func TestCalcAllWonCards(t *testing.T) {
	line := "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	assert.Equal(t, []int{3, 4}, FindAllWonCards(2, line))
}

func TestCalcAllCardsCount(t *testing.T) {
	expectedOp := map[int]int{
		1: 1,
		2: 2,
		3: 4,
		4: 8,
		5: 14,
		6: 1,
	}
	assert.Equal(t, expectedOp, CalcAllCardsCount("test_input.txt"))
}

func TestCalcTotalScratchCards(t *testing.T) {
	input := map[int]int{
		1: 1,
		2: 2,
		3: 4,
		4: 8,
		5: 14,
		6: 1,
	}
	assert.Equal(t, 30, CalcTotalScratchCards(input))
}
