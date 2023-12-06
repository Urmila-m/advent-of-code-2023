package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLowestLocation(t *testing.T) {
	seed, location := FindLowestLocation("lib/test_input.txt")
	assert.Equal(t, 13, seed)
	assert.Equal(t, 35, location)
}

func TestFindLowestLocationFromSeedRange(t *testing.T) {
	seed, location := FindLowestLocationFromSeedRange("lib/test_input.txt")
	assert.Equal(t, 82, seed)
	assert.Equal(t, 46, location)
}
