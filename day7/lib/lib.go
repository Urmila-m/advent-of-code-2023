package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Card int

const (
	T Card = 10 + iota
	J
	Q
	K
	A
)

func (card Card) Value() int {
	return int(card)
}

func (card Card) ValueWithJoker() int {
	if card.String() == "J" {
		return 1
	} else {
		return int(card)
	}
}

func (card Card) String() string {
	if card.Value() > 9 {
		return [...]string{"T", "J", "Q", "K", "A"}[card-10]
	}
	return fmt.Sprintf("%d", card.Value())
}

func (card Card) IsGreaterThan(card2 Card, withJoker bool) bool {
	if withJoker {
		return card.ValueWithJoker() > card2.ValueWithJoker()
	} else {
		return card.Value() > card2.Value()
	}
}

func (card Card) IsLessThan(card2 Card, withJoker bool) bool {
	if withJoker {
		return card.ValueWithJoker() < card2.ValueWithJoker()
	} else {
		return card.Value() < card2.Value()
	}
}

func (card Card) IsEqual(card2 Card, withJoker bool) bool {
	if withJoker {
		return card.ValueWithJoker() == card2.ValueWithJoker()
	} else {
		return card.Value() == card2.Value()
	}
}

type HandType int

const (
	HighCard HandType = 1 + iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (s HandType) Value() int {
	return int(s)
}

func (s HandType) String() string {
	return [...]string{"HighCard", "OnePair", "TwoPair", "ThreeOfAKind", "FullHouse", "FourOfAKind", "FiveOfAKind"}[s-1]
}

func (s HandType) IsGreaterThan(strength HandType) bool {
	return s.Value() > strength.Value()
}

func (s HandType) IsLessThan(strength HandType) bool {
	return s.Value() < strength.Value()
}

func (s HandType) IsEqual(strength HandType) bool {
	return s.Value() == strength.Value()
}

type Hand struct {
	Cards    [5]Card
	Strength HandType
	Bid      int
}

func ParseFile(filePath string, withJoker bool) []Hand {
	file, fErr := os.Open(filePath)

	if fErr != nil {
		log.Fatal(fErr)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	allHands := make([]Hand, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		allHands = append(allHands, ParseHand(line, withJoker))
	}
	return allHands
}

func ParseHand(line string, withJoker bool) Hand {
	split := strings.Fields(line)
	cardsStr := split[0]
	bidStr := split[1]

	bid, convertErr := strconv.Atoi(bidStr)
	if convertErr != nil {
		log.Fatal(convertErr)
	}
	cards := make([]Card, 0, 5)
	for _, card := range cardsStr {
		if unicode.IsDigit(card) {
			card, _ := strconv.Atoi(string(card))
			cards = append(cards, Card(card))
		} else if string(card) == "T" {
			cards = append(cards, T)
		} else if string(card) == "J" {
			cards = append(cards, J)
		} else if string(card) == "Q" {
			cards = append(cards, Q)
		} else if string(card) == "K" {
			cards = append(cards, K)
		} else if string(card) == "A" {
			cards = append(cards, A)
		}
	}
	var strength HandType
	if withJoker {
		strength = FindStrengthWJoker([5]Card(cards))
	} else {
		strength = FindStrength([5]Card(cards))
	}
	return Hand{Cards: [5]Card(cards), Bid: bid, Strength: strength}
}

func SortHands(hands []Hand, withJoker bool) []Hand {
	for i := 0; i < len(hands); i++ {
		minHand := hands[i]
		minHandIndex := i
		remainingHands := hands[i+1:]
		for j, hand := range remainingHands {
			if hand.Strength.IsLessThan(minHand.Strength) {
				minHand = hand
				minHandIndex = j + (i + 1)
			} else if hand.Strength.IsEqual(minHand.Strength) {
				for k := 0; k < 5; k++ {
					if hand.Cards[k].IsEqual(minHand.Cards[k], withJoker) {
						continue
					}
					if hand.Cards[k].IsLessThan(minHand.Cards[k], withJoker) {
						minHand = hand
						minHandIndex = j + (i + 1)
						break
					} else {
						break
					}
				}
			}
		}

		tmp := hands[i]
		hands[i] = minHand
		hands[minHandIndex] = tmp
	}
	return hands
}

func FindCardCountMap(cards [5]Card) map[Card]int {
	cardCount := make(map[Card]int)
	for _, card := range cards {
		_, exists := cardCount[card]
		if exists {
			cardCount[card]++
		} else {
			cardCount[card] = 1
		}
	}
	return cardCount
}

func FindStrength(cards [5]Card) HandType {
	cardCountMap := FindCardCountMap(cards)
	maxCardCount := FindMaxCount(cardCountMap)
	switch len(cardCountMap) {
	case 1:
		return FiveOfAKind
	case 2:
		if maxCardCount == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		if maxCardCount == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	case 4:
		return OnePair
	}
	return HighCard
}

func FindMaxCount(cardCountMap map[Card]int) int {
	maxCount := -1
	for _, count := range cardCountMap {
		if maxCount == -1 || maxCount < count {
			maxCount = count
		}
	}
	return maxCount
}

func FindStrengthWJoker(cards [5]Card) HandType {
	cardCount := FindCardCountMap(cards)
	jokerCount, exists := cardCount[J]
	if !exists {
		return FindStrength(cards)
	}
	delete(cardCount, J)
	maxCountWOJoker := FindMaxCount(cardCount)
	maxCountWJoker := maxCountWOJoker + jokerCount

	switch len(cardCount) {
	case 0:
		return FiveOfAKind
	case 1:
		return FiveOfAKind
	case 2:
		switch maxCountWJoker {
		case 4:
			return FourOfAKind
		case 3:
			return FullHouse
		}
	case 3:
		return ThreeOfAKind
	case 4:
		return OnePair
	}
	return OnePair
}
