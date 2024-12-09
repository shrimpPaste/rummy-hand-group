package internal

import (
	"rummy-group-v2/pkg/app"
)

// 处理刻子的牌
func (h *Hand) find111Cards() {
	jokerLen := len(h.joker)

	if len(h.invalid) == 1 && jokerLen >= 2 {
		h.valid = append(h.valid, h.invalid...)
		h.valid = append(h.valid, h.joker...)
		h.joker = h.joker[2:]
	}

	setCards := h.findSetFromCards([]app.Card{}, h.invalid)

	if len(setCards) < 0 {
		return
	}

	for i, p := range h.pure {
		tSet := app.Card{}
		if p[0].Value == setCards[0].Value && p[0].Suit != setCards[0].Suit && h.judgeIsSeq(p[1:]) {
			tSet = p[0]
		}
		if p[len(p)-1].Value == setCards[0].Value && p[0].Suit != setCards[0].Suit && h.judgeIsSeq(p[:len(p)-1]) {
			tSet = p[len(p)-1]
		}
		setCards = append(setCards, tSet)
		if len(setCards) >= 3 {
			h.invalid = h.handSliceDifference(h.invalid, setCards)
			h.set = append(h.set, setCards)
			h.pure[i] = h.handSliceDifference(h.pure[i], []app.Card{tSet})
		}
	}

	// todo::从间隙牌中找刻子

	if len(setCards) >= 3 {
		return
	}

	if len(h.joker) > 1 && len(setCards) >= 2 {
		setCards = append(setCards, h.joker[0])
		h.setWithJoker = append(h.setWithJoker, setCards)
		h.joker = h.removeByIndex(h.joker, 1)
		h.invalid = h.handSliceDifference(h.invalid, setCards)
	}
}

// findSetFromCards 找一组卡牌中的刻子
func (h *Hand) findSetFromCards(result, cards []app.Card) []app.Card {
	if len(cards) == 0 {
		return result
	}
	if len(result) == 0 {
		result = append(result, cards[0])
		cards = cards[1:]
	}

	// 检查下一张牌是否连续
	if cards[0].Value == result[len(result)-1].Value {
		isSameSuit := false
		for _, r := range result {
			if r.Suit != cards[0].Suit {
				isSameSuit = true
			}
		}
		if isSameSuit {
			result = append(result, cards[0])
		}
	}
	return h.findSetFromCards(result, cards[1:])
}
