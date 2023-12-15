package camelcards

type Hand struct {
	Bid           int
	CardList      string
	Cards         Cards
	Strength      HandStrength
	Winnings      int64
	MatchingCards map[Card]int
	JokerCount    int
	Rank          int
}

type Hands []*Hand

func NewHand(cardList string, bid int, jokerRule bool) (hand *Hand) {
	hand = &Hand{
		Bid:      bid,
		CardList: cardList,
		Cards:    *NewCards(cardList, jokerRule),
	}
	hand.Evaluate()
	return
}

func (c *Cards) FindMatchingCards() (matchingCards map[Card]int) {
	matchingCards = make(map[Card]int)
	for _, card := range c {
		matchingCards[card]++
	}
	return
}

func (h *Hand) Evaluate() (hs HandStrength) {
	h.MatchingCards = h.Cards.FindMatchingCards()
	h.JokerCount = h.MatchingCards[0]
	delete(h.MatchingCards, 0)

	cardGroupCounts := make(map[int]int)
	for _, count := range h.MatchingCards {
		cardGroupCounts[count]++
	}

	switch true {
	case cardGroupCounts[5] == 1:
		hs = FiveOfAKind
	case cardGroupCounts[4] == 1:
		switch h.JokerCount {
		case 0:
			hs = FourOfAKind
		case 1:
			hs = FiveOfAKind
		}
	case cardGroupCounts[3] == 1 && cardGroupCounts[2] == 1:
		hs = FullHouse
	case cardGroupCounts[3] == 1:
		switch h.JokerCount {
		case 0:
			hs = ThreeOfAKind
		case 1:
			hs = FourOfAKind
		case 2:
			hs = FiveOfAKind
		}
	case cardGroupCounts[2] == 2:
		switch h.JokerCount {
		case 0:
			hs = TwoPair
		case 1:
			hs = FullHouse
		}
	case cardGroupCounts[2] == 1:
		switch h.JokerCount {
		case 0:
			hs = OnePair
		case 1:
			hs = ThreeOfAKind
		case 2:
			hs = FourOfAKind
		case 3:
			hs = FiveOfAKind
		}
	case cardGroupCounts[1] == 5:
		hs = HighCard
	default:
		switch h.JokerCount {
		case 4, 5:
			hs = FiveOfAKind
		case 3:
			hs = FourOfAKind
		case 2:
			hs = ThreeOfAKind
		case 1:
			hs = OnePair
		default:
			hs = HighCard
		}
	}
	h.Strength = hs
	return
}

func (h *Hand) CalculateWinnings() int64 {
	h.Winnings = int64(h.Bid * h.Rank)
	return h.Winnings
}
