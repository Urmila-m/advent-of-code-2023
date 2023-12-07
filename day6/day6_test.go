package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPuzzle1Output(t *testing.T) {
	assert.Equal(t, 288, FindPuzzle1Output("lib/test_input.txt"))
}

func TestFindPuzzle2Output(t *testing.T) {
	assert.Equal(t, 71503, FindPuzzle2Output("lib/test_input.txt"))
}
