package internal

import "rummy-logic-v3/pkg/app"

// 处理刻子的牌
//func (h *Hand) find111Cards() {
//	jokerLen := len(h.joker)
//
//	if len(h.invalid) == 1 && jokerLen >= 2 {
//		h.valid = append(h.valid, h.invalid...)
//		h.valid = append(h.valid, h.joker...)
//		h.joker = h.joker[2:]
//	}
//	mostScoreSetCards := h.findHighestScoringSet(h.invalid)
//
//	for i, p := range h.pure {
//		if len(p) <= 3 {
//			continue
//		}
//		// 给pure顺子按照花色分组
//		pureSuitCards := map[string][]app.Card{}
//		h.groupCards(pureSuitCards, p)
//
//		for _, cards := range pureSuitCards {
//			if len(mostScoreSetCards) <= 0 {
//				continue
//			}
//			tSet := app.Card{}
//			for j, p1 := range cards {
//				if p1.Value == mostScoreSetCards[0].Value && p1.Suit != mostScoreSetCards[0].Suit {
//					resPure := h.removeByIndex(cards, j)
//					// 检测切分后是否是合法的顺子
//					if h.judgeIsSeq(resPure) {
//						h.pure[i] = resPure
//						tSet = p1
//						mostScoreSetCards = append(mostScoreSetCards, tSet)
//					} else {
//						// 如果抽离后分数比较高的话也可以抽离
//						if len(pureSuitCards) >= 2 {
//							score1 := h.calculateScore(cards)
//							score2 := h.calculateScore(mostScoreSetCards)
//							if score1 < score2 {
//								// 可以抽离
//								h.pure[i] = resPure
//								tSet = p1
//								mostScoreSetCards = append(mostScoreSetCards, tSet)
//							}
//						}
//					}
//				}
//			}
//			if len(mostScoreSetCards) >= 3 {
//				h.invalid = h.handSliceDifference(h.invalid, mostScoreSetCards)
//				h.set = append(h.set, mostScoreSetCards)
//				mostScoreSetCards = []app.Card{}
//			}
//		}
//	}
//
//	for index, cards := range h.pureWithJoker {
//		if len(mostScoreSetCards) <= 0 {
//			continue
//		}
//		score1 := h.calculateScore(cards)
//		score2 := h.calculateScore(mostScoreSetCards)
//		if score1 > score2 {
//			continue
//		}
//
//		var tempJoker []app.Card
//		var isUsedCard bool
//		for _, card := range cards {
//			if len(mostScoreSetCards) <= 0 {
//				continue
//			}
//
//			if card.Value == mostScoreSetCards[0].Value {
//				isAs := false
//				for _, set := range mostScoreSetCards {
//					if set.Suit == card.Suit {
//						isAs = true
//					}
//				}
//				if !isAs {
//					score1 := h.calculateScore(cards)
//					score2 := h.calculateScore(mostScoreSetCards)
//					if score1 <= score2 {
//						// 可以抽离
//						isUsedCard = true
//						invalid := h.handSliceDifference(cards, []app.Card{card})
//						h.invalid = append(h.invalid, invalid...)
//
//						mostScoreSetCards = append(mostScoreSetCards, card)
//						h.set = append(h.set, mostScoreSetCards)
//
//						h.pureWithJoker[index] = []app.Card{}
//					}
//				}
//			}
//		}
//
//		for _, card := range cards {
//			if card.Suit == app.JokerB || card.Suit == app.JokerA || card.Value == h.wild.Value {
//				tempJoker = append(tempJoker, card)
//
//				diffCards := h.handSliceDifference(h.pureWithJoker[index], tempJoker)
//
//				if isUsedCard {
//					if len(diffCards) < 3 {
//						h.pureWithJoker[index] = nil
//						h.invalid = append(h.invalid, diffCards...)
//					} else {
//						h.pureWithJoker[index] = h.handSliceDifference(h.pureWithJoker[index], tempJoker)
//					}
//				}
//			}
//
//		}
//
//		if len(mostScoreSetCards) >= 3 {
//			h.invalid = h.handSliceDifference(h.invalid, mostScoreSetCards)
//			mostScoreSetCards = []app.Card{}
//			h.joker = append(h.joker, tempJoker...)
//		}
//	}
//
//	if len(h.joker) > 1 && len(mostScoreSetCards) >= 2 {
//		mostScoreSetCards = append(mostScoreSetCards, h.joker[0])
//		h.setWithJoker = append(h.setWithJoker, mostScoreSetCards)
//		h.joker = h.removeByIndex(h.joker, 1)
//		h.invalid = h.handSliceDifference(h.invalid, mostScoreSetCards)
//	}
//}

// findHighestScoringSet 递归找到分值最大的刻子
func (h *Hand) findHighestScoringSet(cards []app.Card) []app.Card {
	var bestSet []app.Card
	var maxScore int

	var find func(invalid []app.Card, lastSet []app.Card, lastScore int)
	find = func(invalid []app.Card, lastSet []app.Card, lastScore int) {
		if len(invalid) < 2 {
			return
		}

		currentSet := h.findSetFromCards([]app.Card{}, invalid)
		currentScore := h.calculateScore(currentSet)

		if currentScore > lastScore {
			bestSet = currentSet
			maxScore = currentScore
		} else {
			bestSet = lastSet
			maxScore = lastScore
		}

		overCards := h.handSliceDifference(invalid, currentSet)
		if len(overCards) > 0 {
			find(overCards, bestSet, maxScore)
		}
	}

	find(cards, []app.Card{}, 0)

	return bestSet
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

	if len(cards) > 0 && cards[0].Value == result[len(result)-1].Value {
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
