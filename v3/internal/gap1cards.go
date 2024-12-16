package internal

import (
	"rummy-logic-v3/pkg/app"
	"sort"
)

// findValidGap1Cards 处理纯顺子中也有可能有卡隆牌
//func (h *Hand) findValidGap1Cards() {
//	// 检测是否已经拥有两个及以上的顺子，并且没有使用过小丑牌。
//	//suitCards := make(map[string][]app.Card, 4)
//	//
//	//for _, cards := range h.pure {
//	//	// 给pure顺子按照花色分组
//	//	h.groupCards(suitCards, cards)
//	//}
//	//
//	//gapScore := map[int][]app.Card{}
//	//
//	//for _, cards := range suitCards {
//	//	//todo:: 去找无效牌 有没有可能和有效牌组合成间隙
//	//	// 没有可能就跳过
//	//
//	//	gapScore = h.handleGapsCards(cards, gapScore)
//	//}
//	//
//	//for _, joker := range h.joker {
//	//	if len(gapScore) == 1 {
//	//		continue
//	//	}
//	//	bestCards, g := h.findAndRemoveMaxGapScore(gapScore)
//	//
//	//	if len(bestCards) > 0 {
//	//		for i := range h.pure {
//	//			h.pure[i] = h.handSliceDifference(h.pure[i], bestCards)
//	//		}
//	//		bestCards = append(bestCards, joker)
//	//
//	//		fmt.Println("找到顺子", bestCards)
//	//		h.pureWithJoker = append(h.pureWithJoker, bestCards)
//	//
//	//		h.joker = h.removeByIndex(h.joker, 0)
//	//		gapScore = g
//	//	}
//	//}
//
//	suitCards := make(map[string][]app.Card, 4)
//	h.groupCards(suitCards, h.GetCards())
//
//	// 移除joker牌
//	for suit, cards := range suitCards {
//		for _, card := range cards {
//			if card.Suit == app.JokerA || card.Suit == app.JokerB || card.Value == h.wild.Value {
//				// 添加joker
//				h.joker = append(h.joker, card)
//
//				cards = h.handSliceDifference(cards, []app.Card{card})
//				suitCards[suit] = cards
//			}
//		}
//	}
//
//	var invalid []app.Card
//	var valid []app.Card
//
//	for _, cards := range suitCards {
//		sort.Slice(cards, func(i, j int) bool {
//			return cards[i].Value < cards[j].Value
//		})
//		sequence := h.findValidSequence(cards)
//
//		if len(sequence) < 3 {
//			invalid = append(invalid, cards...)
//		} else {
//			valid = append(valid, sequence...)
//			invalid = append(invalid, h.handSliceDifference(cards, sequence)...)
//		}
//	}
//
//	//joker [{A 7} {D 7}]
//	//有效牌 [{B 9} {B 10} {B 11} {A 1} {A 12} {A 13}]
//	//无效牌 [{C 9} {B 4} {B 6} {B 13} {A 8}]
//
//	if valid == nil {
//		h.invalid = h.GetCards()
//		return
//	}
//
//	if len(valid) >= 6 {
//		// 有效牌分组
//		validSuitCards := make(map[string][]app.Card, 4)
//		h.groupCards(validSuitCards, valid)
//
//		jokerLength := len(h.GetJoker())
//
//		// 找间隙牌
//		for suit, cards := range validSuitCards {
//			sort.Slice(cards, func(i, j int) bool {
//				return cards[i].Value < cards[j].Value
//			})
//			findKeyValue := cards[0].Value
//			findEndValue := cards[len(cards)-1].Value
//
//			var invalidValue app.Card
//			if findKeyValue == 1 && cards[1].Value == 2 {
//				// 说明他是 1 2 3的顺子
//				continue
//			}
//			if findKeyValue == 1 && cards[1].Value != 2 {
//				// 说明他是 10 11 12 13 1
//				// 或者是 1 12 13的牌型
//				findKeyValue = cards[1].Value
//			}
//
//			for i, v := range invalid {
//				if jokerLength <= 0 {
//					break
//				}
//				if v.Suit != suit {
//					continue
//				}
//
//				if findKeyValue-2 == v.Value {
//					invalidValue = v
//					validSuitCards[suit] = append(validSuitCards[suit], invalidValue)
//					invalid = h.removeByIndex(invalid, i)
//					jokerLength--
//					// 消耗joker
//					validSuitCards[suit] = append(validSuitCards[suit], h.joker[0])
//					h.joker = h.joker[1:]
//				}
//
//				if findEndValue+2 == v.Value {
//					invalidValue = v
//					validSuitCards[suit] = append(validSuitCards[suit], invalidValue)
//					invalid = h.removeByIndex(invalid, i)
//					jokerLength--
//					// 消耗joker
//					validSuitCards[suit] = append(validSuitCards[suit], h.joker[0])
//					h.joker = h.joker[1:]
//				}
//
//				if invalidValue.Value-1 == v.Value {
//					invalidValue = v
//					validSuitCards[suit] = append(validSuitCards[suit], invalidValue)
//					invalid = h.removeByIndex(invalid, i)
//				}
//
//				if invalidValue.Value+1 == v.Value {
//					invalidValue = v
//					validSuitCards[suit] = append(validSuitCards[suit], invalidValue)
//					invalid = h.removeByIndex(invalid, i)
//				}
//			}
//		}
//
//		for _, cards := range validSuitCards {
//			isHaveJoker := false
//			for _, card := range cards {
//				if card.Suit == app.JokerA || card.Suit == app.JokerB || card.Value == h.wild.Value {
//					isHaveJoker = true
//				}
//			}
//			if isHaveJoker {
//				h.pureWithJoker = append(h.pureWithJoker, cards)
//			} else {
//				h.pure = append(h.pure, cards)
//			}
//		}
//	} else {
//		h.pure = append(h.pure, valid)
//	}
//
//	// 从无效牌中找 间隙牌 + joker = 带joker的顺子
//	invalidSuitCards := make(map[string][]app.Card, 4)
//	h.groupCards(invalidSuitCards, invalid)
//	gapScore := map[int][]app.Card{}
//	invalid = []app.Card{}
//
//	for suit, cards := range invalidSuitCards {
//		if len(cards) < 2 {
//			invalid = append(invalid, cards...)
//			delete(invalidSuitCards, suit)
//			continue
//		}
//		sort.Slice(cards, func(i, j int) bool {
//			return cards[i].Value < cards[j].Value
//		})
//
//		gapScore = h.handleGapsCards(cards, gapScore)
//		for _, g := range gapScore {
//			invalidSuitCards[suit] = h.handSliceDifference(invalidSuitCards[suit], g)
//		}
//	}
//
//	for _, joker := range h.joker {
//		bestCards, g := h.findAndRemoveMaxGapScore(gapScore)
//		if len(bestCards) > 0 {
//			bestCards = append(bestCards, joker)
//
//			h.pureWithJoker = append(h.pureWithJoker, bestCards)
//
//			invalid = h.handSliceDifference(invalid, bestCards)
//
//			h.joker = h.removeByIndex(h.joker, 0)
//			gapScore = g
//		}
//	}
//
//	for _, cards := range gapScore {
//		invalid = append(invalid, cards...)
//	}
//	for _, cards := range invalidSuitCards {
//		invalid = append(invalid, cards...)
//	}
//
//	// 牌型没有一个纯顺子 + 第二个顺子就要结束
//	seqNum := 0
//	seqNum += len(h.pure)
//	seqNum += len(h.pureWithJoker)
//
//	if seqNum < 2 {
//		for _, p := range h.pure {
//			if len(p) >= 6 {
//				seqNum += 2
//			}
//		}
//		h.invalid = invalid
//		return
//	}
//
//	// 找刻子
//
//	// 第一回
//	// 1. 先找无效牌中 3张value相等 suit不同的值并且添加到 set中
//
//	// 第二回
//	// 1. 先找无效牌中 3张value相等 suit不同的值添加到 临时set中
//	// 2. 去找纯顺子且纯顺子不能小于1 且被取出后不能影响他是纯顺子的逻辑
//	// 3. 如果纯顺子找不到就去找 带joker的顺子，移除该牌后，计算总顺子数量是否 > 2，且如果移除后的分数要大于未被移除的分数否则不处理
//
//	// 按值分组
//	groupCard := make(map[int][]app.Card)
//	for _, card := range invalid {
//		// 如果不存在该值，初始化空切片
//		if _, exists := groupCard[card.Value]; !exists {
//			groupCard[card.Value] = []app.Card{}
//		}
//
//		// 检查当前值的分组中是否已经存在相同花色的牌
//		alreadyExists := false
//		for _, existingCard := range groupCard[card.Value] {
//			if existingCard.Suit == card.Suit {
//				alreadyExists = true
//				break
//			}
//		}
//
//		// 如果没有相同花色的牌，才追加到分组中
//		if !alreadyExists {
//			groupCard[card.Value] = append(groupCard[card.Value], card)
//		}
//	}
//
//	// 第一回
//	// 1. 先找无效牌中 3张value相等 suit不同的值并且添加到 set中
//	for key, cards := range groupCard {
//		if len(cards) >= 3 {
//			h.set = append(h.set, cards)
//			invalid = h.handSliceDifference(invalid, cards)
//			delete(groupCard, key)
//		}
//	}
//
//	// 第二回
//	// 1. 先找无效牌中 3张value相等 suit不同的值添加到 临时set中
//	// 2. 去找纯顺子且纯顺子不能小于1 且被取出后不能影响他是纯顺子的逻辑
//	// 3. 如果纯顺子找不到就去找 带joker的顺子，移除该牌后，计算总顺子数量是否 > 2，且如果移除后的分数要大于未被移除的分数否则不处理
//	for key, cards := range groupCard {
//		if len(cards) != 2 {
//			continue
//		}
//
//		for i, p := range h.pure {
//			if len(h.pure) == 1 && len(p) == 3 {
//				// 只有一个纯顺子且只有三张，那就不能抽
//				break
//			}
//			// 剩下的都是大于三张的，这里先只处理首尾可能是 刻子的情况
//			first := p[0]
//			last := p[len(p)-1]
//			for _, card := range cards {
//				if card.Value == first.Value && card.Suit != first.Suit {
//					invalid = h.handSliceDifference(invalid, cards)
//					h.pure[i] = h.handSliceDifference(h.pure[i], []app.Card{first})
//					cards = append(cards, first)
//					h.set = append(h.set, cards)
//					delete(groupCard, key)
//					break
//				}
//				if card.Value == last.Value && card.Suit != last.Suit {
//					invalid = h.handSliceDifference(invalid, cards)
//					h.pure[i] = h.handSliceDifference(h.pure[i], []app.Card{last})
//					cards = append(cards, last)
//					h.set = append(h.set, cards)
//					delete(groupCard, key)
//					break
//				}
//			}
//		}
//
//		for j, p := range h.pureWithJoker {
//			//if len(h.pure) == 1 && len(h.pureWithJoker) == 1 && len(p) == 3 {
//			//	// 只有两个顺子，不能抽
//			//	break
//			//}
//
//			first := p[0]
//			last := p[len(p)-1]
//			for i, card := range cards {
//				if card.Value == first.Value && card.Suit != first.Suit {
//					if h.judgeIsValidSeq(p[i+1:]) {
//						h.pureWithJoker[j] = h.handSliceDifference(h.pureWithJoker[j], cards[i+1:])
//					} else {
//						//var tJoker []app.Card
//						invalidAAA := append(invalid, p[i+1:]...)
//						invalidAAA = h.handSliceDifference(invalidAAA, cards)
//						for _, v := range invalidAAA {
//							if v.Suit == app.JokerA || v.Suit == app.JokerB || v.Value == h.wild.Value {
//								h.joker = append(h.joker, v)
//								invalidAAA = h.handSliceDifference(invalidAAA, []app.Card{v})
//								continue
//							}
//						}
//						h.pureWithJoker = append(h.pureWithJoker[:j], h.pureWithJoker[j+1:]...)
//						invalid = h.findGap111(invalidAAA)
//					}
//					invalid = h.handSliceDifference(invalid, cards)
//					cards = append(cards, first)
//					h.set = append(h.set, cards)
//					break
//				}
//				if card.Value == last.Value && card.Suit != last.Suit {
//					if h.judgeIsValidSeq(p[i:]) {
//						h.pureWithJoker[j] = h.handSliceDifference(h.pureWithJoker[j], cards[i:])
//					} else {
//						invalidAAA := append(invalid, p[i:]...)
//						fmt.Println("invalidAAA", invalidAAA)
//					}
//					cards = append(cards, first)
//					fmt.Println("2222222", cards)
//					h.set = append(h.set, cards)
//					break
//				}
//			}
//		}
//	}
//
//	// TODO：：处理特殊情况，如果 无效牌中的间隙牌分值 大于 带joker的顺子的分值，则需要替换他，因为我们是先处理有效牌中的顺子，可能会出现把joker用完的情况
//
//	//todo::h.invalid = h.fixInvalidScoreBigPureWithJokerScore(h.pureWithJoker, invalid)
//	h.invalid = invalid
//
//	//fmt.Println("带joker的顺子", h.pureWithJoker)
//	//fmt.Println("纯顺子", h.pure)
//	//fmt.Println("纯刻子", h.set)
//	//fmt.Println("带joker的刻子", h.setWithJoker)
//	//fmt.Println("无效牌", h.invalid)
//	//fmt.Println("joker", h.joker)
//}
//
//func (h *Hand) findGap111(invalid []app.Card) []app.Card {
//	invalidSuitCards := make(map[string][]app.Card, 4)
//	h.groupCards(invalidSuitCards, invalid)
//	gapScore := map[int][]app.Card{}
//	invalid = []app.Card{}
//
//	for suit, cards := range invalidSuitCards {
//		if len(cards) < 2 {
//			invalid = append(invalid, cards...)
//			delete(invalidSuitCards, suit)
//			continue
//		}
//		sort.Slice(cards, func(i, j int) bool {
//			return cards[i].Value < cards[j].Value
//		})
//
//		gapScore = h.handleGapsCards(cards, gapScore)
//		for _, g := range gapScore {
//			invalidSuitCards[suit] = h.handSliceDifference(invalidSuitCards[suit], g)
//		}
//	}
//
//	for _, joker := range h.joker {
//		bestCards, g := h.findAndRemoveMaxGapScore(gapScore)
//		if len(bestCards) > 0 {
//			bestCards = append(bestCards, joker)
//
//			h.pureWithJoker = append(h.pureWithJoker, bestCards)
//
//			invalid = h.handSliceDifference(invalid, bestCards)
//
//			h.joker = h.removeByIndex(h.joker, 0)
//			gapScore = g
//		}
//	}
//
//	for _, cards := range gapScore {
//		invalid = append(invalid, cards...)
//	}
//	for _, cards := range invalidSuitCards {
//		invalid = append(invalid, cards...)
//	}
//	return invalid
//}
//
//func (h *Hand) fixInvalidScoreBigPureWithJokerScore(pureWithJoker [][]app.Card, invalid []app.Card) []app.Card {
//	// 分组 invalidSuitCards
//	invalidSuitCards := make(map[string][]app.Card, 4)
//	h.groupCards(invalidSuitCards, invalid)
//
//	// 记录 invalid 的 gap 分数
//	invalidGapScore := map[int][]app.Card{}
//	invalid = []app.Card{}
//
//	// 处理 invalidSuitCards 分组
//	for suit, cards := range invalidSuitCards {
//		if len(cards) < 2 {
//			invalid = append(invalid, cards...)
//			delete(invalidSuitCards, suit)
//			continue
//		}
//		sort.Slice(cards, func(i, j int) bool {
//			return cards[i].Value < cards[j].Value
//		})
//
//		invalidGapScore = h.handleGapsCards(cards, invalidGapScore)
//		for _, g := range invalidGapScore {
//			invalidSuitCards[suit] = h.handSliceDifference(invalidSuitCards[suit], g)
//		}
//	}
//
//	// 输出 invalidGapScore
//	invalidCards, score := h.getMostScoreCards(invalidGapScore)
//
//	pureCards, delCard, tSeqCards, tJoker, score2 := h.getMinScoreGroup(pureWithJoker)
//
//	sequence := h.findValidSequence(pureCards)
//	unPureCards := h.handSliceDifference(pureCards, sequence)
//	pureCards = h.handSliceIntersection(pureCards, sequence)
//
//	if score2 > score {
//		for _, p := range pureWithJoker {
//			h.joker = h.handSliceDifference(h.joker, p)
//		}
//	} else {
//		for _, joker := range tJoker {
//			if len(invalidCards) > 0 {
//				invalidCards = append(invalidCards, joker)
//
//				for i, pwj := range h.pureWithJoker {
//					if len(pwj) > 0 && len(tSeqCards) > 0 && pwj[0].Suit == tSeqCards[0].Suit {
//						h.pureWithJoker[i] = h.handSliceDifference(h.pureWithJoker[i], delCard)
//						invalid = h.handSliceDifference(invalid, invalidCards)
//					}
//				}
//
//				h.pureWithJoker = append(h.pureWithJoker, invalidCards)
//				if len(tSeqCards) > 0 {
//					h.pure = append(h.pure, tSeqCards)
//					invalid = append(invalid, unPureCards...)
//				}
//
//				invalid = h.handSliceDifference(invalid, invalidCards)
//				tJoker = h.removeByIndex(tJoker, 0)
//			}
//		}
//
//		//invalid = append(invalid, pureCards...)
//	}
//	return invalid
//}

