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

	overCards := h.handSliceDifference(h.invalid, setCards)
	if len(overCards) >= 2 {
		setCards = h.findSetFromCards([]app.Card{}, overCards)
	}

	for i, p := range h.pure {
		if len(p) <= 3 {
			continue
		}
		// 给pure顺子按照花色分组
		pureSuitCards := map[string][]app.Card{}
		h.groupCards(pureSuitCards, p)

		for _, cards := range pureSuitCards {
			if len(setCards) <= 0 {
				continue
			}
			tSet := app.Card{}
			for j, p1 := range cards {
				if p1.Value == setCards[0].Value && p1.Suit != setCards[0].Suit {
					resPure := h.removeByIndex(cards, j)
					// 检测切分后是否是合法的顺子
					if h.judgeIsSeq(resPure) {
						h.pure[i] = resPure
						tSet = p1
						setCards = append(setCards, tSet)
					} else {
						// 如果抽离后分数比较高的话也可以抽离
						if len(pureSuitCards) >= 2 {
							score1 := h.calculateScore(cards)
							score2 := h.calculateScore(setCards)
							if score1 < score2 {
								// 可以抽离
								h.pure[i] = resPure
								tSet = p1
								setCards = append(setCards, tSet)
							}
						}
					}
				}
			}
			if len(setCards) >= 3 {
				h.invalid = h.handSliceDifference(h.invalid, setCards)
				h.set = append(h.set, setCards)
				setCards = []app.Card{}
			}
		}
	}
	if len(setCards) >= 3 {
		return
	}

	if len(h.joker) > 1 && len(setCards) >= 2 {
		setCards = append(setCards, h.joker[0])
		h.setWithJoker = append(h.setWithJoker, setCards)
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
