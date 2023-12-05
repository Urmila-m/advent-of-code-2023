package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOverallPoints(t *testing.T) {
	overAllPoints := FindOverallPoints("lib/test_input.txt")
	assert.Equal(t, 13, overAllPoints)
}
