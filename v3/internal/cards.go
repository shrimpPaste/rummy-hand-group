package internal

import (
	"github.com/gin-gonic/gin"
	"rummy-logic-v3/pkg/app"
	"sort"
)

// Hand 手牌
type Hand struct {
	cards []app.Card
	wild  app.Card
}

func (h *Hand) Run(r *gin.Engine) {
	r.GET("/", h.WebGet)
}

// GetPure 获取纯顺子
func (h *Hand) GetPure(cards []app.Card) (pureCards, overCards []app.Card) {
	// 移除joker牌
	for _, card := range cards {
		if card.Suit == app.JokerA || card.Suit == app.JokerB || card.Value == h.wild.Value {
			// 添加joker
			overCards = append(overCards, card)

			cards = h.handSliceDifference(cards, []app.Card{card})
		}
	}

	suitCards := make(map[string][]app.Card, 4)
	h.groupCards(suitCards, cards)

	for _, c := range suitCards {
		sort.Slice(c, func(i, j int) bool {
			return c[i].Value < c[j].Value
		})
		sequence := h.findValidSequence(c)

		if len(sequence) < 3 {
			overCards = append(overCards, c...)
		} else {
			pureCards = append(pureCards, sequence...)
			overCards = append(overCards, h.handSliceDifference(c, sequence)...)
		}
	}

	return
}

func (h *Hand) WebGet(c *gin.Context) {
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 1},
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},

		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 12},

		{Suit: app.A, Value: 7},
		{Suit: app.A, Value: 8},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 1},
	})

	h.SetWildJoker(app.Card{Suit: app.A, Value: 1})

	pureCards, overCards := h.GetPure(h.cards)

	c.JSON(200, gin.H{
		"myCards":       getCardsResult(h.cards),
		"calcCards":     getCardsResult([]app.Card{}),
		"pure":          getCardsResult(pureCards),
		"pureWithJoker": getCardsResult([]app.Card{}),
		"set":           getCardsResult([]app.Card{}),
		"setWithJoker":  getCardsResult([]app.Card{}),
		"invalid":       getCardsResult(overCards),
		"joker":         getCardsResult([]app.Card{}),
		"sysJoker":      getCardsResult([]app.Card{h.wild}),
	})
	return
}

func NewHand() *Hand {
	return &Hand{
		cards: []app.Card{},
	}
}
