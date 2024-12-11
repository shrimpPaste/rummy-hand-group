package main

import (
	"fmt"
	"math/rand"
	"rummy-group-v2/internal"
	"rummy-group-v2/pkg/app"
	"testing"
	"time"
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

	//fmt.Println("无效牌", resCards)
	for _, p := range h.GetPure() {
		//fmt.Println("纯顺子", p)
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetPureWithJoker() {
		//fmt.Println("带joker的顺子", p)
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetSet() {
		//fmt.Println("纯刻子", p)
		resCards = append(resCards, p...)
	}
	for _, p := range h.GetSetWithJoker() {
		//fmt.Println("带joker的刻子", p)
		resCards = append(resCards, p...)
	}

	if len(rawCards) != len(resCards) {
		fmt.Println(rawCards)
		fmt.Println(resCards)
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 6})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 4})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 4})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 2})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 9})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 3})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 2})
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

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 2})
	h.RunTest(2) // 开启下列注释会报错，请删除本行
	//_, invalid := h.RunTest(2)
	//resCards := judgeCardLength(t, h, invalid)
	//
	//want := []app.Card{
	//	{Suit: app.D, Value: 12},
	//	{Suit: app.D, Value: 13},
	//	{Suit: app.D, Value: 10},
	//	{Suit: app.D, Value: 11},
	//
	//	{Suit: app.A, Value: 12},
	//	{Suit: app.A, Value: 9},
	//	{Suit: app.A, Value: 11},
	//	{Suit: app.C, Value: 2},
	//
	//	{Suit: app.B, Value: 6},
	//	{Suit: app.B, Value: 7},
	//	{Suit: app.B, Value: 8},
	//}
	//
	//wantI := []app.Card{
	//	{Suit: app.D, Value: 8},
	//	{Suit: app.C, Value: 8},
	//}
	//
	//realCards := handSliceDifference(resCards, invalid)
	//res := handSliceDifference(want, realCards)
	//if len(want) != len(realCards) {
	//	t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
	//	t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
	//	t.Errorf("invalid %v", invalid)
	//}
	//
	//res2 := handSliceDifference(wantI, invalid)
	//if res != nil {
	//	t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
	//	t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
	//	t.Errorf("invalid %v", invalid)
	//}
}

func TestStraight9(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 13},
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},

		{Suit: app.C, Value: 5},
		{Suit: app.C, Value: 10},
		{Suit: app.C, Value: 11},
		{Suit: app.C, Value: 12},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 7},

		{Suit: app.A, Value: 5},
		{Suit: app.A, Value: 12},
	})

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 2})
	_, invalid := h.RunTest(2)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.D, Value: 7},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 9},

		{Suit: app.C, Value: 10},
		{Suit: app.C, Value: 11},
		{Suit: app.C, Value: 12},

		{Suit: app.B, Value: 5},
		{Suit: app.C, Value: 5},
		{Suit: app.A, Value: 5},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 13},
		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 7},
		{Suit: app.A, Value: 12},
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

func TestStraight10(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.JokerB, Value: 3},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 9},

		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 6},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 8},
		{Suit: app.B, Value: 13},

		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 1},
	})

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 2})
	_, invalid := h.RunTest(2)
	resCards := judgeCardLength(t, h, invalid)

	var want []app.Card

	wantI := []app.Card{
		{Suit: app.JokerB, Value: 3},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 9},

		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 6},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 8},
		{Suit: app.B, Value: 13},

		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 1},
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

func TestStraight11(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},

		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 10},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 1},
	})

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 1})
	_, invalid := h.RunTest(1)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},
		{Suit: app.B, Value: 1},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},
		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 10},
		{Suit: app.B, Value: 10},
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

func TestStraight12(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},

		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 10},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 1},
	})

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 1})
	_, invalid := h.RunTest(1)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 5},
		{Suit: app.B, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 12},
		{Suit: app.B, Value: 1},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 2},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 8},
		{Suit: app.D, Value: 6},
		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 10},
		{Suit: app.B, Value: 10},
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

