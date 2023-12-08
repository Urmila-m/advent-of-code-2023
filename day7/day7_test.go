package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcTotalWinnings(t *testing.T) {
	assert.Equal(t, 6440, CalcTotalWinnings("lib/test_input.txt"))
}

func TestCalcTotalWinningsWJoker(t *testing.T) {
	assert.Equal(t, 5905, CalcTotalWinningsWJoker("lib/test_input.txt"))
}
