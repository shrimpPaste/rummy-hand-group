package main

import (
	"fmt"
	"sort"
)

const (
	A      string = "A"      // 黑桃
	B      string = "B"      // 红桃
	C      string = "C"      // 梅花
	D      string = "D"      // 方片
	JokerA string = "JokerA" // 大鬼
	JokerB string = "JokerB" // 小鬼
)

const (
	NotStatus int = 0 // 无状态
	HT        int = 1 // Head and Tail 缺少首尾的牌  2 3 缺少 1 4
	Bd        int = 2 // blocked // 卡隆牌 3 5 缺少 4
)

type Card struct {
	Suit  string
	Value int
}

// Hand 手牌
type Hand struct {
	cards     []Card
	joker     []Card
	valid     []Card
	invalid   []Card
	gap1Cards []Card // 间隙为1的牌
	suitCards map[string][]Card
}

type gapCard struct {
	Cards       []Card
	Status      int
	Score       int
	JokerUseNum int
	// 0: 无状态
	// 1: 1 2 4 牌型 应得到 1 2 4
	// 2: 3 5 7 牌型 应得到 5 7
}

// handSliceDifference 找两个数组之间的差集
func (h *Hand) handSliceDifference(a, b []Card) []Card {
	// 创建一个 map 来存储 b 中的元素
	bMap := make(map[Card]struct{})
	for _, card := range b {
		bMap[card] = struct{}{} // 用空结构体来表示集合中的元素
	}

	var difference []Card
	// 遍历 a 中的每个 card，检查它是否在 b 中
	for _, card := range a {
		if _, found := bMap[card]; !found {
			difference = append(difference, card) // 如果不在 b 中，就加到差集
		}
	}

	return difference
}

// findJoker 找癞子牌
func (h *Hand) findJoker(wild int) {
	for _, card := range h.cards {
		if card.Value == wild || card.Suit == JokerA || card.Suit == JokerB {
			h.joker = append(h.joker, card)
		}
	}
	h.cards = h.handSliceDifference(h.cards, h.joker)
}

// 手牌分组
func (h *Hand) groupCards(cards []Card) {
	for _, card := range cards {
		h.suitCards[card.Suit] = append(h.suitCards[card.Suit], card)
	}
}

