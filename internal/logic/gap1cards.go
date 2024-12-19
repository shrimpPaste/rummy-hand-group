package logic

import (
	"github.com/jinzhu/copier"
	"rummy-logic-v3/pkg/app"
	"sort"
)

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

// findGap 找出当前花色中的顺子
func (h *Hand) findGap(cards []app.Card) (result []app.Card) {
	for i := 0; i < len(cards); i++ {
		if len(result) < 2 && len(cards[i:]) >= 2 {
			result = h.findGapFromCards([]app.Card{}, cards[i:], false)
		}
	}
	return
}

func (h *Hand) removeByIndex(arr []app.Card, index int) []app.Card {
	if index < 0 || index >= len(arr) {
		return []app.Card{}
	}
	return append(arr[:index], arr[index+1:]...)
}

// 递归数组找到连续的值
func (h *Hand) findGapFromCards(result, cards []app.Card, usedGap2 bool) []app.Card {
	if len(cards) == 1 {
		if cards[0].Value == result[len(result)-1].Value+1 {
			result = append(result, cards...)
		}
		return result
	}
	if len(cards) < 2 { // 如果剩余牌数不足，返回当前结果
		return result
	}

	// 初始化结果集（第一次调用时）
	if len(result) == 0 {
		result = append(result, cards[0])
		cards = cards[1:]
	}

	if result[len(result)-1].Value == 1 {
		for index, card := range cards {
			// && len(cards) > 2  ? TODO::为什么当时要写大于2？
			if card.Value == 13 || card.Value == 12 {
				cards = h.removeByIndex(cards, index)
				result = append(result, card)
			}
		}
		// TODO:: 为什么这里要return出去？
		//return result
	}

	if len(cards) <= 0 {
		// 不添加会panic
		return result
	}

	// 检查下一张牌是否连续
	if cards[0].Value == result[len(result)-1].Value+1 {
		result = append(result, cards[0])
	}
	if cards[0].Value == result[len(result)-1].Value+2 && !usedGap2 {
		result = append(result, cards[0])
		usedGap2 = true
	}

	// 递归调用，从下一张牌开始检查
	return h.findGapFromCards(result, cards[1:], usedGap2)
}

func (h *Hand) findGapsByJoker(cards []app.Card, jokers []app.Card) (overCards []app.Card, pureWithJoker []app.Card, remainingJokers []app.Card) {
	// 无 Joker 的情况直接返回
	if len(jokers) == 0 {
		return cards, nil, jokers
	}

	// 分组牌型
	suitCards := make(map[string][]app.Card, 4)
	h.groupCards(suitCards, cards)

	// 使用 Joker 填补间隙
	pureWithJoker, jokers = h.fillGapsWithJokers(suitCards, jokers)

	// 剩余的牌
	for _, c := range suitCards {
		overCards = append(overCards, c...)
	}

	return overCards, pureWithJoker, jokers
}

// 使用 Joker 填补间隙逻辑
func (h *Hand) fillGapsWithJokers(suitCards map[string][]app.Card, jokers []app.Card) ([]app.Card, []app.Card) {
	var result []app.Card

	// 处理间隙为1的
	for suit, cards := range suitCards {
		var overCards, bestCombo []app.Card
		if len(cards) >= 2 {
			// 找间隙为1的
			overCards1, result1, overJokers1 := h.findBestComboGap1(cards, jokers)
			resultScore1 := h.calculateScore(result1)

			// 处理A == 14的情况
			cards2 := make([]app.Card, len(cards))
			_ = copier.Copy(&cards2, cards)

			for i := range cards2 {
				if cards2[i].Value == 1 {
					cards2[i].Value = 14
				}
			}
			overCards2, result2, overJokers2 := h.findBestComboGap1(cards2, jokers)
			resultScore2 := h.calculateScore(result2)

			if resultScore2 > resultScore1 {
				for i := range result2 {
					if result2[i].Value == 14 {
						result2[i].Value = 1
					}
				}

				for i := range overCards2 {
					if overCards2[i].Value == 14 {
						overCards2[i].Value = 1
					}
				}

				bestCombo = result2
				overCards = overCards2
				jokers = overJokers2

			} else {
				bestCombo = result1
				overCards = overCards1
				jokers = overJokers1
			}

			sort.Slice(bestCombo, func(i, j int) bool {
				return bestCombo[i].Value < bestCombo[j].Value
			})

			if len(bestCombo) >= 3 {
				result = append(result, bestCombo...)

				// 更新当前花色的剩余牌
				suitCards[suit] = overCards
			}
		}
	}

	// 处理间隙小于等于3
	for suit, cards := range suitCards {
		var overCards, bestCombo []app.Card
		// 找间隙为1的
		overCards, bestCombo, jokers = h.findBestComboGap2(cards, jokers)
		if len(bestCombo) >= 3 {
			result = append(result, bestCombo...)

			// 更新当前花色的剩余牌
			suitCards[suit] = overCards
		}
	}

	return result, jokers
}

