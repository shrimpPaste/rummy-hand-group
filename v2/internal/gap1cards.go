package internal

import (
	"rummy-group-v2/pkg/app"
	"sort"
)

// 处理卡隆牌
func (h *Hand) findGap1Cards() {
	// 检测是否已经拥有两个及以上的顺子，并且没有使用过小丑牌。
	h.suitCards = make(map[string][]app.Card, 4)
	h.groupCards(h.suitCards, h.invalid)
	h.invalid = []app.Card{}

	gapScore := map[int][]app.Card{}

	for suit, cards := range h.suitCards {
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			h.suitCards[suit] = []app.Card{}
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		gapsCards := h.findGap(cards)
		h.getCardsScore(gapsCards)
		gapScore[h.getCardsScore(gapsCards)] = gapsCards

		//sequence := h.findValidSequence(cards)
		//
		//h.valid = append(h.valid, sequence...)
		h.invalid = append(h.invalid, h.handSliceDifference(cards, gapsCards)...)
		//
		//h.suitCards[suit] = []app.Card{}
	}
	for _, joker := range h.joker {
		bestCards, g := h.findAndRemoveMaxGapScore(gapScore)
		h.valid = append(h.valid, bestCards...)
		h.valid = append(h.valid, joker)
		h.joker = h.joker[1:]
		gapScore = g
	}

	for _, cards := range gapScore {
		h.invalid = append(h.invalid, cards...)
	}
}

func (h *Hand) findAndRemoveMaxGapScore(gapScore map[int][]app.Card) ([]app.Card, map[int][]app.Card) {
	var maxKey int
	var maxCards []app.Card

	// 遍历 map，找到最大键值
	for key, cards := range gapScore {
		if key > maxKey {
			maxKey = key
			maxCards = cards
		}
	}

	// 删除最大键值对应的 entry
	delete(gapScore, maxKey)

	return maxCards, gapScore
}

// getCardsScore 获取当前卡组分数
func (h *Hand) getCardsScore(cards []app.Card) int {
	score := 0
	for _, gap := range cards {
		if gap.Value > 10 || gap.Value == 1 {
			score += 10
			continue
		}
		score += gap.Value
	}
	return score
}

// findValidSequence 找出当前花色中的顺子
func (h *Hand) findGap(cards []app.Card) (result []app.Card) {
	for i := 0; i < len(cards); i++ {
		if len(result) < 2 && len(cards[i:]) >= 2 {
			result = h.findGapFromCards([]app.Card{}, cards[i:], false)
		}
	}
	return
}

// 递归数组找到连续的值
func (h *Hand) findGapFromCards(result, cards []app.Card, usedGap2 bool) []app.Card {
	if len(cards) < 2 { // 如果剩余牌数不足，返回当前结果
		return result
	}

	// 初始化结果集（第一次调用时）
	if len(result) == 0 {
		result = append(result, cards[0])
	}

	if result[len(result)-1].Value == 1 && cards[1].Value == 13 {
		result = append(result, cards[1])
	}
	if result[len(result)-1].Value == 1 && cards[1].Value == 12 {
		result = append(result, cards[1])
		usedGap2 = true
	}

	// 检查下一张牌是否连续
	if cards[1].Value == result[len(result)-1].Value+1 {
		result = append(result, cards[1])
	}
	if cards[1].Value == result[len(result)-1].Value+2 && !usedGap2 {
		result = append(result, cards[1])
		usedGap2 = true
	}

	// 递归调用，从下一张牌开始检查
	return h.findGapFromCards(result, cards[1:], usedGap2)
}
