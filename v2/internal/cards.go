package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"rummy-group-v2/pkg/app"
	"time"
)

// Hand 手牌
type Hand struct {
	cards         []app.Card
	joker         []app.Card
	valid         []app.Card
	pure          [][]app.Card
	pureWithJoker [][]app.Card
	set           [][]app.Card
	setWithJoker  [][]app.Card
	invalid       []app.Card
	gap1Cards     []app.Card // 间隙为1的牌
	wild          *app.Card  // 当前的Joker牌
	suitCards     map[string][]app.Card
}

func (h *Hand) SetCards(cards []app.Card) {
	h.cards = cards
}

func (h *Hand) GetCards() []app.Card {
	return h.cards
}
func (h *Hand) GetPure() [][]app.Card {
	return h.pure
}
func (h *Hand) GetPureWithJoker() [][]app.Card {
	return h.pureWithJoker
}

func (h *Hand) GetSet() [][]app.Card {
	return h.set
}

func (h *Hand) GetSetWithJoker() [][]app.Card {
	return h.setWithJoker
}

// initHand 初始化手牌
func (h *Hand) initHand() {
	//h.cards = []app.Card{
	//	{Value: 2, Suit: app.D},
	//	{Value: 4, Suit: app.D},
	//	{Value: 5, Suit: app.D},
	//	{Value: 6, Suit: app.D},
	//	{Value: 7, Suit: app.D},
	//
	//	{Value: 8, Suit: app.C},
	//	{Value: 1, Suit: app.C},
	//
	//	{Value: 3, Suit: app.B},
	//	{Value: 7, Suit: app.B},
	//	{Value: 13, Suit: app.B},
	//	{Value: 1, Suit: app.B},
	//
	//	{Value: 2, Suit: app.A},
	//	{Value: 12, Suit: app.A},
	//}
	//h.cards = []app.Card{
	//	{Value: 2, Suit: app.D},
	//	{Value: 8, Suit: app.D},
	//
	//	{Value: 2, Suit: app.C},
	//	{Value: 7, Suit: app.C},
	//	{Value: 8, Suit: app.C},
	//	{Value: 9, Suit: app.C},
	//	{Value: 10, Suit: app.C},
	//	{Value: 1, Suit: app.C},
	//
	//	{Value: 4, Suit: app.B},
	//	{Value: 6, Suit: app.B},
	//	{Value: 13, Suit: app.B},
	//	{Value: 1, Suit: app.B},
	//
	//	{Value: 1, Suit: app.A},
	//}
	//fmt.Println(len(h.cards))

	h.cards = []app.Card{
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
}

func (h *Hand) Run(r *gin.Engine) {
	// 初始化手牌
	//h.initHand()
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
	h.findInvalidJoker(6)
	if len(h.joker) < 2 && !h.judgeIsHave1Seq() {
		fmt.Println("没有找到足够的癞子牌支持组成第二组顺子")
		return
	}
	// 有癞子找间隙牌
	h.findInvalidGap1Cards()

	// 找刻子
	h.find111Cards()

	var result []int
	for _, c := range h.valid {
		switch c.Suit {
		case app.A:
			result = append(result, c.Value+48)
		case app.B:
			result = append(result, c.Value+32)
		case app.C:
			result = append(result, c.Value+16)
		case app.D:
			result = append(result, c.Value)
		}
		//hexStr := fmt.Sprintf("%X", c.Value*q)
		//fmt.Println("十六进制", hexStr)

		// rtn = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
		// rtn = [17, 18, 19, 20,21,22,23,24,25,26, 27, 28, 29] // 14, 15, 16,
		//rtn = [33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45] // 30, 31, 32,
		//rtn = [49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61] // 46, 47, 48,
		//rtn = [78, 79] // 62, 63, 64,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, result)
		return
	})

	//fmt.Println("未处理的牌", h.suitCards)
	fmt.Println("有效牌", h.valid)
	fmt.Println("无效牌", h.invalid)
	fmt.Println("joker", h.joker)

	fmt.Println("=================")
	fmt.Println(len(h.valid)+len(h.invalid)+len(h.joker) == len(h.cards))
	//fmt.Println("h.valid", len(h.valid))
	//fmt.Println("h.invalid", len(h.invalid))
	//fmt.Println("h.joker", len(h.joker))
	//fmt.Println("len(h.cards)", len(h.cards))

	// 找间隙为1的牌
	//h.findInvalidGap1Cards()
}

func (h *Hand) SetWildJoker(card *app.Card) {
	h.wild = card
}

func (h *Hand) GetWildJoker() *app.Card {
	return h.wild
}

func (h *Hand) GetJoker() []app.Card {
	return h.joker
}

