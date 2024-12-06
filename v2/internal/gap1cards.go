package internal

import (
	"fmt"
	"rummy-group-v2/pkg/app"
	"sort"
)

// 处理卡隆牌
func (h *Hand) findGap1Cards() {
	// 检测是否已经拥有两个及以上的顺子，并且没有使用过小丑牌。
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

// findGapFromS2L 用于寻找手中牌中每个花色的间隙牌。（从小到大的排序）
func (h *Hand) findGapFromS2L() map[string]app.GapCard {
	// 初始化一个空的映射来存储每个花色的间隙牌信息。
	blackBoard := map[string]app.GapCard{}

	// 遍历每个花色及其对应的牌。
	for suit, cards := range h.suitCards {
		// 如果该花色的牌少于2张，则认为无法形成有效的组合，将其标记为无效牌并跳过。
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}

		// 对牌进行排序，以便后续处理。
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		// 初始化一个间隙牌结构体。
		gapC := app.GapCard{
			Status: app.NotStatus,
			Cards:  []app.Card{},
		}

		// 遍历每张牌，寻找可能的间隙牌组合。
		for i := 0; i < len(cards); i++ {
			// 如果当前间隙牌组合为空，则将当前牌加入组合中，并更新分数。
			if len(gapC.Cards) == 0 {
				gapC.Cards = []app.Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			// 如果当前间隙牌组合的状态未确定，则尝试确定其状态。
			if gapC.Status == app.NotStatus {
				// 检查是否为顺子牌型。
				if cards[i].Value == cards[i-1].Value+1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.HT
					gapC.Score += cards[i].Value
				}

				// 检查是否为间隔一牌的牌型。
				if cards[i].Value == cards[i-1].Value+2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.Bd
					gapC.Score += cards[i].Value
				}
				continue
			}

			// 如果当前间隙牌组合的状态为顺子，则根据特定规则继续添加牌。
			if gapC.Status == app.HT {
				// 因为牌已经按从小到大的顺序排列，所以如果当前牌与组合中的最后一张牌值相差2，且未使用过小丑牌，则认为是合理的间隙牌。
				if gapC.Cards[len(gapC.Cards)-1].Value+2 == cards[i].Value && gapC.JokerUseNum == 0 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}

				// 如果已经使用过小丑牌，且当前牌与组合中的最后一张牌值相差1，则认为是合理的间隙牌。
				if gapC.Cards[len(gapC.Cards)-1].Value+1 == cards[i].Value && gapC.JokerUseNum == 1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
				continue
			}
		}

		// 如果最终的间隙牌组合中至少有2张牌，则将其添加到结果映射中。
		if len(gapC.Cards) >= 2 {
			blackBoard[suit] = gapC
		}
	}

	// 返回结果映射。
	return blackBoard
}

// findGapFromL2S 分析手牌中每个花色的卡片，找出可能的断牌组合。(从大到小的排序）
func (h *Hand) findGapFromL2S() map[string]app.GapCard {
	// 初始化一个映射来存储每种花色的断牌信息。
	blackBoard := map[string]app.GapCard{}

	// 遍历每种花色的卡片。
	for suit, cards := range h.suitCards {
		// 如果该花色的卡片少于2张，则认为无法形成有效的断牌组合，将其标记为无效。
		if len(cards) < 2 {
			h.invalid = append(h.invalid, cards...)
			continue
		}

		// 对卡片进行降序排列，以便后续处理。
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value > cards[j].Value
		})

		// 初始化一个断牌组合信息结构体。
		gapC := app.GapCard{
			Status: app.NotStatus,
			Cards:  []app.Card{},
		}

		// 遍历每张卡片，寻找断牌组合。
		for i := 0; i < len(cards); i++ {
			// 如果当前断牌组合中没有卡片，则将当前卡片加入组合中，并计算其分数。
			if len(gapC.Cards) == 0 {
				gapC.Cards = []app.Card{cards[i]}
				if cards[i].Value == 1 || cards[i].Value > 10 {
					gapC.Score += 10
				} else {
					gapC.Score += cards[i].Value
				}
				continue
			}

			// 如果当前断牌组合的状态未确定，则检查是否可以形成HT或BT牌型。
			if gapC.Status == app.NotStatus {
				// 检查是否可以形成HT牌型。
				if cards[i].Value == cards[i-1].Value-1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.HT
					gapC.Score += cards[i].Value
				}

				// 检查是否可以形成BT牌型。
				if cards[i].Value == cards[i-1].Value-2 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Status = app.Bd
					gapC.Score += cards[i].Value
					gapC.JokerUseNum++
				}
				continue
			}

			// 如果当前断牌组合的状态是HT，则检查是否可以继续添加卡片到组合中。
			if gapC.Status == app.HT {
				// 检查是否可以将当前卡片添加到HT组合中，根据是否已经使用过小丑来决定。
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

			// 如果当前断牌组合的状态是BT，则检查是否可以继续添加卡片到组合中。
			if gapC.Status == app.Bd {
				// 检查是否可以将当前卡片添加到BT组合中，根据是否已经使用过小丑来决定。
				if gapC.Cards[len(gapC.Cards)-1].Value-1 == cards[i].Value && gapC.JokerUseNum == 1 {
					gapC.Cards = append(gapC.Cards, cards[i])
					gapC.Score += cards[i].Value
				}
			}
		}

		// 将最后一个断牌组合信息添加到结果映射中。
		if len(gapC.Cards) >= 2 {
			blackBoard[suit] = gapC
		}
	}

	// 返回结果映射。
	return blackBoard
}
