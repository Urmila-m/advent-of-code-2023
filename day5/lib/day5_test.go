package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCategoryMapping(t *testing.T) {
	input := []string{"seed-to-soil map:", "50 98 2", "52 50 3"}
	expectedOp := CategoryMap{
		Src:     "seed",
		Dst:     "soil",
		Mapping: []SrcDstMapping{{SrcStart: 98, DstStart: 50, RangeLength: 2}, {SrcStart: 50, DstStart: 52, RangeLength: 3}},
	}
	assert.Equal(t, expectedOp, ToCategoryMapping(input))
}

func TestFindInitialSeeds(t *testing.T) {
	input := "seeds: 79 14 55 13"
	assert.Equal(t, []int{79, 14, 55, 13}, FindInitialSeeds(input))
}

func TestParseFile(t *testing.T) {
	_, allCategoryMaps := ParseFile("small_test_input.txt")
	expectedOp := []CategoryMap{{Src: "seed", Dst: "soil", Mapping: []SrcDstMapping{{98, 50, 2}, {50, 52, 48}}}, {Src: "soil", Dst: "fertilizer", Mapping: []SrcDstMapping{{15, 0, 37}, {52, 37, 2}, {0, 39, 15}}}}
	assert.Equal(t, expectedOp, allCategoryMaps)

}

func TestMapCategory(t *testing.T) {
	assert.Equal(t, 51, MapCategory(99, CategoryMap{Src: "seed", Dst: "soil", Mapping: []SrcDstMapping{{98, 50, 2}, {50, 52, 48}}}))
}

func TestFindLocation(t *testing.T) {
	_, allCategoryMaps := ParseFile("test_input.txt")
	assert.Equal(t, 82, FindLocation(79, allCategoryMaps))
}