func (h *Hand) findBestComboGap1(cards []app.Card, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	if len(jokers) <= 0 {
		return cards, nil, jokers
	}

	var result, overCards []app.Card

	singleCards := removeDuplicates(cards)
	overCards = h.handSliceDifference(cards, singleCards)
	isUsed := false

	sort.Slice(singleCards, func(i, j int) bool {
		return singleCards[i].Value > singleCards[j].Value
	})

	var tempResult []app.Card

	for i := 0; i < len(singleCards)-1; i++ {
		for j := i + 1; j < len(singleCards); j++ {
			gap := singleCards[j].Value - singleCards[i].Value
			if len(tempResult) > 0 {
				gap = singleCards[j].Value - tempResult[len(tempResult)-1].Value
			}
			if gap == -1 {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], singleCards[j])
				} else {
					tempResult = append(tempResult, singleCards[j])
				}
				break
			} else if gap == -2 && !isUsed {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], jokers[0], singleCards[j])
					i++
				} else {
					tempResult = append(tempResult, jokers[0], singleCards[j])
					i++
				}
				jokers = jokers[1:]
				isUsed = true
			} else {
				if len(tempResult) != 0 {
					overCards = append(overCards, singleCards[j])
				} else {
					overCards = append(overCards, singleCards[i])
				}
				break
			}
		}

	}

	if len(tempResult) >= 3 {
		result = append(result, tempResult...)
	}

	if len(tempResult) == 2 && len(jokers) > 0 {
		result = append(result, tempResult...)
		result = append(result, jokers[0])
		jokers = jokers[1:]
	}

	return overCards, result, jokers
}

func (h *Hand) findBestComboGap2(cards []app.Card, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	if len(jokers) <= 0 {
		return cards, nil, jokers
	}

	var result, overCards []app.Card

	singleCards := removeDuplicates(cards)
	overCards = h.handSliceDifference(cards, singleCards)

	sort.Slice(singleCards, func(i, j int) bool {
		return singleCards[i].Value > singleCards[j].Value
	})
	isUsed := false

	var tempResult []app.Card

	for i := 0; i <= len(singleCards)-1; i++ {
		if i == len(singleCards)-1 && len(tempResult) > 0 && singleCards[i].Value-tempResult[len(tempResult)-1].Value == 1 {
			tempResult = append(tempResult, singleCards[i])
		}
		if i == len(singleCards)-1 && len(tempResult) == 0 {
			overCards = append(overCards, singleCards[i])
		}
		for j := i + 1; j < len(singleCards); j++ {
			gap := singleCards[j].Value - singleCards[i].Value
			if len(tempResult) > 0 {
				gap = singleCards[j].Value - tempResult[len(tempResult)-1].Value
			}

			if gap == -1 {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], singleCards[j])
				} else {
					tempResult = append(tempResult, singleCards[j])
				}
				break
			} else if gap == -3 && len(jokers) > 1 && !isUsed {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], jokers[0], jokers[1], singleCards[j])
					i++ // 跳过当前处理的卡
				} else {
					tempResult = append(tempResult, jokers[0], jokers[1], singleCards[j])
					i++
				}
				jokers = jokers[2:]
				isUsed = true
			} else {
				if len(tempResult) != 0 {
					overCards = append(overCards, singleCards[j])
				} else {
					overCards = append(overCards, singleCards[i])
				}
				break
			}
		}
	}

	if len(tempResult) >= 3 {
		result = append(result, tempResult...)
	}

	if len(tempResult) == 2 && len(jokers) > 0 {
		result = append(result, tempResult...)
		result = append(result, jokers[0])
		jokers = jokers[1:]
	}

	return overCards, result, jokers
}
