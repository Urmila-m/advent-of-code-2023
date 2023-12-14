package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}
	assert.Equal(t, expectedOp, ParseFile("test_input.txt"))
}

func TestAddColumnForEmpty(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
	}
	assert.Equal(t, expectedOp, AddColumnForEmpty(ParseFile("test_input.txt")))
}

func TestFindAllGalaxies(t *testing.T) {
	expectedOp := []Point{
		{0, 4},
		{1, 9},
		{2, 0},
		{5, 8},
		{6, 1},
		{7, 12},
		{10, 9},
		{11, 0},
		{11, 5},
	}
	assert.Equal(t, expectedOp, FindAllGalaxies("test_input.txt"))
}

func TestFindDistance(t *testing.T) {
	assert.Equal(t, 15, FindDistance(Point{x: 0, y: 4}, Point{x: 10, y: 9}))
}

func TestFindSumOfGalaxyDistances(t *testing.T) {
	assert.Equal(t, 374, FindSumOfGalaxyDistances("test_input.txt"))
}

func TestParseFile2(t *testing.T) {
	expectedOp := [][]rune{
		{'.', '.', '.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	}
	matrix, emptyRows, emptyCols := ParseFile2("test_input.txt")
	assert.Equal(t, expectedOp, matrix)
	assert.Equal(t, []int{3, 7}, emptyRows)
	assert.Equal(t, []int{2, 5, 8}, emptyCols)
}

func TestFindAllGalaxies2(t *testing.T) {
	allGalaxies, _, _ := FindAllGalaxies2("test_input.txt")
	assert.Equal(t, []Point{{0, 3}, {1, 7}, {2, 0}, {4, 6}, {5, 1}, {6, 9}, {8, 7}, {9, 0}, {9, 4}}, allGalaxies)
}

func TestFindSumDistance2(t *testing.T) {
	assert.Equal(t, 1030, FindSumDistance2("test_input.txt", 10))
	assert.Equal(t, 8410, FindSumDistance2("test_input.txt", 100))
}
