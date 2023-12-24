package lib

import (
	d10 "day10/lib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTrench(t *testing.T) {
	for _, cube := range FindTrench(ParseFile("test_input.txt")) {
		fmt.Printf("+%v\n", cube)
	}
}

func TestCalcCubeHoldCapacity(t *testing.T) {
	allBorderPoints := make([]d10.Point, 0)
	for _, cube := range FindTrench(ParseFile("test_input.txt")) {
		allBorderPoints = append(allBorderPoints, cube.Position)
	}
	assert.Equal(t, 62, CalcCubeHoldCapacity(allBorderPoints))

	assert.Equal(t, 952408144115, CalcCubeHoldCapacity(FindTrench2(ParseFile("test_input.txt"))))
}
