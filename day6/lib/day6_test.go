package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	expectedOp := map[string][]int{
		"Time":     []int{7, 15, 30},
		"Distance": []int{9, 40, 200},
	}
	assert.Equal(t, expectedOp, ParseFile("test_input.txt"))
}

func TestMapTimeDistance(t *testing.T) {
	expectedOp := map[int]int{
		7:  9,
		15: 40,
		30: 200,
	}
	assert.Equal(t, expectedOp, MapTimeDistance(ParseFile("test_input.txt")))
}
func TestFindDistance(t *testing.T) {
	assert.Equal(t, 10, FindDistance(2, 7))
}

func TestFindNumWaysToBeat(t *testing.T) {
	assert.Equal(t, 4, FindNumWaysToBeat(7, 9))
	assert.Equal(t, 8, FindNumWaysToBeat(15, 40))
}

func TestParseFile2(t *testing.T) {
	time, distance := ParseFile2("test_input.txt")
	assert.Equal(t, 71530, time)
	assert.Equal(t, 940200, distance)
}