func (h *Hand) RunTest(wild int) ([]app.Card, []app.Card) {
	h.findValidGap1Cards()

	//// 分组
	//h.groupCards(h.suitCards, h.cards)
	//// 找顺子
	//h.findSequences()
	//// 第一轮鉴定
	//if !h.judgeIsHave1Seq() {
	//	//fmt.Println("Waring::没有找到一个无赖字的同花顺子")
	//	return h.valid, h.invalid
	//}
	//// 找癞子
	//h.findInvalidJoker(wild)
	//
	//if len(h.joker) < 1 && !h.judgeIsHave2Seq() {
	//	//fmt.Println("Waring::没有找到足够的癞子牌支持组成第二组顺子")
	//	return h.valid, h.invalid
	//}
	//
	//h.findValidGap1Cards()
	//
	//// 有癞子找间隙牌
	//h.findInvalidGap1Cards()
	//
	//// 找刻子
	//h.find111Cards()
	//
	//// 再次找间隙为1的牌，因为找刻子有可能腾出来joker
	//h.findInvalidGap1Cards()

	return h.valid, h.invalid
}

func (h *Hand) RunWeb(r *gin.Engine) {
	//h.findValidGap1Cards()

	r.GET("/", func(c *gin.Context) {
		desk := InitializeDeck()

		ShuffleDeck(desk)
		headCard := DealCards(&desk, 13)
		h.SetCards(headCard)

		h.pure = [][]app.Card{}
		h.pureWithJoker = [][]app.Card{}
		h.set = [][]app.Card{}
		h.setWithJoker = [][]app.Card{}
		h.invalid = []app.Card{}
		h.valid = []app.Card{}
		h.joker = []app.Card{}

		rand.NewSource(time.Now().UnixNano())
		jokerV := rand.Intn(12) + 1

		jokerC := app.Card{Suit: app.A, Value: jokerV}
		h.SetWildJoker(&jokerC)

		suitCards := make(map[string][]app.Card, 4)
		h.groupCards(suitCards, h.GetCards())

		var handCards []app.Card
		for _, cards := range suitCards {
			handCards = append(handCards, cards...)
		}
		myCards := getCardsResult(handCards)

		var calcCardsRaw []app.Card
		var pureCards, pureWithJoker, set, setWithJoker, invalidCards []app.Card
		h.findValidGap1Cards()

		for _, p := range h.pure {
			calcCardsRaw = append(calcCardsRaw, p...)
			pureCards = append(pureCards, p...)
		}
		for _, cards := range h.pureWithJoker {
			calcCardsRaw = append(calcCardsRaw, cards...)
			pureWithJoker = append(pureWithJoker, cards...)
		}
		for _, cards := range h.set {
			calcCardsRaw = append(calcCardsRaw, cards...)
			set = append(set, cards...)
		}
		for _, cards := range h.setWithJoker {
			calcCardsRaw = append(calcCardsRaw, cards...)
			setWithJoker = append(setWithJoker, cards...)
		}
		for _, cards := range h.invalid {
			calcCardsRaw = append(calcCardsRaw, cards)
			invalidCards = append(invalidCards, cards)
		}

		c.JSON(200, gin.H{
			"myCards":       myCards,
			"calcCards":     getCardsResult(calcCardsRaw),
			"pure":          getCardsResult(pureCards),
			"pureWithJoker": getCardsResult(pureWithJoker),
			"set":           getCardsResult(set),
			"setWithJoker":  getCardsResult(setWithJoker),
			"invalid":       getCardsResult(invalidCards),
			"joker":         getCardsResult(h.GetJoker()),
			"sysJoker":      getCardsResult([]app.Card{*h.GetWildJoker()}),
		})
		return
	})
	return
}

func getCardsResult(cards []app.Card) []int {
	var myCards []int
	for _, c := range cards {
		if c.Suit == app.A {
			myCards = append(myCards, c.Value+48)
		} else if c.Suit == app.B {
			myCards = append(myCards, c.Value+32)
		} else if c.Suit == app.C {
			myCards = append(myCards, c.Value+16)
		} else if c.Suit == app.D {
			myCards = append(myCards, c.Value)
		} else if c.Suit == app.JokerA {
			myCards = append(myCards, 79)
		} else if c.Suit == app.JokerB {
			myCards = append(myCards, 78)
		}
	}
	return myCards
}

func NewHand() *Hand {
	return &Hand{
		cards:         []app.Card{},
		pure:          make([][]app.Card, 0),
		pureWithJoker: make([][]app.Card, 0),
		set:           make([][]app.Card, 0),
		setWithJoker:  make([][]app.Card, 0),
		joker:         []app.Card{},
		valid:         []app.Card{},
		invalid:       []app.Card{},
		gap1Cards:     []app.Card{},
		suitCards:     make(map[string][]app.Card, 4),
	}
}
