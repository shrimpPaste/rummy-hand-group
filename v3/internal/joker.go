package internal

import "rummy-logic-v3/pkg/app"

// 关于joker处理

func (h *Hand) findJoker(cards []app.Card) (jokers []app.Card, overCards []app.Card) {
	// 移除joker牌
	for _, card := range cards {
		if card.Suit == app.JokerA || card.Suit == app.JokerB || card.Value == h.wild.Value {
			// 添加joker
			jokers = append(jokers, card)
		} else {
			overCards = append(overCards, card)
		}
	}
	return jokers, overCards
}
