package internal

import (
	"fmt"
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
		{Value: 3, Suit: app.D},
		{Value: 4, Suit: app.D},
		{Value: 5, Suit: app.D},
		{Value: 6, Suit: app.D},

		{Value: 2, Suit: app.C},
		{Value: 4, Suit: app.C},
		{Value: 5, Suit: app.C},

		{Value: 5, Suit: app.B},
		{Value: 6, Suit: app.B},
		{Value: 3, Suit: app.B},
		{Value: 6, Suit: app.B},

		{Value: 2, Suit: app.A},
		{Value: 3, Suit: app.A},
	})

	h.SetWildJoker(app.Card{Suit: app.A, Value: 6})

	jokers, overCards := h.findJoker(h.cards)

	tempOverCards := overCards
	// TODO::第一步先去找牌堆中所有的三条,同时剩下的仍然能组成顺子
	overCards, setCards, scoreMapCards := h.findSet(overCards)

	pureCards, overCards := h.GetPure(overCards)
	// TODO:: 第一步鉴定是否有顺子没有则中断
	if h.judgeIsHave1Seq(pureCards) {
		// 先找三条后还能找到找到顺子
		fmt.Println(scoreMapCards)
	} else {
		overCards = tempOverCards
		// 找不到顺子
		pureCards, overCards = h.GetPure(overCards)
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
	}

	// TODO:: 第二步找无效牌中间隙牌+joker分值最高的牌
	overCards, pureWithCards, jokers := h.findGapMostScoreCards(overCards, jokers)

	// TODO:: 第四步从无效牌中找到两个相同值但是花色不同的牌 (带joker的癞子)
	overCards, setWithJoker, jokers := h.findSetWithJoker(overCards, jokers)

	c.JSON(200, gin.H{
		"myCards":       getCardsResult(h.cards),
		"calcCards":     getCardsResult([]app.Card{}),
		"pure":          getCardsResult(pureCards),
		"pureWithJoker": getCardsResult(pureWithCards),
		"set":           getCardsResult(setCards),
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
