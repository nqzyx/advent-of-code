package main

import (
	"strconv"

	"golang.org/x/exp/slices"
)

type HandStrength int

const (
	_ HandStrength = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h HandStrength) String() string {
	switch h {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	default:
		return "Unknown"
	}
}

type (
	Card  int
	Cards [5]Card
)

func (c Card) String() string {
	switch c {
	case 10:
		return "T"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	default:
		return strconv.Itoa(int(c))
	}
}

type Hand struct {
	Bid           int
	Cards         Cards
	Rank          int
	Strength      HandStrength
	Winnings      int64
	MatchingCards map[Card]int
}

type Game []*Hand

func NewCards(cardList string) (cards *Cards) {
	cards = new(Cards)
	for i, c := range cardList {
		switch c {
		case '2', '3', '4', '5', '6', '7', '8', '9':
			cards[i] = Card(c - '0')
		case 'T':
			cards[i] = 10
		case 'J':
			cards[i] = 11
		case 'Q':
			cards[i] = 12
		case 'K':
			cards[i] = 13
		case 'A':
			cards[i] = 14
		}
	}
	return
}

func NewHand(cardList string, bid int) (hand *Hand) {
	hand = &Hand{
		Bid:   bid,
		Cards: *NewCards(cardList),
	}
	hand.CalculateStrength()
	return
}

func (h *Hand) CalculateStrength() {
	matching := make(map[Card]int)

	for _, card1 := range h.Cards {
		matching[card1]++
	}

	var pairCount, threeCount, fourCount, fiveCount int
	for _, count := range matching {
		switch count {
		case 2:
			pairCount++
		case 3:
			threeCount++
		case 4:
			fourCount++
		case 5:
			fiveCount++
		}
	}

	switch true {
	case fiveCount == 1:
		h.Strength = FiveOfAKind
	case fourCount == 1:
		h.Strength = FourOfAKind
	case threeCount == 1 && pairCount == 1:
		h.Strength = FullHouse
	case threeCount == 1:
		h.Strength = ThreeOfAKind
	case pairCount == 2:
		h.Strength = TwoPair
	case pairCount == 1:
		h.Strength = OnePair
	default:
		h.Strength = HighCard
	}
}

func (g Game) RankHands() {
	slices.SortFunc(g, func(a *Hand, b *Hand) int {
		if a.Strength != b.Strength {
			return int(a.Strength - b.Strength)
		}

		for i, cardA := range a.Cards {
			cardB := b.Cards[i]
			if cardA != cardB {
				return int(cardA - cardB)
			}
		}

		return 0
	})
}

func (g Game) CalculateWinnings() (winnings int) {
	for rank, hand := range g {
		winnings += hand.Bid * (rank + 1)
	}
	return
}
