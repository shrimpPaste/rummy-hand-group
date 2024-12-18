package internal

import (
	"rummy-logic-v3/pkg/app"
)

// 鉴定牌型是否有一个及以上的顺子
func (h *Hand) judgeIsHave1Seq(cards []app.Card) bool {
	// 该函数调用应该在第一轮找顺子的时候判断
	if len(cards) >= 3 {
		return true
	}
	return false
}

func (h *Hand) judgeIsHave2Seq(pureCards, pureWithCards []app.Card) bool {
	if len(pureCards) >= 6 || (len(pureCards) >= 3 && len(pureWithCards) >= 3) {
		return true
	}
	return false
}
