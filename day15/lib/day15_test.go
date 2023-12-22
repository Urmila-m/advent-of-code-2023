package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHash(t *testing.T) {
	assert.Equal(t, 52, FindHash("HASH"))

}

func TestFindHashSum(t *testing.T) {
	assert.Equal(t, 1320, FindHashSum(ParseFile("test_input.txt")))
}

func TestFindResultingConfiguration(t *testing.T) {
	expectedOp := map[int][]map[string]int{
		0: {{"rn": 1}, {"cm": 2}},
		1: {},
		3: {{"ot": 7}, {"ab": 5}, {"pc": 6}},
	}
	fmt.Println(t, expectedOp, FindResultingConfiguration(ParseFile("test_input.txt")))
}

func TestSumAllFocusingPower(t *testing.T) {
	assert.Equal(t, 145, SumAllFocusingPower(FindResultingConfiguration(ParseFile("test_input.txt"))))
}
