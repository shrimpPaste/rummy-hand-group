package internal

import (
	"math/rand"
	"rummy-logic-v3/pkg/app"
	"time"
)

// handSliceDifference 找两个数组之间的差集
func (h *Hand) handSliceDifference(a, b []app.Card) []app.Card {
	// 用 map 记录 b 中每张卡片的数量
	bCount := make(map[app.Card]int)
	for _, card := range b {
		bCount[card]++ // 记录每张卡片出现的次数
	}

	var difference []app.Card
	// 遍历 a，检查每张卡片是否在 b 中以及出现的次数
	for _, card := range a {
		if count, found := bCount[card]; found && count > 0 {
			bCount[card]-- // b 中减少一次计数
		} else {
			difference = append(difference, card) // 如果 b 中没有或计数为 0，则加入差值
		}
	}

	return difference
}

// handSliceIntersection 找两个数组的交集
func (h *Hand) handSliceIntersection(a, b []app.Card) []app.Card {
	// 创建一个 map 来存储 b 中的元素
	bMap := make(map[app.Card]struct{})
	for _, card := range b {
		bMap[card] = struct{}{} // 用空结构体来表示集合中的元素
	}

	var intersection []app.Card
	// 遍历 a 中的每个 card，检查它是否在 b 中
	for _, card := range a {
		if _, found := bMap[card]; found {
			intersection = append(intersection, card) // 如果在 b 中，就加到交集
		}
	}

	return intersection
}

// 手牌分组
func (h *Hand) groupCards(suitCards map[string][]app.Card, cards []app.Card) {
	for _, card := range cards {
		suitCards[card.Suit] = append(suitCards[card.Suit], card)
	}
}

func (h *Hand) calculateScore(cards []app.Card) int {
	score := 0
	for _, card := range cards {
		if card.Value == h.GetWildJoker().Value {
			continue
		}

		if card.Value == 1 || card.Value > 10 {
			score += 10
		} else {
			score += card.Value
		}
	}
	return score
}

// InitializeDeck 初始化牌堆 （两副牌）
func InitializeDeck() (deck []app.Card) {
	for i := 0; i < 2; i++ {
		for _, suit := range []string{app.A, app.B, app.C, app.D} {
			for value := 1; value <= 13; value++ {
				deck = append(deck, app.Card{Suit: suit, Value: value})
			}
		}

		// 添加大小王
		deck = append(deck, app.Card{Suit: app.JokerA, Value: 0})
		deck = append(deck, app.Card{Suit: app.JokerB, Value: 0})
	}

	return
}

func DealCards(deck *[]app.Card, numCards int) []app.Card {
	// numCards不能超过排堆大小
	if numCards > len(*deck) {
		panic("too many cards requested")
	}
	hand := (*deck)[:numCards]
	*deck = (*deck)[numCards:]
	return hand
}

func ShuffleDeck(deck []app.Card) []app.Card {
	rand.NewSource(time.Now().UnixNano()) // 设置随机种子
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func getCardsResult(cards []app.Card) []int {
	var myCards []int
	for _, c := range cards {
		if c.Suit == app.A {
			myCards = append(myCards, c.Value+48)
		} else if c.Suit == app.B {
			myCards = append(myCards, c.Value+32)
		} else if c.Suit == app.C {
			myCards = append(myCards, c.Value+16)
		} else if c.Suit == app.D {
			myCards = append(myCards, c.Value)
		} else if c.Suit == app.JokerA {
			myCards = append(myCards, 79)
		} else if c.Suit == app.JokerB {
			myCards = append(myCards, 78)
		}
	}

	if len(myCards) == 0 {
		return []int{0}
	}
	return myCards
}

func removeDuplicates(cards []app.Card) []app.Card {
	// 使用 map 来记录已经出现过的 Card
	seen := make(map[app.Card]bool)
	var result []app.Card

	for _, card := range cards {
		// 如果 map 中没有这个 Card，则添加到结果中，并标记为已见
		if _, ok := seen[card]; !ok {
			seen[card] = true
			result = append(result, card)
		}
	}

	return result
}
