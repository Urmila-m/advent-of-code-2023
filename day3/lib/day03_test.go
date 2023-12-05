package lib

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestFindLeftNumber(t *testing.T) {
	leftNumber, present := FindImmLeftNum(3, "617*......")
	if present {
		assert.Equal(t, 617, leftNumber, "They should be equal")
	}
}

func TestFindAdjNumsInAdjLine(t *testing.T) {
	/*
		6.749.
		...*..
		.699..
	*/
	assert.Equal(t, []int{749}, FindAdjNumsInAdjLine(3, "6.749."), "They should be equal")
	assert.Equal(t, []int{699}, FindAdjNumsInAdjLine(3, ".699.."))

}

func TestFindAllAdjNums(t *testing.T) {
	allLines := ParseLinesFromFile("../test_input.txt")
	allAdjNum := FindAllAdjNums(allLines)
	sort.Slice(allAdjNum, func(i int, j int) bool {
		return allAdjNum[i] < allAdjNum[j]
	})
	expectedOp := []int{467, 35, 633, 617, 592, 755, 664, 598}
	sort.Slice(expectedOp, func(i int, j int) bool {
		return expectedOp[i] < expectedOp[j]
	})
	assert.Equal(t, allAdjNum, expectedOp)
	partNumberSum := 0
	for _, num := range allAdjNum {
		partNumberSum += num
	}
	assert.Equal(t, partNumberSum, 4361)
}
