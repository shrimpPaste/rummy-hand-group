package internal

import (
	"rummy-logic-v3/pkg/app"
	"sort"
)

// 关于找顺子的逻辑

// findSequences 找出花色中的顺子
func (h *Hand) findSequences() {
	//for suit, cards := range h.suitCards {
	//	if len(cards) < 3 {
	//		h.invalid = append(h.invalid, cards...)
	//		h.suitCards[suit] = []app.Card{}
	//		continue
	//	}
	//	sort.Slice(cards, func(i, j int) bool {
	//		return cards[i].Value < cards[j].Value
	//	})
	//
	//	sequence := h.findValidSequence(cards)
	//
	//	if len(sequence) < 3 {
	//		h.invalid = append(h.invalid, cards...)
	//	} else {
	//		h.valid = append(h.valid, sequence...)
	//		h.invalid = append(h.invalid, h.handSliceDifference(cards, sequence)...)
	//	}
	//
	//	h.suitCards[suit] = []app.Card{}
	//}
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

	if result[len(result)-1].Value == 1 {
		for _, card := range cards {
			if card.Value == 13 || card.Value == 12 {
				isExist := false
				for _, r := range result {
					if r.Value == 13 || r.Value == 12 {
						isExist = true
					}
				}
				if !isExist {
					result = append(result, card)
				} else {
					continue
				}
			}
		}
		if len(result) >= 3 {
			cards = h.handSliceDifference(cards, result)
		}
		if len(cards) < 2 {
			return result
		}
		return h.findSequenceFromCards(result, cards[1:])
	}

	// 检查下一张牌是否连续
	if cards[1].Value == result[len(result)-1].Value+1 {
		result = append(result, cards[1])
	}

	// 递归调用，从下一张牌开始检查
	return h.findSequenceFromCards(result, cards[1:])
}

// GetPure 获取纯顺子
func (h *Hand) GetPure(cards []app.Card) (pureCards, overCards []app.Card) {
	suitCards := make(map[string][]app.Card, 4)
	h.groupCards(suitCards, cards)

	for _, c := range suitCards {
		sort.Slice(c, func(i, j int) bool {
			return c[i].Value < c[j].Value
		})

		sequence := h.findBestSeqLogic(c)

		if len(sequence) < 3 {
			overCards = append(overCards, c...)
		} else {
			pureCards = append(pureCards, sequence...)
			overCards = append(overCards, h.handSliceDifference(c, sequence)...)
		}
	}

	return
}

func (h *Hand) findBestSeqLogic(cards []app.Card) (pureCards []app.Card) {
	seqCards := append([]app.Card{}, cards...)
	valid1, _, score1 := findBestSequence1(seqCards)
	valid2, _, score2 := findBestSequence2(seqCards)

	if score1 > score2 {
		pureCards = valid1
	} else {
		pureCards = valid2
	}

	return pureCards
}

func findBestSequence1(cards []app.Card) (valid, invalid []app.Card, score int) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	calculateScore := func(seq []app.Card) int {
		s := 0
		for i := range seq {
			if seq[i].Value > 10 || seq[i].Value == 1 {
				s += 10
			} else {
				s += seq[i].Value
			}
		}
		return s
	}

	seq := []app.Card{cards[0]}

	for i := 1; i < len(cards); i++ {
		if seq[len(seq)-1].Value+1 == cards[i].Value {
			seq = append(seq, cards[i])
		} else {
			invalid = append(invalid, cards[i])
		}
	}

	if len(seq) >= 3 {
		valid = seq
		score += calculateScore(valid)

		if len(invalid) >= 3 {
			valid2, invalid2, score2 := findBestSequence1(invalid)
			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[app.Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []app.Card
				for _, card := range invalid {
					if !valid2Map[card] {
						remainingInvalid = append(remainingInvalid, card)
					}
				}
				invalid = remainingInvalid
			}
			if len(invalid2) > 0 {
				invalid = append(invalid, invalid2...)
			}
		}
	} else {
		if len(invalid) >= 3 {
			valid2, invalid2, score2 := findBestSequence1(invalid)
			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[app.Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []app.Card
				for _, card := range invalid {
					if !valid2Map[card] {
						remainingInvalid = append(remainingInvalid, card)
					}
				}
				invalid = remainingInvalid
			}
			if len(invalid2) > 0 {
				invalid = append(invalid, invalid2...)
			}
		}
		invalid = append(invalid, seq...)
	}

	return valid, invalid, score
}

func findBestSequence2(cards []app.Card) (valid, invalid []app.Card, score int) {
	// 将所有的1替换为14，便于处理A作为最大值的情况
	for i := range cards {
		if cards[i].Value == 1 {
			cards[i].Value = 14
		}
	}

	// 按Value从大到小排序
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value > cards[j].Value
	})

	// 用于存储有效序列
	seq := []app.Card{cards[0]}
	cards = cards[1:]

	for i := 0; i < len(cards); i++ {
		if seq[len(seq)-1].Value-1 == cards[i].Value {
			seq = append(seq, cards[i])
		} else {
			invalid = append(invalid, cards[i])
		}
	}

	if len(seq) >= 3 {
		valid = seq

		// 对invalid中的卡片递归查找序列
		if len(invalid) >= 3 {
			valid2, invalid2, _ := findBestSequence2(invalid)

			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				//score += score2

				valid2Map := make(map[app.Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []app.Card
				for _, card := range invalid {
					if !valid2Map[card] {
						remainingInvalid = append(remainingInvalid, card)
					}
				}
				invalid = remainingInvalid
			}

			if len(invalid2) > 0 {
				invalid = invalid2
			}
		}
	} else {
		if len(invalid) >= 3 {
			valid2, invalid2, _ := findBestSequence2(invalid)

			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				//score += score2

				valid2Map := make(map[app.Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []app.Card
				for _, card := range invalid {
					if !valid2Map[card] {
						remainingInvalid = append(remainingInvalid, card)
					}
				}
				invalid = remainingInvalid
			}

			if len(invalid2) > 0 {
				invalid = invalid2
			}
		}
		invalid = append(invalid, seq...)
	}

	// 计算分数，并恢复A的值
	for i := range valid {
		if valid[i].Value > 10 {
			score += 10
		} else {
			score += valid[i].Value
		}
		if valid[i].Value == 14 {
			valid[i].Value = 1
		}
	}

	for i := range invalid {
		if invalid[i].Value == 14 {
			invalid[i].Value = 1
		}
	}
	return valid, invalid, score
}
