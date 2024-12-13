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

func main() {
	cards1 := []Card{
		{Suit: "A", Value: 1},
		{Suit: "A", Value: 2},
		{Suit: "A", Value: 3},
		{Suit: "A", Value: 12},
		{Suit: "A", Value: 12},
	}
	var valid, invalid []Card

	valid1, invalid1, score1 := findBestSequence(cards1)
	valid2, invalid2, score2 := findBestSequence2(cards1)

	//score2 = 0
	if score1 > score2 {
		valid = valid1
		invalid = invalid1
	} else {
		valid = valid2
		invalid = invalid2
	}

	fmt.Printf("Test Case 1: Valid: %v, Invalid: %v\n", valid, invalid)

}