// 通过map 返回int值最高的
func (h *Hand) getMostScoreCards(cardsMap map[int][]app.Card) ([]app.Card, int) {
	// 内部函数，用于计算最大分数
	calcMostScore := func(cardsMap map[int][]app.Card) int {
		maxScore := 0
		for k := range cardsMap {
			if k > maxScore {
				maxScore = k
			}
		}
		return maxScore
	}

	// 调用内部函数计算最大分数
	mostScore := calcMostScore(cardsMap)
	// 返回结果
	return cardsMap[mostScore], mostScore
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

func (h *Hand) findGapMostScoreCards(overCards, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	// TODO:: 当癞子变多，找不到带joker的顺子则一个都找不到
	// TODO:: 这里只能找到间隙为1的，找间隙牌应该再实现一个不区分癞子的，把他都找出来，假设癞子是红桃6，同时你有两个红桃6，一个红桃8
	// TODO:: 那么你应该能够得到 68 6这样子的一个带joker的顺子 进入这种情况是要不满足两个顺子的情况。
	suitCards := make(map[string][]app.Card, 4)
	h.groupCards(suitCards, overCards)
	gapScore := map[int][]app.Card{}

	var result []app.Card

	for suit, cards := range suitCards {
		if len(cards) < 2 {
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		gapScore = h.handleGapsCards(cards, gapScore)
		for _, g := range gapScore {
			suitCards[suit] = h.handSliceDifference(suitCards[suit], g)
		}
	}

	for _, joker := range jokers {
		bestCards, g := h.findAndRemoveMaxGapScore(gapScore)
		if len(bestCards) > 0 {
			bestCards = append(bestCards, joker)

			result = append(result, bestCards...)

			overCards = h.handSliceDifference(overCards, bestCards)

			jokers = h.handSliceDifference(jokers, bestCards)
			//jokers = h.removeByIndex(jokers, i)
			gapScore = g
		}
	}

	//for _, cards := range gapScore {
	//	overCards = append(overCards, cards...)
	//}

	return overCards, result, jokers
}

func (h *Hand) handleGapsCards(cards []app.Card, gapScore map[int][]app.Card) (score map[int][]app.Card) {
	gapsCards := h.findGap(cards)

	if len(gapsCards) >= 2 {
		gapScore[h.calculateScore(gapsCards)] = gapsCards
	}

	if len(gapsCards) < 2 {
		return gapScore
	}

	overCards := h.handSliceDifference(cards, gapsCards)
	gapScore = h.handleGapsCards(overCards, gapScore)

	return gapScore
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
			overCards, bestCombo, jokers = h.findBestComboGap1(cards, jokers)
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
	isUsed := false
	// 间隙为1的
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if i > len(cards)-1 {
				break
			}
			gap := cards[j].Value - cards[i].Value
			if gap == 0 {
				overCards = append(overCards, cards[i])
				if len(result) > 0 {
					result = result[:len(result)-1]
				}

				continue
			} else if gap == 1 {
				if len(result) == 0 {
					result = append(result, cards[i], cards[j])
				} else {
					result = append(result, cards[j])
				}
				i++
				continue
			} else if gap == 2 && !isUsed {
				if len(result) >= 2 && result[len(result)-1].Value == cards[i].Value {
					result = append(result, jokers[0])
					result = append(result, cards[j])
				} else {
					result = append(result, cards[i])
					result = append(result, jokers[0])
					result = append(result, cards[j])
				}

				jokers = jokers[1:]
				i = j
				isUsed = true
			} else {
				if len(result) == 2 && len(jokers) >= 1 {
					result = append(result, jokers[0])
					jokers = jokers[1:]
					overCards = append(overCards, cards[i+1:]...)
					i = len(cards)
				} else {
					overCards = append(overCards, cards[i])
					i++
					break
				}
			}
		}
	}

	if len(result) == 2 && len(jokers) >= 1 {
		result = append(result, jokers[0])
		jokers = jokers[1:]
	}

	if len(overCards) >= 2 && len(jokers) >= 1 {
		overCards2, result2, jokers2 := h.findBestComboGap1(overCards, jokers)
		if len(result2) >= 3 {
			result = append(result, result2...)
			jokers = jokers2
			overCards = overCards2
		}
	}
	return overCards, result, jokers
}

func (h *Hand) findBestComboGap2(cards []app.Card, jokers []app.Card) ([]app.Card, []app.Card, []app.Card) {
	if len(jokers) <= 0 {
		return cards, nil, jokers
	}

	var result, overCards []app.Card
	isUsed := false
	// 间隙为1的
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			gap := cards[j].Value - cards[i].Value
			if gap == 0 {
				overCards = append(overCards, cards[i])
				i++
				continue
			} else if gap == 1 {
				result = append(result, cards[i], cards[j])
				i++
				continue
			} else if gap == 3 && !isUsed {
				if len(result) == 0 {
					result = append(result, cards[i])
				}
				result = append(result, jokers[0], jokers[1])
				result = append(result, cards[j])

				jokers = jokers[2:]
				i = j
				isUsed = true
			} else {
				overCards = append(overCards, cards[i])
				i++
				break
			}
		}
	}

	if len(overCards) >= 2 && len(jokers) >= 2 {
		overCards2, result2, jokers2 := h.findBestComboGap2(overCards, jokers)
		if len(result2) >= 3 {
			result = append(result, result2...)
			jokers = jokers2
			overCards = overCards2
		}
	}

	return overCards, result, jokers
}
