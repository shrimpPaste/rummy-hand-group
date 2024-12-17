package main

import (
	"fmt"
	"sort"
)

type Card struct {
	Suit  string
	Value int
}

func findBestSequence(cards []Card) (valid, invalid []Card, score int) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	calculateScore := func(seq []Card) int {
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

	seq := []Card{cards[0]}

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
			valid2, invalid2, score2 := findBestSequence(invalid)
			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []Card
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
			valid2, invalid2, score2 := findBestSequence(invalid)
			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []Card
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

func findBestSequence2(cards []Card) (valid, invalid []Card, score int) {
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
	seq := []Card{cards[0]}
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
			valid2, invalid2, score2 := findBestSequence2(invalid)

			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []Card
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
			valid2, invalid2, score2 := findBestSequence2(invalid)

			if len(valid2) >= 3 {
				valid = append(valid, valid2...)
				score += score2

				valid2Map := make(map[Card]bool)
				for _, card := range valid2 {
					valid2Map[card] = true
				}

				var remainingInvalid []Card
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

func handSliceDifference(a, b []Card) []Card {
	// 用 map 记录 b 中每张卡片的数量
	bCount := make(map[Card]int)
	for _, card := range b {
		bCount[card]++ // 记录每张卡片出现的次数
	}

	var difference []Card
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

func findGap(cards []Card, jokers []Card) ([]Card, []Card, []Card) {
	var result, overCards []Card

	singleCards := removeDuplicates(cards)
	overCards = handSliceDifference(cards, singleCards)
	isUsed := false

	sort.Slice(singleCards, func(i, j int) bool {
		return singleCards[i].Value > singleCards[j].Value
	})

	var tempResult []Card

	for i := 0; i < len(singleCards)-1; i++ {
		for j := i + 1; j < len(singleCards); j++ {
			gap := singleCards[j].Value - singleCards[i].Value
			if len(tempResult) > 0 {
				gap = singleCards[j].Value - tempResult[len(tempResult)-1].Value
			}
			if gap == -1 {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], singleCards[j])
				} else {
					tempResult = append(tempResult, singleCards[j])
				}
				break
			} else if gap == -2 && !isUsed {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], jokers[0], singleCards[j])
					i++
				} else {
					tempResult = append(tempResult, jokers[0], singleCards[j])
					i++
				}
				jokers = jokers[1:]
				isUsed = true
			} else {
				if len(tempResult) != 0 {
					overCards = append(overCards, singleCards[j])
				} else {
					overCards = append(overCards, singleCards[i])
				}
				break
			}
		}

	}

	if len(tempResult) >= 3 {
		result = append(result, tempResult...)
	}

	if len(tempResult) == 2 && len(jokers) > 0 {
		result = append(result, tempResult...)
		result = append(result, jokers[0])
		jokers = jokers[1:]
	}

	if len(tempResult) == 0 && len(singleCards) > 0 && len(jokers) > 2 {
		result = append(result, singleCards[0], jokers[0], jokers[1])
		overCards = singleCards[1:]
		jokers = jokers[2:]
	}

	return overCards, result, jokers
}

//func findGapTo14(cards []Card, jokers []Card) ([]Card, []Card, []Card) {
//	var result, overCards []Card
//
//	singleCards := removeDuplicates(cards)
//	overCards = handSliceDifference(cards, singleCards)
//	isUsed := false
//
//	for i := range singleCards {
//		if singleCards[i].Value == 1 {
//			singleCards[i].Value = 14
//		}
//	}
//
//	sort.Slice(singleCards, func(i, j int) bool {
//		return singleCards[i].Value > singleCards[j].Value
//	})
//
//	var tempResult []Card
//
//	for i := 0; i < len(singleCards)-1; i++ {
//		for j := i + 1; j < len(singleCards); j++ {
//			gap := singleCards[j].Value - singleCards[i].Value
//			if len(tempResult) > 0 {
//				gap = singleCards[j].Value - tempResult[len(tempResult)-1].Value
//			}
//			if gap == -1 {
//				if len(tempResult) == 0 {
//					tempResult = append(tempResult, singleCards[i], singleCards[j])
//				} else {
//					tempResult = append(tempResult, singleCards[j])
//				}
//				break
//			} else if gap == -2 && !isUsed {
//				if len(tempResult) == 0 {
//					tempResult = append(tempResult, singleCards[i], jokers[0], singleCards[j])
//					i++
//				} else {
//					tempResult = append(tempResult, jokers[0], singleCards[j])
//					i++
//				}
//				jokers = jokers[1:]
//				isUsed = true
//			} else {
//				if len(tempResult) != 0 {
//					overCards = append(overCards, singleCards[j])
//				} else {
//					overCards = append(overCards, singleCards[i])
//				}
//				break
//			}
//		}
//
//	}
//
//	if len(tempResult) >= 3 {
//		result = append(result, tempResult...)
//	}
//
//	if len(tempResult) == 2 && len(jokers) > 0 {
//		result = append(result, tempResult...)
//		result = append(result, jokers[0])
//		jokers = jokers[1:]
//	}
//
//	for i := range result {
//		if result[i].Value == 14 {
//			result[i].Value = 1
//		}
//	}
//
//	for i := range overCards {
//		if overCards[i].Value == 14 {
//			overCards[i].Value = 1
//		}
//	}
//
//	return overCards, result, jokers
//}

