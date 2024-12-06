package internal

import "rummy-group-v2/pkg/app"

// 关于joker处理

func (h *Hand) findJoker(wild int) {
	for _, card := range h.cards {
		if card.Value == wild || card.Suit == app.JokerA || card.Suit == app.JokerB {
			h.joker = append(h.joker, card)
		}
	}
	h.cards = h.handSliceDifference(h.cards, h.joker)
}
