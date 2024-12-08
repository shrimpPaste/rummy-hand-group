package internal

import "rummy-group-v2/pkg/app"

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
