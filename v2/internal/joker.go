package internal

import "rummy-group-v2/pkg/app"

// 关于joker处理

// 找无效牌中的joker
func (h *Hand) findInvalidJoker(wild int) {
	for _, card := range h.invalid {
		if card.Value == wild || card.Suit == app.JokerA || card.Suit == app.JokerB {
			h.joker = append(h.joker, card)
		}
	}
	h.invalid = h.handSliceDifference(h.invalid, h.joker)
}
