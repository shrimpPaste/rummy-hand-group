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
		{Value: 4, Suit: app.D},
		{Value: 5, Suit: app.D},
		{Value: 6, Suit: app.D},
		{Value: 7, Suit: app.D},

		{Value: 8, Suit: app.C},
		{Value: 1, Suit: app.C},

		{Value: 3, Suit: app.B},
		{Value: 7, Suit: app.B},
		{Value: 13, Suit: app.B},
		{Value: 1, Suit: app.B},

		{Value: 2, Suit: app.A},
		{Value: 12, Suit: app.A},
	}
}

func (h *Hand) Run() {
	// 初始化手牌
	h.initHand()
	// 分组
	h.groupCards(h.suitCards, h.cards)
	// 找顺子
	h.findSequences()
	// 第一轮鉴定
	if !h.judgeIsHave1Seq() {
		fmt.Println("没有找到一个无赖字的同花顺子")
		return
	}
	// 找癞子
	h.findInvalidJoker(2)
	if len(h.joker) < 2 && !h.judgeIsHave1Seq() {
		fmt.Println("没有找到足够的癞子牌支持组成第二组顺子")
		return
	}
	// 有癞子找间隙牌
	h.findGap1Cards()

	//fmt.Println("未处理的牌", h.suitCards)
	fmt.Println("有效牌", h.valid)
	fmt.Println("无效牌", h.invalid)
	//fmt.Println("joker", h.joker)

	// 找间隙为1的牌
	h.findGap1Cards()
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
