package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindStrength(t *testing.T) {
	cards := [5]Card{3, 2, T, 3, K}
	assert.Equal(t, "OnePair", FindStrength(cards).String())
}
func TestParseHand(t *testing.T) {
	hand := Hand{Cards: [5]Card{3, 2, T, 3, K}, Strength: OnePair, Bid: 765}
	assert.Equal(t, hand, ParseHand("32T3K 765", false))
}

func TestFindStrengthWJoker(t *testing.T) {
	assert.Equal(t, FourOfAKind, FindStrengthWJoker([5]Card{T, 5, 5, J, 5}))
	assert.Equal(t, FourOfAKind, FindStrengthWJoker([5]Card{K, T, J, J, T}))
	assert.Equal(t, FourOfAKind, FindStrengthWJoker([5]Card{Q, Q, Q, J, A}))
}

func TestSortHands(t *testing.T) {
	sortedHands := []Hand{
		{Cards: [5]Card{3, 2, T, 3, K}, Bid: 765, Strength: OnePair},
		{Cards: [5]Card{K, T, J, J, T}, Bid: 220, Strength: TwoPair},
		{Cards: [5]Card{K, K, 6, 7, 7}, Bid: 28, Strength: TwoPair},
		{Cards: [5]Card{T, 5, 5, J, 5}, Bid: 684, Strength: ThreeOfAKind},
		{Cards: [5]Card{Q, Q, Q, J, A}, Bid: 483, Strength: ThreeOfAKind},
	}
	assert.Equal(t, sortedHands, SortHands(ParseFile("test_input.txt", false), false))

	hands := []Hand{
		{[5]Card{5, 9, 6, K, 8}, HighCard, 23},
		{[5]Card{3, 7, 2, 9, A}, HighCard, 14},
		{[5]Card{5, 6, 7, J, Q}, HighCard, 398},
	}
	sortedHands = []Hand{
		{[5]Card{3, 7, 2, 9, A}, HighCard, 14},
		{[5]Card{5, 6, 7, J, Q}, HighCard, 398},
		{[5]Card{5, 9, 6, K, 8}, HighCard, 23},
	}
	assert.Equal(t, sortedHands, SortHands(hands, false))

	sortedHands = []Hand{
		{Cards: [5]Card{3, 2, T, 3, K}, Bid: 765, Strength: OnePair},
		{Cards: [5]Card{K, K, 6, 7, 7}, Bid: 28, Strength: TwoPair},
		{Cards: [5]Card{T, 5, 5, J, 5}, Bid: 684, Strength: FourOfAKind},
		{Cards: [5]Card{Q, Q, Q, J, A}, Bid: 483, Strength: FourOfAKind},
		{Cards: [5]Card{K, T, J, J, T}, Bid: 220, Strength: FourOfAKind},
	}
	assert.Equal(t, sortedHands, SortHands(ParseFile("test_input.txt", true), true))
}
