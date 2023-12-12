package lib

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestPipe_FindAdjacentPipes(t *testing.T) {
	lines := ParseFile("test_input2.txt")
	pipe := Pipe{Letter: 'S', Position: Point{1, 1}}
	expectedOp := map[Point]rune{
		Point{1, 2}: '-',
		Point{2, 1}: '|',
	}
	assert.Equal(t, expectedOp, pipe.FindAdjacentPipes(lines))
}

func areSlicesEqual(slice1, slice2 []Pipe) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func TestFindMainLoop(t *testing.T) {
	actualOp := FindMainLoop("test_input3.txt")
	expectedOp := []Pipe{
		Pipe{'S', Point{1, 1}},
		Pipe{'|', Point{2, 1}},
		Pipe{'L', Point{3, 1}},
		Pipe{'-', Point{3, 2}},
		Pipe{'J', Point{3, 3}},
		Pipe{'|', Point{2, 3}},
		Pipe{'7', Point{1, 3}},
		Pipe{'-', Point{1, 2}},
		Pipe{'S', Point{1, 1}},
	}
	seriallyEqual := areSlicesEqual(expectedOp, actualOp)
	slices.Reverse(expectedOp)
	reverseEqual := areSlicesEqual(expectedOp, actualOp)
	assert.True(t, seriallyEqual || reverseEqual)

	expectedOp = []Pipe{
		{'S', Point{2, 0}},
		{'|', Point{3, 0}},
		{'L', Point{4, 0}},
		{'J', Point{4, 1}},
		{'F', Point{3, 1}},
		{'-', Point{3, 2}},
		{'-', Point{3, 3}},
		{'J', Point{3, 4}},
		{'7', Point{2, 4}},
		{'L', Point{2, 3}},
		{'|', Point{1, 3}},
		{'7', Point{0, 3}},
		{'F', Point{0, 2}},
		{'J', Point{1, 2}},
		{'F', Point{1, 1}},
		{'J', Point{2, 1}},
		{'S', Point{2, 0}},
	}
	actualOp = FindMainLoop("test_input5.txt")
	seriallyEqual = areSlicesEqual(expectedOp, actualOp)
	slices.Reverse(expectedOp)
	reverseEqual = areSlicesEqual(expectedOp, actualOp)
	assert.True(t, seriallyEqual || reverseEqual)
}

func TestFindNumOfStepsForFarthestPoint(t *testing.T) {
	assert.Equal(t, 4, FindNumOfStepsForFarthestPoint("test_input3.txt"))
	assert.Equal(t, 8, FindNumOfStepsForFarthestPoint("test_input5.txt"))
}

func TestFindNumOfInsidePoints(t *testing.T) {
	borderPipes := FindMainLoop("test_input6.txt")

	borderPoints := make([]Point, 0)
	for _, pipe := range borderPipes[:len(borderPipes)-1] {
		borderPoints = append(borderPoints, pipe.Position)
	}
	assert.Equal(t, 8, FindNumOfInsidePoints(FindAreaUsingShoeLace(borderPoints), len(borderPoints)))
}