func removeDuplicates(cards []Card) []Card {
	// 使用 map 来记录已经出现过的 Card
	seen := make(map[Card]bool)
	var result []Card

	for _, card := range cards {
		// 如果 map 中没有这个 Card，则添加到结果中，并标记为已见
		if _, ok := seen[card]; !ok {
			seen[card] = true
			result = append(result, card)
		}
	}

	return result
}

func findGap3(cards []Card, jokers []Card) ([]Card, []Card, []Card) {
	var result, overCards []Card

	singleCards := removeDuplicates(cards)
	overCards = handSliceDifference(cards, singleCards)

	sort.Slice(singleCards, func(i, j int) bool {
		return singleCards[i].Value > singleCards[j].Value
	})
	isUsed := false

	var tempResult []Card

	for i := 0; i <= len(singleCards)-1; i++ {
		if i == len(singleCards)-1 && len(tempResult) > 0 && singleCards[i].Value-tempResult[len(tempResult)-1].Value == 1 {
			tempResult = append(tempResult, singleCards[i])
		}
		if i == len(singleCards)-1 && len(tempResult) == 0 {
			overCards = append(overCards, singleCards[i])
		}
		for j := i + 1; j < len(singleCards); j++ {
			gap := singleCards[j].Value - singleCards[i].Value
			if len(tempResult) > 0 {
				gap = singleCards[j].Value - tempResult[len(tempResult)-1].Value
			}

			if gap == -1 {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], singleCards[j])
				} else {
					tempResult = append(tempResult, singleCards[j])
				}
				break
			} else if gap == -3 && len(jokers) > 1 && !isUsed {
				if len(tempResult) == 0 {
					tempResult = append(tempResult, singleCards[i], jokers[0], jokers[1], singleCards[j])
					i++ // 跳过当前处理的卡
				} else {
					tempResult = append(tempResult, jokers[0], jokers[1], singleCards[j])
					i++
				}
				jokers = jokers[2:]
				isUsed = true
			} else {
				if len(tempResult) != 0 {
					overCards = append(overCards, singleCards[j])
				} else {
					overCards = append(overCards, singleCards[i])
				}
				break
			}
		}
	}

	if len(tempResult) >= 3 {
		result = append(result, tempResult...)
	}

	if len(tempResult) == 2 && len(jokers) > 0 {
		result = append(result, tempResult...)
		result = append(result, jokers[0])
		jokers = jokers[1:]
	}

	return overCards, result, jokers
}

func calculateScore(cards []Card) int {
	score := 0
	for _, card := range cards {
		if card.Value == 2 {
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

func main() {
	//cards1 := []Card{
	//	{Suit: "A", Value: 2},
	//	{Suit: "A", Value: 4},
	//	{Suit: "A", Value: 6},
	//	{Suit: "A", Value: 12},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 1},
	//	{Suit: "A", Value: 3},
	//	{Suit: "A", Value: 3},
	//	{Suit: "A", Value: 4},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 1},
	//	{Suit: "A", Value: 1},
	//	{Suit: "A", Value: 7},
	//	{Suit: "A", Value: 8},
	//	{Suit: "A", Value: 10},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 4},
	//	{Suit: "A", Value: 4},
	//	{Suit: "A", Value: 8},
	//	{Suit: "A", Value: 8},
	//	{Suit: "A", Value: 9},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 13},
	//	{Suit: "A", Value: 3},
	//	{Suit: "A", Value: 7},
	//	//{Suit: "A", Value: 10},
	//	//{Suit: "A", Value: 12},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 7},
	//	{Suit: "A", Value: 8},
	//	{Suit: "A", Value: 10},
	//	{Suit: "A", Value: 13},
	//}

	//cards1 := []Card{
	//	{Suit: "A", Value: 1},
	//	{Suit: "A", Value: 2},
	//}
	//
	//jokers := []Card{
	//	{Suit: "D", Value: 5},
	//}
	//
	//var result, overCards, overJokers []Card
	//
	//overCards1, result1, overJokers1 := findGap(cards1, jokers)
	//
	//resultScore := calculateScore(result1)
	//
	//// 处理A == 14的情况
	//cards2 := cards1
	//for i := range cards2 {
	//	if cards2[i].Value == 1 {
	//		cards2[i].Value = 14
	//	}
	//}
	//overCards2, result2, overJokers2 := findGap(cards2, jokers)
	//resultScore2 := calculateScore(result2)
	//
	//if resultScore2 > resultScore {
	//	for i := range result2 {
	//		if result2[i].Value == 14 {
	//			result2[i].Value = 1
	//		}
	//	}
	//
	//	for i := range overCards2 {
	//		if overCards2[i].Value == 14 {
	//			overCards2[i].Value = 1
	//		}
	//	}
	//
	//	result = result2
	//	overCards = overCards2
	//	overJokers = overJokers2
	//
	//} else {
	//	result = result1
	//	overCards = overCards1
	//	overJokers = overJokers1
	//}
	//fmt.Println("overCards", overCards, "result", result, "jokers", overJokers)

	cards3 := []Card{
		{Suit: "A", Value: 7},
		{Suit: "A", Value: 10},
	}
	jokers2 := []Card{
		{Suit: "D", Value: 5},
		{Suit: "D", Value: 5},
	}
	overCards3, result3, overJokers3 := findGap3(cards3, jokers2)
	fmt.Println("overCards3", overCards3, "result3", result3, "jokers3", overJokers3)
}
