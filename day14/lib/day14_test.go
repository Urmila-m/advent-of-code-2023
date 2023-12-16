package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollToNorth(t *testing.T) {
	expectedOp := [][]rune{
		{'O', 'O', 'O', 'O', '.', '#', '.', 'O', '.', '.'},
		{'O', 'O', '.', '.', '#', '.', '.', '.', '.', '#'},
		{'O', 'O', '.', '.', 'O', '#', '#', '.', '.', 'O'},
		{'O', '.', '.', '#', '.', 'O', 'O', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '#', '.', '.', '.', '.', '#', '.', '#'},
		{'.', '.', 'O', '.', '.', '#', '.', 'O', '.', 'O'},
		{'.', '.', 'O', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
	}
	lines := RollToNorth(ParseFile("test_input.txt"))
	assert.Equal(t, expectedOp, lines)
}

func TestRollToSouth(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '#'},
		{'.', '.', '.', 'O', '.', '#', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'O', '.', 'O', '.', '.', '.', '.', 'O', '#', 'O'},
		{'O', '.', '#', '.', '.', 'O', '.', '#', '.', '#'},
		{'O', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'O', 'O', '.', '.', '.', '.', 'O', 'O', '.', '.'},
		{'#', 'O', 'O', '.', '.', '#', '#', '#', '.', '.'},
		{'#', 'O', 'O', '.', 'O', '#', '.', '.', '.', 'O'},
	}
	lines := RollToSouth(ParseFile("test_input.txt"))
	assert.Equal(t, expectedOp, lines)
}

func TestRollToWest(t *testing.T) {
	expectedOp := [][]rune{
		{'O', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'O', 'O', 'O', '.', '#', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '#', '#', '.', '.', '.'},
		{'O', 'O', '.', '#', 'O', 'O', '.', '.', '.', '.'},
		{'O', 'O', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'O', '.', '#', 'O', '.', '.', '.', '#', '.', '#'},
		{'O', '.', '.', '.', '.', '#', 'O', 'O', '.', '.'},
		{'O', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
		{'#', 'O', 'O', '.', '.', '#', '.', '.', '.', '.'},
	}
	assert.Equal(t, expectedOp, RollToWest(ParseFile("test_input.txt")))
}

func TestRollToEast(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '.', 'O', '#', '.', '.', '.', '.'},
		{'.', 'O', 'O', 'O', '#', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '#', '#', '.', '.', '.'},
		{'.', 'O', 'O', '#', '.', '.', '.', '.', 'O', 'O'},
		{'.', '.', '.', '.', '.', '.', 'O', 'O', '#', '.'},
		{'.', 'O', '#', '.', '.', '.', 'O', '#', '.', '#'},
		{'.', '.', '.', '.', 'O', '#', '.', '.', 'O', 'O'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'O'},
		{'#', '.', '.', '.', '.', '#', '#', '#', '.', '.'},
		{'#', '.', '.', 'O', 'O', '#', '.', '.', '.', '.'},
	}
	assert.Equal(t, expectedOp, RollToEast(ParseFile("test_input.txt")))
}

func TestRollACycle(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '#', '.', '.', '.', 'O', '#'},
		{'.', '.', '.', 'O', 'O', '#', '#', '.', '.', '.'},
		{'.', 'O', 'O', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', 'O', 'O', 'O', '#', '.'},
		{'.', 'O', '#', '.', '.', '.', 'O', '#', '.', '#'},
		{'.', '.', '.', '.', 'O', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'O', 'O', 'O', 'O'},
		{'#', '.', '.', '.', 'O', '#', '#', '#', '.', '.'},
		{'#', '.', '.', 'O', 'O', '#', '.', '.', '.', '.'},
	}
	after1Cycle := RollACycle(ParseFile("test_input.txt"))
	DisplayMatrix(after1Cycle)
	assert.Equal(t, expectedOp, after1Cycle)
}

func TestRollNCycle(t *testing.T) {
	expectedAfter2Cycle := [][]rune{
		{'.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '#', '.', '.', '.', 'O', '#'},
		{'.', '.', '.', '.', '.', '#', '#', '.', '.', '.'},
		{'.', '.', 'O', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', 'O', 'O', 'O', '#', '.'},
		{'.', 'O', '#', '.', '.', '.', 'O', '#', '.', '#'},
		{'.', '.', '.', '.', 'O', '#', '.', '.', '.', 'O'},
		{'.', '.', '.', '.', '.', '.', '.', 'O', 'O', 'O'},
		{'#', '.', '.', 'O', 'O', '#', '#', '#', '.', '.'},
		{'#', '.', 'O', 'O', 'O', '#', '.', '.', '.', 'O'},
	}
	expectedAfter3Cycle := [][]rune{
		{'.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '#', '.', '.', '.', 'O', '#'},
		{'.', '.', '.', '.', '.', '#', '#', '.', '.', '.'},
		{'.', '.', 'O', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', 'O', 'O', 'O', '#', '.'},
		{'.', 'O', '#', '.', '.', '.', 'O', '#', '.', '#'},
		{'.', '.', '.', '.', 'O', '#', '.', '.', '.', 'O'},
		{'.', '.', '.', '.', '.', '.', '.', 'O', 'O', 'O'},
		{'#', '.', '.', '.', 'O', '#', '#', '#', '.', 'O'},
		{'#', '.', 'O', 'O', 'O', '#', '.', '.', '.', 'O'},
	}
	matrix := ParseFile("test_input.txt")
	after1Cycle := RollACycle(matrix)
	after2Cycle := RollACycle(after1Cycle)
	DisplayMatrix(after2Cycle)
	assert.Equal(t, expectedAfter2Cycle, after2Cycle)
	fmt.Println("----------")
	after3Cycle := RollACycle(after2Cycle)
	DisplayMatrix(after3Cycle)
	assert.Equal(t, expectedAfter3Cycle, after1Cycle)
}

func TestFindCycleRepeatPosition(t *testing.T) {
	repeatsAfter, repeatStartPoint, allNonRepeatMatrices := FindCycleRepeatPosition(ParseFile("test_input.txt"))
	for _, matrix := range allNonRepeatMatrices {
		DisplayMatrix(matrix)
		fmt.Println("----------")
	}
	DisplayMatrix(RollACycle(allNonRepeatMatrices[len(allNonRepeatMatrices)-1]))
	assert.Equal(t, 10, repeatsAfter)
	assert.Equal(t, 3, repeatStartPoint)

	for i, matrix := range allNonRepeatMatrices {
		fmt.Println(i, CalcNorthLoad(matrix))
	}
}

func TestCalcNorthLoad(t *testing.T) {
	assert.Equal(t, 136, CalcNorthLoad(RollToNorth(ParseFile("test_input.txt"))))
}

func TestFindNorthLoadAfterSpin(t *testing.T) {
	assert.Equal(t, 64, FindNorthLoadAfterSpin(1000000000, "test_input.txt"))
}
