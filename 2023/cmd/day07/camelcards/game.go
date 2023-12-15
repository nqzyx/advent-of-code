package camelcards

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Game struct {
	Hands         Hands
	JokerRule     bool
	TotalWinnings int64
}

func NewGame(handList []string, jokerRule bool) *Game {
	hands := make(Hands, 0, len(handList))
	for _, handString := range handList {
		if len(handString) == 0 {
			continue
		}
		handParts := strings.Split(handString, " ")
		bid, _ := strconv.Atoi(handParts[1])
		hands = append(hands, NewHand(handParts[0], bid, jokerRule))
	}
	g := &Game{
		Hands:     hands,
		JokerRule: jokerRule,
	}
	g.TotalWinnings = g.CalculateWinnings()
	return g
}

func (g *Game) RankHands() {
	slices.SortFunc(g.Hands, func(a *Hand, b *Hand) int {
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

func (g *Game) CalculateWinnings() int64 {
	g.RankHands()
	var winnings int64
	for rank, hand := range g.Hands {
		winnings += hand.CalculateWinnings(rank)
	}
	g.TotalWinnings = winnings

	return g.TotalWinnings
}
