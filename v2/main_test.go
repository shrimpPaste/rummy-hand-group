package main

import (
	"rummy-group-v2/internal"
	"rummy-group-v2/pkg/app"
	"testing"
)

func handSliceDifference(a, b []app.Card) []app.Card {
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

func judgeCardLength(t *testing.T, rawCards, resCards []app.Card) {
	if len(rawCards) != len(resCards) {
		res := handSliceDifference(rawCards, resCards)
		t.Errorf("结果长度不一致, 原长度%d, 返回长度 %d, \n 他们缺少: %v", len(rawCards), len(resCards), res)
	}
}

func TestStraight1(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 4},
		{Suit: app.D, Value: 5},
		{Suit: app.D, Value: 6},

		{Value: 2, Suit: app.C},
		{Value: 4, Suit: app.C},
		{Value: 5, Suit: app.C},

		{Value: 5, Suit: app.B},
		{Value: 6, Suit: app.B},
		{Value: 3, Suit: app.B},
		{Value: 6, Suit: app.B},

		{Value: 2, Suit: app.A},
		{Value: 3, Suit: app.A},
	})

	valid, invalid := h.RunTest(6)
	judgeCardLength(t, h.GetCards(), append(valid, invalid...))

	want := []app.Card{
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 4},
		{Suit: app.D, Value: 5},

		{Suit: app.D, Value: 6},
		{Value: 2, Suit: app.C},
		{Value: 4, Suit: app.C},
		{Value: 5, Suit: app.C},

		{Value: 5, Suit: app.B},
		{Value: 6, Suit: app.B},
		{Value: 3, Suit: app.B},

		{Value: 6, Suit: app.B},
		{Value: 2, Suit: app.A},
		{Value: 3, Suit: app.A},
	}

	res := handSliceDifference(want, valid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(valid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, valid, res)
		t.Errorf("invalid %v", invalid)
	}
}
