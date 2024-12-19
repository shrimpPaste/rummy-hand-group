package logic

import (
	"rummy-logic-v3/pkg/app"
)

type Prompt struct {
	hand *Hand
}

func (p *Prompt) Calculate() [][]int {
	h := p.hand

	// 1. 找纯刻子
	overCards, setCards, _ := h.findSet(h.cards)
	// 2. 找纯顺子
	pureCards, overCards := h.GetPure(overCards)
	if !h.judgeIsHave1Seq(pureCards) {
		// 顺子不够，先把抽离出的刻子放回去再次检查
		overCards = append(overCards, setCards...)
		setCards = []app.Card{}
		pureCards, overCards = h.GetPure(overCards)
		if !h.judgeIsHave1Seq(pureCards) {
			resp := p.GetResponse(h.cards)
			return resp
		}
	}

	// 3. 找joker
	jokers, overCards := h.findJoker(overCards)
	var setWithJoker, pureWithCards, setCards2 []app.Card
	if len(pureCards) < 6 {
		// 顺子不够就要先找顺子，满足两个顺子才行
		if len(setCards) > 0 {
			overCards = append(overCards, setCards...)
			setCards = []app.Card{}
		}
	}

	// 找带joker的顺子
	overCards, pureWithCards, jokers = h.findGapsByJoker(overCards, jokers)

	if h.judgeIsHave2Seq(pureCards, pureWithCards) {
		overCards, setCards2, _ = h.findSet(overCards)

		if len(setCards2) > 0 {
			setCards2 = h.handSliceDifference(setCards2, setCards)
			setCards = append(setCards, setCards2...)
		}

		overCards, setWithJoker, jokers = h.findSetWithJoker2(overCards, jokers)
	} else {
		overCards = append(overCards, setCards...)
		overCards = append(overCards, pureCards...)

		pureCards, overCards = h.GetPure(overCards)
		setCards = []app.Card{}
	}

	respOverCards := make([]app.Card, len(overCards))
	copy(respOverCards, overCards)
	respOverCards = append(respOverCards, jokers...)

	resp := p.GetResponse(pureCards, pureWithCards, setCards, setWithJoker, respOverCards)
	return resp
}

func (p *Prompt) GetResponse(cards ...[]app.Card) [][]int {
	var res [][]int

	for _, card := range cards {
		if len(card) > 0 {
			res = append(res, GetCardsResult(card))
		}
	}
	return res
}

func NewPrompt(cards []app.Card, joker app.Card) *Prompt {
	h := NewHand()
	h.SetCards(cards)
	h.SetWildJoker(joker)
	return &Prompt{
		hand: h,
	}
}
