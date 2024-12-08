package internal

import (
	"fmt"
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

	// todo::从有效牌当中找到joker牌并且取出来还是pure的顺子
	for _, card := range h.valid {
		if card.Value == wild {
			tJoker = append(tJoker, card)
		}
	}
	tSeq = h.handSliceDifference(h.valid, tJoker)

	// 小于3就不是pure顺
	if (len(tSeq)) < 3 {
		fmt.Println("抽离后他不是一个合法的顺子")
		return
	}

	if !h.judgeIsSeq(tSeq) {
		fmt.Println("judge后他不是一个合法的顺子")
		return
	}

	h.joker = append(h.joker, h.handSliceDifference(h.valid, tSeq)...)
	h.valid = h.handSliceIntersection(h.valid, tSeq)
}
