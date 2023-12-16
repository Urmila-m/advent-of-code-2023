package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	lines := RollToNorth("test_input.txt")
	assert.Equal(t, expectedOp, lines)
}

func TestCalcNorthLoad(t *testing.T) {
	assert.Equal(t, 136, CalcNorthLoad("test_input.txt"))
}
