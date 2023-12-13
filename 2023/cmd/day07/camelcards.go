package main

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

type Cards [5]int

type Hand struct {
	Bid      int
	Cards    Cards
	Rank     int
	Strength HandStrength
	Winnings int64
}

func NewCards(cardList string) (cards *Cards) {
	cards = new(Cards)
	for i, c := range cardList {
		switch c {
		case '2', '3', '4', '5', '6', '7', '8', '9':
			cards[i] = int(c - '0')
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

func NewHand(cardList string, bid int) *Hand {
	return &Hand{
		Bid:   bid,
		Cards: *NewCards(cardList),
	}
}
