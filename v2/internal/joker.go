package internal

import (
	"rummy-group-v2/pkg/app"
)

// 关于joker处理

// 找无效牌中的joker
func (h *Hand) findInvalidJoker(wild int) {
	for _, card := range h.invalid {
		if card.Value == wild || card.Suit == app.JokerA || card.Suit == app.JokerB {
			h.joker = append(h.joker, card)
		}
	}
	h.invalid = h.handSliceDifference(h.invalid, h.joker)

	var tJoker []app.Card
	var tSeq []app.Card

	validSuitCards := map[string][]app.Card{}

	h.groupCards(validSuitCards, h.valid)

	for _, cards := range validSuitCards {
		for _, card := range cards {
			if card.Value == wild {
				tJoker = append(tJoker, card)
			}
		}
		tSeq = h.handSliceDifference(h.valid, tJoker)

		// 小于3就不是pure顺
		if (len(tSeq)) < 3 {
			//fmt.Println("抽离后他不是一个合法的顺子")
			continue
		}

		if !h.judgeIsSeq(tSeq) {
			//fmt.Println("judge后他不是一个合法的顺子")
			continue
		}

		h.pure = append(h.pure, h.handSliceIntersection(h.valid, tSeq))
		h.valid = h.handSliceDifference(h.valid, append(tJoker, tSeq...))
		h.joker = append(h.joker, tJoker...)
	}

}
