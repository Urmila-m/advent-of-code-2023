package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	expectedOp := map[string]map[string]string{
		"AAA": {"left": "BBB", "right": "CCC"},
		"BBB": {"left": "DDD", "right": "EEE"},
		"CCC": {"left": "ZZZ", "right": "GGG"},
		"DDD": {"left": "DDD", "right": "DDD"},
		"EEE": {"left": "EEE", "right": "EEE"},
		"GGG": {"left": "GGG", "right": "GGG"},
		"ZZZ": {"left": "ZZZ", "right": "ZZZ"},
	}
	instructions, actualOp := ParseFile("test_input2.txt")
	assert.Equal(t, expectedOp, actualOp)
	assert.Equal(t, "RL", instructions)
}

func TestFindNumOfSteps(t *testing.T) {
	instructions, network := ParseFile("test_input2.txt")
	mPath := FindNumOfSteps(instructions, network, "AAA", "ZZZ")
	assert.Equal(t, 2, mPath.NoOfSteps)
}

func TestFindAllPaths(t *testing.T) {
	instructions, network := ParseFile("test_input3.txt")
	allPaths := FindAllPaths(instructions, network)

	expectedOp := []Path{
		{
			Source:     "11A",
			Dest:       "11Z",
			NoOfSteps:  2,
			Trajectory: []string{"11B", "11Z"},
		},
		{
			Source:     "22A",
			Dest:       "22Z",
			NoOfSteps:  3,
			Trajectory: []string{"22B", "22C", "22Z"},
		},
	}
	assert.Equal(t, expectedOp, allPaths)
}

/*
	Test if the number of steps from src -> dst is equal to dst -> dst for all paths.

Passing this test would mean, that if we move with the `instructions` infinitely in any path(`..A` -> `..Z`) , they would follow src -> dst -> dst -> dst...
This also means that no. of steps from src -> dst and dst -> dst are multiple of number of moves(L/R) in the instructions.
*/
func TestEqualityOfNoOfSteps(t *testing.T) {
	instructions, network := ParseFile("../puzzle_input.txt")

	for _, path := range FindAllPaths(instructions, network) {
		dstToDstPath := FindNumOfSteps(instructions, network, path.Dest, path.Dest)
		assert.Equal(t, path.NoOfSteps, dstToDstPath.NoOfSteps)

		// Comparing the Trajectory yields that both the trajectory converges at some point and then follows the same pattern.
		// Check if the last two elements of the trajectory are same or not for the two paths
		assert.Equal(t, path.Trajectory[len(path.Trajectory)-2:], dstToDstPath.Trajectory[len(dstToDstPath.Trajectory)-2:])
	}
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 18, FindLCM([]int{6, 9, 18}))
}