// findSequences 找顺子
func (h *Hand) findSequences() {
	for suit, cards := range h.suitCards {
		if len(cards) < 3 {
			h.invalid = append(h.invalid, cards...)
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		var sequence []Card
		for i := 0; i < len(cards)-2; i++ {
			if cards[i+1].Value-cards[i].Value == 1 && cards[i+2].Value-cards[i+1].Value == 1 {
				sequence = append(sequence, cards[i:i+3]...)
				i += 2
			}
		}
		h.valid = append(h.valid, sequence...)
		h.suitCards[suit] = h.handSliceDifference(h.suitCards[suit], sequence)

		h.invalid = append(h.invalid, h.handSliceDifference(cards, sequence)...)
	}
}

// 判断牌型
func (h *Hand) judgeCardType() {
	fmt.Println("手牌类型判断")
}

func (h *Hand) findGap1Cards() {
	h.suitCards = make(map[string][]Card, 4)
	h.groupCards(h.invalid)
	h.invalid = []Card{}

	//fmt.Println("无效牌的分组", h.suitCards)

	gapCardBlackBoard := h.judgeMostScore(h.findGapFromS2L(), h.findGapFromL2S())

	fmt.Println()
	fmt.Println("黑板")
	for suit, gapC := range gapCardBlackBoard {
		fmt.Printf("花色 %s 的最佳牌组: ", suit)
		if gapC.Status == HT {
			// HT 牌型
			fmt.Println("首尾牌型")
			for _, card := range gapC.Cards {
				fmt.Printf("%d ", card.Value)
			}
			fmt.Println("分值", gapC.Score)
			fmt.Printf("癞子使用次数: %d\n", gapC.JokerUseNum)
			fmt.Println()
		}

	}
}

// findGapFromS2L 找间隙从小到大
func (h *Hand) findGapFromS2L() map[string]gapCard {
	blackBoard := map[string]gapCard{}

	for suit, cards := range h.suitCards {
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		gapC := gapCard{
			Status: NotStatus,
			Cards:  []Card{},
		}

		for i := 0; i < len(cards); i++ {
			if len(gapC.Cards) == 0 {
				gapC.Cards = []Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == NotStatus {
				// HT 牌型
				if cards[i].Value == cards[i-1].Value+1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = HT
					gapC.Score += cards[i].Value
				}

				// BT 牌型
				if cards[i].Value == cards[i-1].Value+2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = Bd
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == HT {
				// 因为该牌已经是从小到大了， 如果已经是 2 3 下一张应该就是5 才可能是合理的间隙牌否则就是顺子了，不应该出现在这里
				if gapC.Cards[len(gapC.Cards)-1].Value+2 == cards[i].Value && gapC.JokerUseNum == 0 {
					// 因为假设后面还有6的情况需要兼容这个情况
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}

				if gapC.Cards[len(gapC.Cards)-1].Value+1 == cards[i].Value && gapC.JokerUseNum == 1 {
					// 如果前面是 2 3 5，此时是6
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
				continue
			}
		}

		// 输出最后一组
		if len(gapC.Cards) >= 2 {
			blackBoard[suit] = gapC
		}
	}

	return blackBoard
}

// findGapFromL2S 找间隙从大到小
func (h *Hand) findGapFromL2S() map[string]gapCard {
	blackBoard := map[string]gapCard{}
	for suit, cards := range h.suitCards {
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value > cards[j].Value
		})

		gapC := gapCard{
			Status: NotStatus,
			Cards:  []Card{},
		}

		for i := 0; i < len(cards); i++ {
			if len(gapC.Cards) == 0 {
				gapC.Cards = []Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == NotStatus {
				// HT 牌型
				if cards[i].Value == cards[i-1].Value-1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = HT
					gapC.Score += cards[i].Value
				}

				// BT 牌型
				if cards[i].Value == cards[i-1].Value-2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = Bd
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == HT {
				// 因为该牌已经是从小到大了， 如果已经是 2 3 下一张应该就是5 才可能是合理的间隙牌否则就是顺子了，不应该出现在这里
				if gapC.Cards[len(gapC.Cards)-1].Value-2 == cards[i].Value && gapC.JokerUseNum == 0 {
					// 因为假设后面还有6的情况需要兼容这个情况
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}

				if gapC.Cards[len(gapC.Cards)-1].Value-1 == cards[i].Value && gapC.JokerUseNum == 1 {
					// 如果前面是 2 3 5，此时是6
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
				continue
			}
		}

		// 输出最后一组
		if len(gapC.Cards) >= 2 {
			blackBoard[suit] = gapC
		}
	}

	return blackBoard
}

// 鉴定哪一个牌型得分最高
func (h *Hand) judgeMostScore(S2L, L2S map[string]gapCard) map[string]gapCard {
	blackBoard := map[string]gapCard{}

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

func main() {
	hand := Hand{
		cards: []Card{
			{Value: 2, Suit: B},
			{Value: 3, Suit: B},

			{Value: 5, Suit: B},
			{Value: 6, Suit: B},

			{Value: 8, Suit: B},
			//{Value: 9, Suit: B},
			//{Value: 7, Suit: B},

			//{Value: 10, Suit: B},
			//{Value: 11, Suit: B},
			//{Value: 12, Suit: B},
			//{Value: 3, Suit: B},
			//
			////{Value: 5, Suit: D},
			////{Value: 5, Suit: A},
			////{Value: 5, Suit: C},
			//{Value: 1, Suit: C},
			//{Value: 3, Suit: C},
			//{Value: 4, Suit: C},
			//{Value: 14, Suit: JokerA},
			//{Value: 15, Suit: JokerB},
		},
		suitCards: make(map[string][]Card, 4),
	}
	// 找癞子牌
	hand.findJoker(0)
	//fmt.Println("总体癞子", hand.joker)

	// 分组排序
	hand.groupCards(hand.cards)
	//fmt.Println("分颜色排序", hand.suitCards)

	// 找顺子
	hand.findSequences()

	fmt.Println("未处理的牌", hand.suitCards)
	fmt.Println("有效牌", hand.valid)
	fmt.Println("无效牌", hand.invalid)
	fmt.Println("joker", hand.joker)

	// 找间隙为1的牌
	hand.findGap1Cards()
}
