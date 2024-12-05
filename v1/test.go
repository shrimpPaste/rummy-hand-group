package stright

import "sort"

// groupCards 组牌，返回有效组牌和无效牌
func groupCards2(hand Hand, wildcardValue int) (valid Hand, invalid Hand) {
	valid = make(Hand, 0, len(hand)) // 预分配空间，避免多次扩容
	invalid = make(Hand, 0, len(hand))

	// 对花色进行分类
	cardMap := make(map[Suit]Hand)
	var wildcards Hand

	// 分类牌，癞子牌单独存储
	for _, card := range hand {
		if card.Value == wildcardValue || card.Suit == JokerA || card.Suit == JokerB {
			// 癞子单独存储
			wildcards = append(wildcards, card)
		} else {
			cardMap[card.Suit] = append(cardMap[card.Suit], card)
		}
	}

	// 遍历每种花色的手牌
	for _, cards := range cardMap {
		// 对当前花色的手牌按点数排序
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})

		// 临时存放顺子
		var currentStraight []Card

		// 遍历当前花色的牌，找到所有合法的顺子
		for _, card := range cards {
			// 如果是顺子的开始
			if len(currentStraight) == 0 {
				currentStraight = append(currentStraight, card)
			} else if card.Value == currentStraight[len(currentStraight)-1].Value+1 {
				// 如果当前的牌是顺子的下一个
				currentStraight = append(currentStraight, card)
			} else {
				// 当前顺子结束，处理并清空临时顺子
				if len(currentStraight) >= 3 {
					valid = append(valid, currentStraight...)
				} else {
					invalid = append(invalid, currentStraight...)
				}
				currentStraight = []Card{card} // 重置顺子
			}
		}

		// 处理最后一个顺子
		if len(currentStraight) >= 3 {
			valid = append(valid, currentStraight...)
		} else {
			invalid = append(invalid, currentStraight...)
		}
	}

	// 特殊处理：处理 `12 13 1` 顺子的优先级问题
	// 先查找 invalid 中是否有 12 和 13，如果存在且 valid 中有 1，则尝试修正顺子
	for i := 0; i < len(invalid)-1; i++ {
		// 查找 12 13
		if invalid[i].Value == 12 && invalid[i+1].Value == 13 {
			// 检查 valid 中是否有 1
			if containsCard(valid, 1) {
				// 找到 valid 中的 1 开头的顺子
				var start1Straight Hand
				for _, v := range valid {
					if v.Value == 1 {
						continue
					}
					if len(start1Straight) == 0 {
						start1Straight = append(start1Straight, v)
					} else if v.Value == start1Straight[len(start1Straight)-1].Value+1 {
						start1Straight = append(start1Straight, v)
					}
				}

				// 合并 valid 和 invalid
				valid = append(valid, invalid[i:]...)
				invalid = append(invalid[:i], invalid[i+2:]...)

				// 如果 1 开头的顺子长度小于 3，则不符合顺子规则，加入 invalid
				if len(start1Straight) < 3 {
					invalid = append(invalid, start1Straight...)
				}
			}
		}
	}

	// 排序 valid 和 invalid
	sort.Slice(valid, func(i, j int) bool {
		// 优先将 12 和 13 排在前面
		if valid[i].Value == 12 || valid[i].Value == 13 {
			return true
		}
		if valid[j].Value == 12 || valid[j].Value == 13 {
			return false
		}
		// 对其他牌按 Value 排序
		return valid[i].Value < valid[j].Value
	})

	sort.Slice(invalid, func(i, j int) bool {
		return invalid[i].Value < invalid[j].Value
	})

	return valid, invalid
}

// containsCard 判断手牌中是否包含指定的牌值
func containsCard(hand Hand, value int) bool {
	for _, card := range hand {
		if card.Value == value {
			return true
		}
	}
	return false
}
