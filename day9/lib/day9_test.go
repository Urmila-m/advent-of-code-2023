package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	expectedOp := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	assert.Equal(t, expectedOp, ParseFile("test_input.txt"))
}

func TestPredictNextValue(t *testing.T) {
	assert.Equal(t, 18, PredictNextValue([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, PredictNextValue([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, PredictNextValue([]int{10, 13, 16, 21, 30, 45}))
}

func TestSumExtrapolatedValues(t *testing.T) {
	assert.Equal(t, 114, SumExtrapolatedValues("test_input.txt", "next"))
	assert.Equal(t, 2, SumExtrapolatedValues("test_input.txt", "previous"))
}

func TestPredictPreviousValue(t *testing.T) {
	assert.Equal(t, -3, PredictPreviousValue([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 0, PredictPreviousValue([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 5, PredictPreviousValue([]int{10, 13, 16, 21, 30, 45}))
}
