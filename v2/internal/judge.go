package internal

import (
	"rummy-group-v2/pkg/app"
	"sort"
)

// 鉴定牌型是否有两个及以上的顺子
func (h *Hand) judgeIsHave2Seq() bool {
	for _, p := range h.pure {
		// 给pure顺子按照花色分组
		pureSuitCards := map[string][]app.Card{}
		h.groupCards(pureSuitCards, p)

		// 该函数调用应该在第一轮找顺子的时候判断
		if len(pureSuitCards) >= 2 {
			return true
		}
	}

	return false
}

// 鉴定牌型是否有一个及以上的顺子
func (h *Hand) judgeIsHave1Seq() bool {
	// 该函数调用应该在第一轮找顺子的时候判断
	if len(h.valid) >= 3 {
		return true
	}
	return false
}

// 鉴定哪一个牌型得分最高
func (h *Hand) judgeMostScore(S2L, L2S map[string]app.GapCard) map[string]app.GapCard {
	blackBoard := map[string]app.GapCard{}

	for suit, gapC := range S2L {
		for suit2, gapC2 := range L2S {
			if suit == suit2 {
				if gapC.Score > gapC2.Score {
					blackBoard[suit] = gapC
				} else {
					blackBoard[suit2] = gapC2
				}
			}
		}
	}
	return blackBoard
}

// 鉴定牌型是否为顺子
func (h *Hand) judgeIsSeq(cards []app.Card) bool {
	// 1. 对卡牌进行颜色分组
	suitCards := map[string][]app.Card{}
	h.groupCards(suitCards, cards)

	for _, cards := range suitCards {
		if len(cards) < 3 {
			// judge:: 小于3一定不是合法的顺子
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		sequence := h.findValidSequence(cards)

		if len(sequence) < 3 {
			continue
		}
		return true
	}
	return false
}

// 鉴定牌型是否为顺子
func (h *Hand) judgeIsValidSeq(cards []app.Card) bool {
	if len(cards) < 3 {
		return false
	}
	// 7 8 joker
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	var jokers []app.Card

	for _, card := range cards {
		if card.Suit == app.JokerA || card.Suit == app.JokerB || card.Value == h.wild.Value {
			jokers = append(jokers, card)
			cards = h.handSliceDifference(cards, jokers)
			continue
		}
	}

	gapValue := 0
	for i, card := range cards {
		if i == len(cards)-1 {
			break
		}

		gapV := 0
		if card.Value == 1 && cards[i+1].Value != 2 {
			//   1 9 10 11 12 13 的牌型
			gapV = 14 - cards[len(cards)-1].Value
			continue
		}
		gapV = card.Value - cards[i+1].Value
		if gapV != 1 && gapV != -1 {
			gapValue += gapV
		}
	}
	if len(jokers)*2 >= gapValue*-1 {
		return true
	}
	return false
}
