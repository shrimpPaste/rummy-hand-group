package internal

import "rummy-group-v2/pkg/app"

// 鉴定牌型是否有两个及以上的顺子
func (h *Hand) judgeIsHave2Seq() bool {
	// 该函数调用应该在第一轮找顺子的时候判断
	if len(h.valid) > 6 {
		return true
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
func (h *Hand) judgeIsSeq() bool {
	return false
}
