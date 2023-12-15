package camelcards

import (
	"strconv"
)

type Card int

func (c Card) String() string {
	switch c {
	case 10:
		return "T"
	case 11, 0: // 0 when JokerRule is in force...
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

type Cards [5]Card

func NewCards(cardList string, jokerRule bool) (cards *Cards) {
	cards = new(Cards)
	for i, c := range cardList {
		switch c {
		case '2', '3', '4', '5', '6', '7', '8', '9':
			cards[i] = Card(c - '0')
		case 'T':
			cards[i] = 10
		case 'J':
			if jokerRule { // handle the Joker Rule
				cards[i] = 0
			} else {
				cards[i] = 11
			}
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
