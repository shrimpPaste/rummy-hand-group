package internal

import (
	"fmt"
	"rummy-group-v2/pkg/app"
	"sort"
)

// Hand 手牌
type Hand struct {
	cards     []app.Card
	joker     []app.Card
	valid     []app.Card
	invalid   []app.Card
	gap1Cards []app.Card // 间隙为1的牌
	suitCards map[string][]app.Card
}

// initHand 初始化手牌
func (h *Hand) initHand() {
	h.cards = []app.Card{
		{Value: 4, Suit: app.B},
		{Value: 5, Suit: app.B},
		{Value: 7, Suit: app.B},
		{Value: 9, Suit: app.B},
		{Value: 11, Suit: app.B},
	}
}

func (h *Hand) findJoker(wild int) {
	for _, card := range h.cards {
		if card.Value == wild || card.Suit == app.JokerA || card.Suit == app.JokerB {
			h.joker = append(h.joker, card)
		}
	}
	h.cards = h.handSliceDifference(h.cards, h.joker)
}

// handSliceDifference 找两个数组之间的差集
func (h *Hand) handSliceDifference(a, b []app.Card) []app.Card {
	// 创建一个 map 来存储 b 中的元素
	bMap := make(map[app.Card]struct{})
	for _, card := range b {
		bMap[card] = struct{}{} // 用空结构体来表示集合中的元素
	}

	var difference []app.Card
	// 遍历 a 中的每个 card，检查它是否在 b 中
	for _, card := range a {
		if _, found := bMap[card]; !found {
			difference = append(difference, card) // 如果不在 b 中，就加到差集
		}
	}

	return difference
}

// 手牌分组
func (h *Hand) groupCards(cards []app.Card) {
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

		var sequence []app.Card
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

func (h *Hand) findGap1Cards() {
	h.suitCards = make(map[string][]app.Card, 4)
	h.groupCards(h.invalid)
	h.invalid = []app.Card{}

	gapCardBlackBoard := h.judgeMostScore(h.findGapFromS2L(), h.findGapFromL2S())

	fmt.Println()
	fmt.Println("黑板")
	for suit, gapC := range gapCardBlackBoard {
		fmt.Printf("花色 %s 的最佳牌组: ", suit)
		if gapC.Status == app.HT {
			// app.HT 牌型
			fmt.Println("首尾牌型")
			for _, card := range gapC.Cards {
				fmt.Printf("%d ", card.Value)
			}
			fmt.Println("分值", gapC.Score)
			fmt.Printf("癞子使用次数: %d\n", gapC.JokerUseNum)
			fmt.Println()
		}

		if gapC.Status == app.Bd {
			// app.HT 牌型
			fmt.Println("卡隆牌型")
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
func (h *Hand) findGapFromS2L() map[string]app.GapCard {
	blackBoard := map[string]app.GapCard{}

	for suit, cards := range h.suitCards {
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		gapC := app.GapCard{
			Status: app.NotStatus,
			Cards:  []app.Card{},
		}

		for i := 0; i < len(cards); i++ {
			if len(gapC.Cards) == 0 {
				gapC.Cards = []app.Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == app.NotStatus {
				// app.HT 牌型
				if cards[i].Value == cards[i-1].Value+1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.HT
					gapC.Score += cards[i].Value
				}

				// BT 牌型
				if cards[i].Value == cards[i-1].Value+2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.Bd
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == app.HT {
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
func (h *Hand) findGapFromL2S() map[string]app.GapCard {
	blackBoard := map[string]app.GapCard{}
	for suit, cards := range h.suitCards {
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value > cards[j].Value
		})

		gapC := app.GapCard{
			Status: app.NotStatus,
			Cards:  []app.Card{},
		}

		for i := 0; i < len(cards); i++ {
			if len(gapC.Cards) == 0 {
				gapC.Cards = []app.Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			if gapC.Status == app.NotStatus {
				// app.HT 牌型
				if cards[i].Value == cards[i-1].Value-1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.HT
					gapC.Score += cards[i].Value
				}

				// BT 牌型
				if cards[i].Value == cards[i-1].Value-2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.Bd
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}
				continue
			}

			if gapC.Status == app.HT {
				if gapC.Cards[len(gapC.Cards)-1].Value-2 == cards[i].Value && gapC.JokerUseNum == 0 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}

				if gapC.Cards[len(gapC.Cards)-1].Value-1 == cards[i].Value && gapC.JokerUseNum == 1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
				continue
			}

			// BT 牌型
			if gapC.Status == app.Bd {
				if gapC.Cards[len(gapC.Cards)-1].Value-1 == cards[i].Value && gapC.JokerUseNum == 1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
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

func (h *Hand) Run() {
	// 初始化手牌
	h.initHand()
	// 找癞子
	h.findJoker(0)
	// 分组
	h.groupCards(h.cards)
	// 找顺子
	h.findSequences()

	fmt.Println("未处理的牌", h.suitCards)
	fmt.Println("有效牌", h.valid)
	fmt.Println("无效牌", h.invalid)
	fmt.Println("joker", h.joker)

	// 找间隙为1的牌
	h.findGap1Cards()
}

func NewHand() *Hand {
	return &Hand{
		cards:     []app.Card{},
		joker:     []app.Card{},
		valid:     []app.Card{},
		invalid:   []app.Card{},
		gap1Cards: []app.Card{},
		suitCards: make(map[string][]app.Card, 4),
	}
}
