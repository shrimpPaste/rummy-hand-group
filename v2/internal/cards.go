package internal

import (
	"fmt"
	"rummy-group-v2/pkg/app"
)

// Hand 手牌
type Hand struct {
	cards     []app.Card
	joker     []app.Card
	valid     []app.Card
	invalid   []app.Card
	gap1Cards []app.Card // 间隙为1的牌
	suitCards map[string][]app.Card
}

// initHand 初始化手牌
func (h *Hand) initHand() {
	h.cards = []app.Card{
		{Value: 2, Suit: app.D},
		{Value: 3, Suit: app.D},
		{Value: 4, Suit: app.D},
		{Value: 5, Suit: app.D},
		{Value: 6, Suit: app.D},
		{Value: 7, Suit: app.D},

		{Value: 8, Suit: app.C},
		{Value: 1, Suit: app.C},

		{Value: 3, Suit: app.B},
		{Value: 7, Suit: app.B},
		//{Value: 13, Suit: app.B},
		{Value: 4, Suit: app.B}, //
		{Value: 5, Suit: app.B},

		{Value: 2, Suit: app.A},
		{Value: 12, Suit: app.A},
	}
}

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

// 手牌分组
func (h *Hand) groupCards(cards []app.Card) {
	for _, card := range cards {
		h.suitCards[card.Suit] = append(h.suitCards[card.Suit], card)
	}
}

func (h *Hand) Run() {
	// 初始化手牌
	h.initHand()
	// 分组
	h.groupCards(h.cards)
	// 找顺子
	h.findSequences()
	// 第一轮鉴定
	if !h.judgeIsHave2Seq() {
		_ = fmt.Errorf("没有找到2连顺子")
		return
	}
	// 找癞子
	//h.findJoker(2)

	fmt.Println("未处理的牌", h.suitCards)
	fmt.Println("有效牌", h.valid)
	fmt.Println("无效牌", h.invalid)
	fmt.Println("joker", h.joker)

	// 找间隙为1的牌
	//h.findGap1Cards()
}

func NewHand() *Hand {
	return &Hand{
		cards:     []app.Card{},
		joker:     []app.Card{},
		valid:     []app.Card{},
		invalid:   []app.Card{},
		gap1Cards: []app.Card{},
		suitCards: make(map[string][]app.Card, 4),
	}
}
