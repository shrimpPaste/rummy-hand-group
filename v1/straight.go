package stright

import (
	"math/rand"
	"sort"
	"time"
)

// Suit 类型
type Suit string

const (
	A      Suit = "A"      // 黑桃
	B      Suit = "B"      // 红桃
	C      Suit = "C"      // 梅花
	D      Suit = "D"      // 方片
	JokerA Suit = "JokerA" // 大鬼
	JokerB Suit = "JokerB" // 小鬼
)

// Card 定义
type Card struct {
	Suit  Suit
	Value int // 1~13（Ace 到 King），Joker 的值为 0
}

// Hand 定义
type Hand []Card

// 初始化牌堆 （两副牌）
func initializeDeck() (deck Hand) {
	for i := 0; i < 2; i++ {
		for _, suit := range []Suit{A, B, C, D} {
			for value := 1; value <= 13; value++ {
				deck = append(deck, Card{Suit: suit, Value: value})
			}
		}

		// 添加大小王
		deck = append(deck, Card{Suit: JokerA, Value: 0})
		deck = append(deck, Card{Suit: JokerB, Value: 0})
	}

	return
}

func dealCards(deck *Hand, numCards int) Hand {
	// numCards不能超过排堆大小
	if numCards > len(*deck) {
		panic("too many cards requested")
	}
	hand := (*deck)[:numCards]
	*deck = (*deck)[numCards:]
	return hand
}

