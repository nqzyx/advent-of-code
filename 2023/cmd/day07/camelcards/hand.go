package camelcards

type Hand struct {
	Bid           int
	CardList      string
	Cards         Cards
	Strength      HandStrength
	Winnings      int64
	MatchingCards map[Card]int
	JokerCount    int
}

type Hands []*Hand

func NewHand(cardList string, bid int, jokerRule bool) (hand *Hand) {
	hand = &Hand{
		Bid:           bid,
		CardList:      cardList,
		Cards:         *NewCards(cardList, jokerRule),
		MatchingCards: make(map[Card]int),
	}
	hand.CalculateStrength()
	return
}

func (h *Hand) CalculateStrength() {
	for _, card := range h.Cards {
		if card == 0 {
			h.JokerCount++
		} else {
			h.MatchingCards[card]++
		}
	}

	cardGroupCounts := make(map[int]int)
	for _, count := range h.MatchingCards {
		cardGroupCounts[count]++
	}

	switch true {
	case cardGroupCounts[5] == 1:
		h.Strength = FiveOfAKind
	case cardGroupCounts[4] == 1:
		switch h.JokerCount {
		case 1:
			h.Strength = FiveOfAKind
		case 0:
			h.Strength = FourOfAKind
		}
	case cardGroupCounts[3] == 1 && cardGroupCounts[2] == 1:
		h.Strength = FullHouse
	case cardGroupCounts[3] == 1:
		switch h.JokerCount {
		case 2:
			h.Strength = FiveOfAKind
		case 1:
			h.Strength = FourOfAKind
		case 0:
			h.Strength = ThreeOfAKind
		}
	case cardGroupCounts[2] == 2:
		switch h.JokerCount {
		case 1:
			h.Strength = FullHouse
		case 0:
			h.Strength = TwoPair
		}
	case cardGroupCounts[2] == 1:
		switch h.JokerCount {
		case 3:
			h.Strength = FiveOfAKind
		case 2:
			h.Strength = FourOfAKind
		case 1:
			h.Strength = ThreeOfAKind
		case 0:
			h.Strength = OnePair
		}
	default:
		switch h.JokerCount {
		case 4, 5:
			h.Strength = FiveOfAKind
		case 3:
			h.Strength = FourOfAKind
		case 2:
			h.Strength = ThreeOfAKind
		case 1:
			h.Strength = OnePair
		case 0:
			h.Strength = HighCard
		}
	}
}

func (h *Hand) CalculateWinnings(rank int) int64 {
	h.Winnings = int64(h.Bid * (rank + 1))
	return h.Winnings
}
