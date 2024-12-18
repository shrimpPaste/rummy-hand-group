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

func (h *Hand) findSetWithJoker2(cards, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	result := make(map[int][]app.Card)
	var overCards []app.Card
	var setCards []app.Card
	// 按值分组
	for _, card := range cards {
		if result[card.Value] == nil {
			result[card.Value] = append(result[card.Value], card)
			continue
		}

		isExist := false
		for _, v := range result[card.Value] {
			if v.Suit == card.Suit {
				isExist = true
				break
			}
		}
		if isExist {
			overCards = append(overCards, card)
		} else {
			result[card.Value] = append(result[card.Value], card)
		}
	}

	// 消耗1张Joker牌
	for i, r := range result {
		if len(jokers) < 1 {
			break
		}
		if len(r) == 2 && len(jokers) >= 1 {
			setCards = append(setCards, r...)
			setCards = append(setCards, jokers[0])
			jokers = jokers[1:]
			delete(result, i)
			continue
		}
	}

	// 消耗2张Joker牌
	for i, r := range result {
		if len(jokers) < 2 {
			break
		}
		if len(r) == 1 && len(jokers) >= 2 {
			setCards = append(setCards, r...)
			setCards = append(setCards, jokers[0], jokers[1])
			jokers = jokers[2:]
			delete(result, i)
			continue
		}
	}

	for _, r := range result {
		overCards = append(overCards, r...)
	}
	return overCards, setCards, jokers
}

func (h *Hand) findSet(cards []app.Card) (overCards, setCards []app.Card, result map[int][]app.Card) {
	result = make(map[int][]app.Card)
	// 按值分组
	for _, card := range cards {
		if result[card.Value] == nil {
			result[card.Value] = append(result[card.Value], card)
			continue
		}

		isExist := false
		for _, v := range result[card.Value] {
			if v.Suit == card.Suit {
				isExist = true
				break
			}
		}
		if isExist {
			overCards = append(overCards, card)
		} else {
			result[card.Value] = append(result[card.Value], card)
		}
	}

	for i, r := range result {
		if len(r) >= 3 {
			setCards = append(setCards, r...)
			delete(result, i)
			continue
		}
		overCards = append(overCards, r...)
	}

	return overCards, setCards, result
}