// 随机打乱牌
func shuffleDeck(deck Hand) Hand {
	rand.NewSource(time.Now().UnixNano()) // 设置随机种子
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

// StraightScore 顺子评分
type StraightScore struct {
	NotJoker bool // 是否包含鬼牌 true == 有鬼牌
	Score    int  // 分值
}

// 获取 Hand a 和 b 的差集 (a - b)
func handSliceDifference(a, b Hand) Hand {
	// 创建一个 map 来存储 b 中的元素
	bMap := make(map[Card]struct{})
	for _, card := range b {
		bMap[card] = struct{}{} // 用空结构体来表示集合中的元素
	}

	var difference Hand
	// 遍历 a 中的每个 card，检查它是否在 b 中
	for _, card := range a {
		if _, found := bMap[card]; !found {
			difference = append(difference, card) // 如果不在 b 中，就加到差集
		}
	}

	return difference
}

func groupCards(hand Hand, wildcardValue int) (valid map[Suit]Hand, invalid map[Suit]Hand) {
	valid = make(map[Suit]Hand)
	invalid = make(map[Suit]Hand)

	winnerStraightRule := make(map[Suit]*StraightScore)

	// 对花色进行分类
	cardMap := make(map[Suit]Hand)
	var wildcards Hand

	for _, card := range hand {
		if card.Value == wildcardValue || card.Suit == JokerA || card.Suit == JokerB {
			// 鬼牌癞子单独存储
			wildcards = append(wildcards, card)
		} else {
			cardMap[card.Suit] = append(cardMap[card.Suit], card)
		}
	}

	for suit, cards := range cardMap {
		// 如果cards长度小于3 全部都是废牌
		if len(cards) < 3 {
			invalid[suit] = cards
			continue
		}

		// 按点数排序
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		// 临时存放一个顺子
		var currentStraight Hand

		// 遍历手牌，找到所有合法顺子
		for _, card := range cards {
			// 如果是顺子的开始
			if len(currentStraight) == 0 {
				currentStraight = append(currentStraight, card)
			} else if card.Value == currentStraight[len(currentStraight)-1].Value+1 {
				// 如果当前的值刚好是历史的大1，那就说明是顺子
				currentStraight = append(currentStraight, card)
			} else {
				// 问：为什么这里要有两次和下面一样对 valid和invalid的赋值操作？
				// 答：如果card刚好是数组最后一个，在上面的判断已经循环结束，就不能走到这里
				if len(currentStraight) >= 3 {
					valid[suit] = append(valid[suit], currentStraight...)
				} else {
					invalid[suit] = append(invalid[suit], currentStraight...)
				}
				// 重置顺子
				currentStraight = []Card{card}
			}
		}

		// 所以这里还需要一次赋值操作，这是对数组最后一个也是连续顺子的赋值
		if len(currentStraight) >= 3 {
			valid[suit] = append(valid[suit], currentStraight...)
		} else {
			invalid[suit] = append(invalid[suit], currentStraight...)
		}

		// 特殊处理
		// 12 13 1 的分值比 1 2 3的价值更大
		for i := 0; i < len(invalid[suit])-1; i++ {
			// 目标：如果 第一轮无效牌是 12 13 就需要去找 正确牌 顺子 1 2 3 ... 取出1 还能把 2 3 4 .. 组合成正常牌
			// 如果第一个值为1

			if invalid[suit][i].Value == 12 && invalid[suit][i+1].Value == 13 {
				if valid[suit][0].Value == 1 {
					var start1Straight Hand // 1 开头的顺子

					for _, v := range valid[suit] {
						if v.Value == 1 {
							continue
						}
						if len(start1Straight) == 0 {
							start1Straight = append(start1Straight, v)
						} else if v.Value == start1Straight[len(start1Straight)-1].Value+1 {
							start1Straight = append(start1Straight, v)
						}
					}

					//temp := make(Hand, len(invalid))
					//copy(temp, invalid) // 将 valid 的所有元素复制到 temp 中
					//invalid = invalid[:i]
					//valid = append(valid, temp[i:]...)
					valid[suit] = append(valid[suit], invalid[suit][i:]...)
					invalid[suit] = append(invalid[suit][:i], invalid[suit][i+2:]...)

					if len(start1Straight) < 3 {
						valid[suit] = handSliceDifference(valid[suit], start1Straight)
						invalid[suit] = append(invalid[suit], start1Straight...) // 把取出 1后不符合 顺子规则的数组添加到无效牌当
					}

					sort.Slice(valid[suit], func(i, j int) bool {
						// 如果 i 是 12 或 13，且 j 不是，i 排在前面
						if valid[suit][i].Value == 12 || valid[suit][i].Value == 13 {
							return true
						}
						// 如果 j 是 12 或 13，且 i 不是，j 排在前面
						if valid[suit][j].Value == 12 || valid[suit][j].Value == 13 {
							return false
						}
						// 对其他的情况，按照 Value 排序
						return valid[suit][i].Value < valid[suit][j].Value
					})
				}
			}
		}
		sort.Slice(invalid[suit], func(i, j int) bool {
			return invalid[suit][i].Value < invalid[suit][j].Value
		})
	}

	for s, v := range valid {
		tScore := &StraightScore{
			NotJoker: false,
			Score:    0,
		}
		for _, c := range v {
			if c.Value == 1 && c.Value > 10 {
				tScore.Score += 10
			} else {
				tScore.Score += c.Value
			}
			if c.Value == wildcardValue {
				// 包含癞子
				tScore.NotJoker = true
				tScore.Score -= c.Value
			}
		}
		winnerStraightRule[s] = tScore
	}

	// 一次只消费一张癞子牌

	// 处理间隙为 1 的
	for s, v := range valid {
		for _, iv := range invalid {
			for _, c := range iv {
				if v[len(v)-1].Value+2 == c.Value || v[0].Value-2 == c.Value {
					valid[s] = append(valid[s], c)
				}
			}
		}
	}

	// 添加癞子牌
	//for suit, cards := range invalid {
	//	for _, card := range cards {
	//		if card.Value == wildcardValue {
	//			wildcards = append(wildcards, card)
	//			invalid[suit] = handSliceDifference(invalid[suit], wildcards)
	//		}
	//	}
	//}

	// 找只差一张牌的
	//var maybeValid []Hand
	//for _, i := range invalid {
	//	if len(i) < 2 {
	//		continue
	//	}
	//	first := i[0]
	//
	//	if first.Value == 1 {
	//		for j := 1; j < len(i); j++ {
	//			if i[j].Value-first.Value == 13 || i[j].Value-first.Value == 12 {
	//				maybeValid = append(maybeValid, Hand{
	//					i[j], first,
	//				})
	//			}
	//		}
	//		continue
	//	}
	//	if first.Value == 12 {
	//		for j := 1; j < len(i); j++ {
	//			if i[j].Value-first.Value == 12 || i[j].Value-first.Value == 1 {
	//				maybeValid = append(maybeValid, Hand{
	//					i[j], first,
	//				})
	//			}
	//		}
	//		continue
	//	}
	//
	//	for j := 1; j < len(i); j++ {
	//		if i[j].Value-first.Value >= -2 && i[j].Value-first.Value <= 2 {
	//			maybeValid = append(maybeValid, Hand{
	//				i[j], first,
	//			})
	//		}
	//	}
	//}
	//
	//// 循环癞子牌
	//for _, r := range maybeValid {
	//	if len(wildcards) >= 1 {
	//		invalid[r[0].Suit] = handSliceDifference(invalid[r[0].Suit], r)
	//		r = append(r, wildcards[0])
	//		wildcards = wildcards[1:]
	//		valid[r[0].Suit] = append(valid[r[0].Suit], r...)
	//	}
	//}

	return
}

// find111Card 找刻子
func find111Card(valid, invalid map[Suit]Hand) (map[Suit]Hand, map[Suit]Hand) {
	diff111Suit := findSameValueDifferentSuit(invalid)

	// 对每个找出的刻子（hand）
	for _, hand := range diff111Suit {
		// 对于 invalid 中的每个花色，更新 invalid[s]
		for s, h := range invalid {
			invalid[s] = handSliceDifference(h, hand) // 移除手牌中的刻子
		}

		// 将刻子添加到 valid 中
		valid[A] = append(valid[A], hand...)
	}

	return valid, invalid
}

func findSameValueDifferentSuit(cards map[Suit]Hand) map[int]Hand {
	valueMap := make(map[int]Hand)
	for _, hand := range cards {
		for _, card := range hand {
			valueMap[card.Value] = append(valueMap[card.Value], card)
		}
	}

	result := make(map[int]Hand)
	for value, hand := range valueMap {
		if len(hand) < 3 {
			continue
		}

		suitCount := make(map[Suit]bool) // 用于记录已添加的花色
		tempHand := Hand{}

		for _, card := range hand {
			if !suitCount[card.Suit] { // 只添加未添加过的花色
				suitCount[card.Suit] = true
				tempHand = append(tempHand, card)
			}
		}

		if len(tempHand) >= 3 {
			result[value] = tempHand
		}
	}

	return result
}
