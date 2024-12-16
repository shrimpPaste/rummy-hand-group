package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
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
	isTest := false

	var jokerValueRand int
	if isTest {
		desk := InitializeDeck()
		ShuffleDeck(desk)
		headCard := DealCards(&desk, 13)

		for _, cc := range headCard {
			fmt.Printf("{Suit: app.%s, Value: %d},\n", cc.Suit, cc.Value)
		}

		suitCards := make(map[string][]app.Card, 4)
		h.groupCards(suitCards, headCard)

		var headCardRes []app.Card
		for _, cards := range suitCards {
			headCardRes = append(headCardRes, cards...)
		}

		h.SetCards(headCardRes)

		jokerValueRand = rand.Intn(13) + 1
		fmt.Println("Joker值是", jokerValueRand)

	} else {
		h.SetCards([]app.Card{
			{Suit: app.B, Value: 12},
			{Suit: app.A, Value: 4},
			{Suit: app.D, Value: 4},
			{Suit: app.B, Value: 4},
			{Suit: app.D, Value: 6},
			{Suit: app.D, Value: 8},
			{Suit: app.B, Value: 11},
			{Suit: app.D, Value: 3},
			{Suit: app.A, Value: 8},
			{Suit: app.B, Value: 10},
			{Suit: app.D, Value: 9},
			{Suit: app.D, Value: 5},
			{Suit: app.A, Value: 12},
		})
		jokerValueRand = 5
	}

	jokerRand := app.Card{Suit: app.D, Value: jokerValueRand}

	suitRand := rand.Intn(4)
	if suitRand == 0 {
		jokerRand.Suit = app.A
	} else if suitRand == 1 {
		jokerRand.Suit = app.C
	} else if suitRand == 2 {
		jokerRand.Suit = app.B
	}

	h.SetWildJoker(jokerRand)

	// TODO::第一步先去找牌堆中所有的三条,同时剩下的仍然能组成顺子
	overCards, setCards, scoreMapCards := h.findSet(h.cards)

	// TODO:: scoreMapCards 找到的顺子没有被删掉

	pureCards, overCards := h.GetPure(overCards)
	// TODO:: 第一步鉴定是否有顺子没有则中断
	if h.judgeIsHave1Seq(pureCards) {
		// 先找三条后还能找到找到顺子
		fmt.Println("pure", scoreMapCards)

		// TODO:: 一开始找的时候，不要抽离Joker，在找完纯顺子再去找Joker。
	} else {
		c.JSON(200, gin.H{
			"myCards":       getCardsResult(h.cards),
			"calcCards":     getCardsResult([]app.Card{}),
			"pure":          getCardsResult([]app.Card{}),
			"pureWithJoker": getCardsResult([]app.Card{}),
			"set":           getCardsResult([]app.Card{}),
			"setWithJoker":  getCardsResult([]app.Card{}),
			"invalid":       getCardsResult(h.cards),
			"joker":         getCardsResult([]app.Card{}),
			"sysJoker":      getCardsResult([]app.Card{h.wild}),
		})
		return
	}

	jokers, overCards := h.findJoker(overCards)

	// TODO:: 第二步找无效牌中间隙牌+joker分值最高的牌
	//overCards, pureWithCards, jokers := h.findGapMostScoreCards(overCards, jokers)

	var setWithJoker, pureWithCards, setCards2 []app.Card

	// TODO::1. 如果有一个joker，就要去找间隙 < 3
	overCards, pureWithCards, jokers = h.findGapsByJoker(overCards, jokers)
	// TODO::2. 如果有2个joker，就要去找间隙 == 3 如果还是没有 就去两个joker + 一个点数最大的牌。

	if len(pureCards) >= 6 || (len(pureCards) >= 3 && len(pureWithCards) >= 3) {
		// TODO:: 第三步从无效牌中找到两个相同值但是花色不同的牌 (不带joker的癞子)
		overCards, setCards2, scoreMapCards = h.findSet(overCards)

		if len(setCards2) > 0 {
			setCards2 = h.handSliceDifference(setCards2, setCards)
			setCards = append(setCards, setCards2...)
		}

		// TODO:: 第四步从无效牌中找到两个相同值但是花色不同的牌 (带joker的癞子)
		overCards, setWithJoker, jokers = h.findSetWithJoker2(overCards, jokers)
	} else {
		overCards = append(overCards, setCards...)
		setCards = []app.Card{}
	}

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
