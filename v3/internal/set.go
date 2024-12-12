package internal

import "rummy-logic-v3/pkg/app"

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

func (h *Hand) findSetWithJoker(cards, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	var valid []app.Card

	if len(jokers) == 0 {
		return cards, valid, jokers
	}

	for _, card := range cards {
		if len(valid) == 0 {
			valid = append(valid, card)
			continue
		}

		for _, v := range valid {
			if card.Value == v.Value && card.Suit != v.Suit {
				valid = append(valid, card)
			}
		}
	}

	if len(valid) == 2 && len(jokers) >= 1 {
		// 消耗一张joker牌

		jokers = jokers[1:]
		valid = append(valid, jokers[0])
		cards = h.handSliceDifference(cards, valid)

		return cards, valid, jokers
	}
	return cards, valid, jokers
}
