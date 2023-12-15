package camelcards

type HandStrength int

const (
	HighCard     HandStrength = 10
	OnePair      HandStrength = 20
	TwoPair      HandStrength = 22
	ThreeOfAKind HandStrength = 30
	FullHouse    HandStrength = 32
	FourOfAKind  HandStrength = 40
	FiveOfAKind  HandStrength = 50
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
