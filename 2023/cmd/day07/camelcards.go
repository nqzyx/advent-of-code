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

type Card byte

const (
	Two Card = '2' + iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Cards [5]Card

type Hand struct {
	Bid      int
	Cards    Cards
	Rank     int
	Strength HandStrength
	Winnings int64
}

func NewCards(cardList string) (cards *Cards) {
	for i, c := range cards {
		if ok := c.(Card); ok {
			cards[i] = c
		} else {
		}
		cards[i] = c
	}
	return
}

func NewHand(cardList string, bid int) *Hand {
	return &Hand{
		Bid:   bid,
		Cards: *NewCards(cardList),
	}
}
