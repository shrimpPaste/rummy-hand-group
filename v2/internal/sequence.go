package internal

import (
	"rummy-group-v2/pkg/app"
	"sort"
)

// 关于找顺子的逻辑

// findSequences 找出花色中的顺子
func (h *Hand) findSequences() {
	for suit, cards := range h.suitCards {
		if len(cards) < 3 {
			h.invalid = append(h.invalid, cards...)
			h.suitCards[suit] = []app.Card{}
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		sequence := h.findValidSequence(cards)

		if len(sequence) < 3 {
			h.invalid = append(h.invalid, cards...)
		} else {
			h.valid = append(h.valid, sequence...)
		}

		h.suitCards[suit] = []app.Card{}
	}
}

// findValidSequence 找出当前花色中的顺子
func (h *Hand) findValidSequence(cards []app.Card) (result []app.Card) {
	for i := 0; i < len(cards); i++ {
		if len(result) < 3 && len(cards[i:]) >= 3 {
			result = h.findSequenceFromCards([]app.Card{}, cards[i:])
		} else {
			break
		}
	}
	return
}

// 递归数组找到连续的值
func (h *Hand) findSequenceFromCards(result, cards []app.Card) []app.Card {
	if len(cards) < 2 { // 如果剩余牌数不足，返回当前结果
		return result
	}

	// 初始化结果集（第一次调用时）
	if len(result) == 0 {
		result = append(result, cards[0])
	}

	// 检查下一张牌是否连续
	if cards[1].Value == result[len(result)-1].Value+1 {
		result = append(result, cards[1])
	}

	// 递归调用，从下一张牌开始检查
	return h.findSequenceFromCards(result, cards[1:])
}
