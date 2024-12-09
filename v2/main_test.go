package main

import (
	"fmt"
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

func judgeCardLength(t *testing.T, h *internal.Hand, valid []app.Card) []app.Card {
	rawCards := h.GetCards()
	resCards := append(valid)

	for _, p := range h.GetPure() {
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetPureWithJoker() {
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetSet() {
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetSetWithJoker() {
		resCards = append(resCards, p...)
	}

	if len(rawCards) != len(resCards) {
		res := handSliceDifference(rawCards, resCards)
		t.Fatal(fmt.Printf("结果长度不一致, 原长度%d, 返回长度 %d, \n 他们缺少: %v", len(rawCards), len(resCards), res))
	}
	return resCards
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

	_, invalid := h.RunTest(6)
	resCards := judgeCardLength(t, h, invalid)

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

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight2(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 4},
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 5},
		{Suit: app.C, Value: 13},

		{Suit: app.B, Value: 3},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 11},

		{Suit: app.A, Value: 7},
	})

	_, invalid := h.RunTest(4)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},

		{Suit: app.D, Value: 4},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 13},
	}

	wantI := []app.Card{
		{Suit: app.C, Value: 5},
		{Suit: app.C, Value: 13},

		{Suit: app.B, Value: 3},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 11},

		{Suit: app.A, Value: 7},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight3(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 4},
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},
		{Suit: app.D, Value: 10},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},

		{Suit: app.D, Value: 13},
		{Suit: app.C, Value: 13},
		{Suit: app.B, Value: 13},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 11},

		{Suit: app.B, Value: 3},
	})

	_, invalid := h.RunTest(4)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},

		{Suit: app.D, Value: 10},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},

		{Suit: app.D, Value: 13},
		{Suit: app.C, Value: 13},
		{Suit: app.B, Value: 13},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 11},
		{Suit: app.D, Value: 4},
	}

	wantI := []app.Card{
		{Suit: app.B, Value: 3},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight4(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.JokerB, Value: 14},

		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 9},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 10},
		{Suit: app.C, Value: 11},
		{Suit: app.C, Value: 1},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 2},

		{Suit: app.A, Value: 8},
	})

	_, invalid := h.RunTest(2)
	resCards := judgeCardLength(t, h, invalid)

	var want []app.Card

	wantI := []app.Card{
		{Suit: app.JokerB, Value: 14},

		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 9},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 10},
		{Suit: app.C, Value: 11},
		{Suit: app.C, Value: 1},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 2},

		{Suit: app.A, Value: 8},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight5(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 10},

		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 8},
		{Suit: app.C, Value: 9},
		{Suit: app.C, Value: 11},

		{Suit: app.B, Value: 13},
		{Suit: app.B, Value: 2},

		{Suit: app.A, Value: 2},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 1},
	})

	_, invalid := h.RunTest(9)
	resCards := judgeCardLength(t, h, invalid)

	var want []app.Card

	wantI := []app.Card{
		{Suit: app.JokerB, Value: 14},

		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 9},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 10},
		{Suit: app.C, Value: 11},
		{Suit: app.C, Value: 1},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 2},

		{Suit: app.A, Value: 8},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight6(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 1},
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 5},
		{Suit: app.D, Value: 6},

		{Suit: app.C, Value: 2},
		{Suit: app.C, Value: 7},
		{Suit: app.C, Value: 1},

		{Suit: app.B, Value: 13},
		{Suit: app.B, Value: 9},

		{Suit: app.A, Value: 2},
		{Suit: app.A, Value: 5},
	})

	_, invalid := h.RunTest(3)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 1},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 5},
		{Suit: app.D, Value: 6},

		{Suit: app.C, Value: 2},
		{Suit: app.C, Value: 7},
		{Suit: app.C, Value: 1},

		{Suit: app.B, Value: 13},
		{Suit: app.B, Value: 9},

		{Suit: app.A, Value: 2},
		{Suit: app.A, Value: 5},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight7(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 10},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},

		{Suit: app.C, Value: 2},
		{Suit: app.C, Value: 5},

		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 7},
		{Suit: app.B, Value: 8},

		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 11},
	})

	_, invalid := h.RunTest(2)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 7},
		{Suit: app.B, Value: 8},

		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 10},
		{Suit: app.C, Value: 2},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},
		{Suit: app.C, Value: 5},
		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 11},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}
func TestStraight8(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 10},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 8},

		{Suit: app.C, Value: 8},
		{Suit: app.C, Value: 2},

		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 7},
		{Suit: app.B, Value: 8},

		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 11},
	})

	_, invalid := h.RunTest(2)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.D, Value: 12},
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 10},
		{Suit: app.D, Value: 11},

		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 9},
		{Suit: app.A, Value: 11},
		{Suit: app.C, Value: 2},

		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 7},
		{Suit: app.B, Value: 8},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 8},
		{Suit: app.C, Value: 8},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}
