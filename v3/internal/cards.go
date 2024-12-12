package internal

import (
	"github.com/gin-gonic/gin"
	"rummy-logic-v3/pkg/app"
)

// Hand 手牌
type Hand struct {
	cards []app.Card
	wild  app.Card
}

func (h *Hand) Run(r *gin.Engine) {
	r.GET("/", h.WebGet)
}

func (h *Hand) WebGet(c *gin.Context) {
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 7},

		{Suit: app.C, Value: 9},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 13},

		{Suit: app.A, Value: 7},
		{Suit: app.A, Value: 8},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 13},
		{Suit: app.A, Value: 1},
	})

	h.SetWildJoker(app.Card{Suit: app.A, Value: 7})

	jokers, overCards := h.findJoker(h.cards)

	pureCards, overCards := h.GetPure(overCards)
	// TODO:: 第一步鉴定是否有顺子没有则中断
	if !h.judgeIsHave1Seq(pureCards) {
		c.JSON(200, gin.H{
			"myCards":       getCardsResult(h.cards),
			"calcCards":     getCardsResult([]app.Card{}),
			"pure":          getCardsResult([]app.Card{}),
			"pureWithJoker": getCardsResult([]app.Card{}),
			"set":           getCardsResult([]app.Card{}),
			"setWithJoker":  getCardsResult([]app.Card{}),
			"invalid":       getCardsResult(overCards),
			"joker":         getCardsResult([]app.Card{}),
			"sysJoker":      getCardsResult([]app.Card{h.wild}),
		})
		return
	}

	// TODO:: 第二步找无效牌中间隙牌+joker分值最高的牌
	overCards, pureWithCards, jokers := h.findGapMostScoreCards(overCards, jokers)

	// TODO:: 第三步从无效牌中找到两个相同值但是花色不同的牌 (带joker的癞子)
	overCards, setWithJoker, jokers := h.findSetWithJoker(overCards, jokers)

	c.JSON(200, gin.H{
		"myCards":       getCardsResult(h.cards),
		"calcCards":     getCardsResult([]app.Card{}),
		"pure":          getCardsResult(pureCards),
		"pureWithJoker": getCardsResult(pureWithCards),
		"set":           getCardsResult([]app.Card{}),
		"setWithJoker":  getCardsResult(setWithJoker),
		"invalid":       getCardsResult(overCards),
		"joker":         getCardsResult(jokers),
		"sysJoker":      getCardsResult([]app.Card{h.wild}),
	})
	return
}

func NewHand() *Hand {
	return &Hand{
		cards: []app.Card{},
	}
}
