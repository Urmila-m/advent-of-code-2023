package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	expectedOp := [][]string{
		{
			"#.##..##.",
			"..#.##.#.",
			"##......#",
			"##......#",
			"..#.##.#.",
			"..##..##.",
			"#.#.##.#.",
		},
		{
			"#...##..#",
			"#....#..#",
			"..##..###",
			"#####.##.",
			"#####.##.",
			"..##..###",
			"#....#..#",
		},
	}
	assert.Equal(t, expectedOp, ParseFile("test_input.txt"))
}

func TestFindHorizontalMirror(t *testing.T) {
	chunks := ParseFile("test_input.txt")
	mirrorX, exists := FindHorizontalMirror(chunks[0])
	assert.Equal(t, []int{}, mirrorX)
	assert.Equal(t, false, exists)
	mirrorX, exists = FindHorizontalMirror(chunks[1])
	assert.Equal(t, []int{4}, mirrorX)
	assert.Equal(t, true, exists)
	mirrorX, exists = FindHorizontalMirror(chunks[2])
	assert.Equal(t, []int{2}, mirrorX)
	assert.Equal(t, true, exists)
}

func TestFindVerticalMirror(t *testing.T) {
	chunks := ParseFile("test_input.txt")
	mirrorX, exists := FindVerticalMirror(chunks[0])
	assert.Equal(t, []int{5}, mirrorX)
	assert.Equal(t, true, exists)
	mirrorX, exists = FindVerticalMirror(chunks[1])
	assert.Equal(t, []int{}, mirrorX)
	assert.Equal(t, false, exists)
}

func TestSummarizeNotes(t *testing.T) {
	assert.Equal(t, 605, SummarizeNotes("test_input.txt"))
}

func TestFindSmudgeHorizontalMirror(t *testing.T) {
	chunks := ParseFile("test_input.txt")
	mirrorX, exists := FindSmudgeHorizontalMirror(chunks[0])
	assert.Equal(t, true, exists)
	assert.Equal(t, 3, mirrorX)
	mirrorX, exists = FindSmudgeHorizontalMirror(chunks[1])
	assert.Equal(t, true, exists)
	assert.Equal(t, 1, mirrorX)
	mirrorX, exists = FindSmudgeHorizontalMirror(chunks[2])
	assert.Equal(t, true, exists)
	assert.Equal(t, 5, mirrorX)
}

func TestFindSmudgeVerticalMirror(t *testing.T) {
	chunk := ParseFile("test_input4.txt")[0]
	mirrorY, exists := FindSmudgeVerticalMirror(chunk)
	assert.Equal(t, true, exists)
	assert.Equal(t, 9, mirrorY)
}

func TestSummarizeNotesWithSmudge(t *testing.T) {
	assert.Equal(t, 400, SummarizeNotesWithSmudge("test_input2.txt"))
}
