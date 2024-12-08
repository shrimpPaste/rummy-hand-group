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

	// todo::他只能找无效牌但是不能找到有效牌的
	setCards := h.findSetFromCards([]app.Card{}, h.invalid)
	// todo::找出一个花色中的刻子，一张牌应该怎么处理，两张牌怎么处理

	if len(h.joker) > 1 {
		h.valid = append(h.valid, setCards...)
		h.valid = append(h.valid, h.joker[0])
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