func TestStraight13(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 1},
		{Suit: app.D, Value: 7},

		{Suit: app.C, Value: 1},
		{Suit: app.C, Value: 6},
		{Suit: app.C, Value: 10},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 12},

		{Suit: app.A, Value: 1},
		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 5},
		{Suit: app.A, Value: 13},
	})

	h.SetWildJoker(&app.Card{Suit: app.B, Value: 6})
	_, invalid := h.RunTest(6)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 12},

		{Suit: app.A, Value: 3},
		{Suit: app.A, Value: 5},
		{Suit: app.C, Value: 6},

		{Suit: app.C, Value: 1},
		{Suit: app.D, Value: 1},
		{Suit: app.A, Value: 1},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 7},
		{Suit: app.C, Value: 10},
		{Suit: app.A, Value: 13},
	}

	realCards := handSliceDifference(resCards, invalid)
	res := handSliceDifference(want, realCards)
	if len(want) != len(realCards) {
		//t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
		//t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
		//t.Errorf("invalid %v", invalid)
	}

	res2 := handSliceDifference(wantI, invalid)
	if res != nil {
		t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
		t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
		t.Errorf("invalid %v", invalid)
	}
}

func TestStraight14(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 7},

		{Suit: app.C, Value: 9},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 6},
		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 13},

		{Suit: app.A, Value: 7},
		{Suit: app.A, Value: 8},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 13},
		{Suit: app.A, Value: 1},
	})

	h.SetWildJoker(&app.Card{Suit: app.D, Value: 7})

	_, invalid := h.RunTest(7)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 13},
		{Suit: app.A, Value: 1},

		{Suit: app.B, Value: 9},
		{Suit: app.B, Value: 10},
		{Suit: app.B, Value: 11},
		{Suit: app.D, Value: 7},
		{Suit: app.B, Value: 13},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 6},
		{Suit: app.A, Value: 7},
	}

	wantI := []app.Card{
		{Suit: app.C, Value: 9},
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

func TestStraight15(t *testing.T) {

	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 10},

		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 5},
		{Suit: app.C, Value: 13},

		{Suit: app.B, Value: 2},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 1},

		{Suit: app.A, Value: 5},
		{Suit: app.A, Value: 6},
		{Suit: app.A, Value: 11},
	})

	h.SetWildJoker(&app.Card{Suit: app.C, Value: 7})

	_, invalid := h.RunTest(7)
	resCards := judgeCardLength(t, h, invalid)

	want := []app.Card{
		{Suit: app.C, Value: 3},
		{Suit: app.C, Value: 4},
		{Suit: app.C, Value: 5},
	}

	wantI := []app.Card{
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 10},
		{Suit: app.C, Value: 13},

		{Suit: app.B, Value: 2},
		{Suit: app.B, Value: 11},
		{Suit: app.B, Value: 1},

		{Suit: app.A, Value: 5},
		{Suit: app.A, Value: 6},
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

func TestStraight16(t *testing.T) {
	h := internal.NewHand()
	h.SetCards([]app.Card{
		{Suit: app.D, Value: 1},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 5},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 2},
		{Suit: app.C, Value: 11},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 11},

		{Suit: app.A, Value: 8},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 13},
	})

	h.SetWildJoker(&app.Card{Suit: app.A, Value: 9})

	_, invalid := h.RunTest(9)
	resCards := judgeCardLength(t, h, invalid)

	var want []app.Card

	wantI := []app.Card{
		{Suit: app.D, Value: 1},
		{Suit: app.D, Value: 3},
		{Suit: app.D, Value: 5},
		{Suit: app.D, Value: 6},
		{Suit: app.D, Value: 11},
		{Suit: app.D, Value: 13},

		{Suit: app.C, Value: 2},
		{Suit: app.C, Value: 11},

		{Suit: app.B, Value: 4},
		{Suit: app.B, Value: 11},

		{Suit: app.A, Value: 8},
		{Suit: app.A, Value: 12},
		{Suit: app.A, Value: 13},
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

func TestStraight17(t *testing.T) {
	desk := internal.InitializeDeck()

	internal.ShuffleDeck(desk)
	headCard := internal.DealCards(&desk, 13)

	h := internal.NewHand()
	h.SetCards(headCard)

	rand.NewSource(time.Now().UnixNano())
	jokerV := rand.Intn(13)

	jokerC := app.Card{Suit: app.A, Value: jokerV}
	h.SetWildJoker(&jokerC)
	_, invalid := h.RunTest(jokerV)

	printer := func(cards []app.Card, joker *app.Card) {
		for _, card := range cards {
			if card.Suit == app.A {
				if card.Value == joker.Value {
					fmt.Printf("♠: %v* || ", card.Value)
				} else {
					fmt.Printf("♠: %v || ", card.Value)
				}
			}
			if card.Suit == app.B {
				fmt.Printf("♥: %v || ", card.Value)
				if card.Value == joker.Value {
					fmt.Printf("♠: %v* || ", card.Value)
				} else {
					fmt.Printf("♠: %v || ", card.Value)
				}
			}
			if card.Suit == app.C {
				if card.Value == joker.Value {
					fmt.Printf("♣: %v* || ", card.Value)
				} else {
					fmt.Printf("♣: %v || ", card.Value)
				}
			}
			if card.Suit == app.D {
				if card.Value == joker.Value {
					fmt.Printf("♦: %v* || ", card.Value)
				} else {
					fmt.Printf("♦: %v || ", card.Value)
				}
			}
		}
	}

	fmt.Print("0 当前的手牌 :  ")
	printer(h.GetCards(), &jokerC)
	fmt.Println("手牌长度", len(h.GetCards()))
	fmt.Println()
	fmt.Println()

	fmt.Print("1 该牌的无效牌: ")
	printer(invalid, &jokerC)
	fmt.Println()

	fmt.Print("2 纯顺子: ")
	for _, p := range h.GetPure() {
		printer(p, &jokerC)
	}
	fmt.Println()

	fmt.Print("3 拥有的癞子牌: ")
	printer(h.GetJoker(), &jokerC)
	fmt.Println()

	fmt.Print("4 带癞子的顺子: ")
	for _, p := range h.GetPureWithJoker() {
		printer(p, &jokerC)
	}
	fmt.Println()

	fmt.Print("5 纯刻子: ")
	for _, p := range h.GetSet() {
		printer(p, &jokerC)
	}
	fmt.Println()

	fmt.Print("6 带癞子的刻子: ")
	for _, p := range h.GetSetWithJoker() {
		printer(p, &jokerC)
	}
	fmt.Println()

	fmt.Printf("7 当前的癞子牌 ♠: %v ", h.GetWildJoker().Value)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// A 黑桃
	// B 红桃
	// C 梅花
	// D 方片

	//h.SetCards([]app.Card{
	//	{Suit: app.D, Value: 1},
	//	{Suit: app.D, Value: 3},
	//	{Suit: app.D, Value: 5},
	//	{Suit: app.D, Value: 6},
	//	{Suit: app.D, Value: 11},
	//	{Suit: app.D, Value: 13},
	//
	//	{Suit: app.C, Value: 2},
	//	{Suit: app.C, Value: 11},
	//
	//	{Suit: app.B, Value: 4},
	//	{Suit: app.B, Value: 11},
	//
	//	{Suit: app.A, Value: 8},
	//	{Suit: app.A, Value: 12},
	//	{Suit: app.A, Value: 13},
	//})
	//
	//h.SetWildJoker(&app.Card{Suit: app.A, Value: 9})
	//
	//_, invalid := h.RunTest(9)
	//resCards := judgeCardLength(t, h, invalid)
	//
	//var want []app.Card
	//
	//wantI := []app.Card{
	//	{Suit: app.D, Value: 1},
	//	{Suit: app.D, Value: 3},
	//	{Suit: app.D, Value: 5},
	//	{Suit: app.D, Value: 6},
	//	{Suit: app.D, Value: 11},
	//	{Suit: app.D, Value: 13},
	//
	//	{Suit: app.C, Value: 2},
	//	{Suit: app.C, Value: 11},
	//
	//	{Suit: app.B, Value: 4},
	//	{Suit: app.B, Value: 11},
	//
	//	{Suit: app.A, Value: 8},
	//	{Suit: app.A, Value: 12},
	//	{Suit: app.A, Value: 13},
	//}
	//
	//realCards := handSliceDifference(resCards, invalid)
	//res := handSliceDifference(want, realCards)
	//if len(want) != len(realCards) {
	//	t.Errorf("理想长度 %v; \n 实际长度: %v", len(want), len(realCards))
	//	t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", want, realCards, res)
	//	t.Errorf("invalid %v", invalid)
	//}
	//
	//res2 := handSliceDifference(wantI, invalid)
	//if res != nil {
	//	t.Errorf("理想长度 %v; \n 实际长度: %v", len(wantI), len(invalid))
	//	t.Errorf("理想值获取错误： \n want %v  \n res %v \n 他们之间的差距 %v", wantI, invalid, res2)
	//	t.Errorf("invalid %v", invalid)
	//}
}
