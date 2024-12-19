package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"rummy-logic-v3/internal/logic"
	"rummy-logic-v3/pkg/app"
	"rummy-logic-v3/pkg/response"
	"strconv"
)

type Hand struct{}

func (h Hand) HandPrompt(c *gin.Context) {
	type Params struct {
		Cards [][]int `json:"cards"`
		Joker int     `json:"joker"`
	}

	var params Params

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err)
		return
	}

	var cards []app.Card
	for _, cardsRaw := range params.Cards {
		for _, rankRaw := range cardsRaw {
			rank := fmt.Sprintf("%x", rankRaw)
			var first, second string
			if len(rank) > 0 {
				first = string(rank[0]) // 第一个字符
			}
			if len(rank) > 1 {
				second = string(rank[1]) // 第二个字符
			}

			if second == "" {
				cards = append(cards, app.Card{Suit: app.D, Value: rankRaw})
				continue
			}

			v, _ := strconv.ParseInt(second, 16, 10)
			switch first {
			case "1":
				cards = append(cards, app.Card{Suit: app.C, Value: int(v)})
			case "2":
				cards = append(cards, app.Card{Suit: app.B, Value: int(v)})
			case "3":
				cards = append(cards, app.Card{Suit: app.A, Value: int(v)})
			case "4":
				if v == 15 {
					cards = append(cards, app.Card{Suit: app.JokerA, Value: int(v)})
				} else {
					cards = append(cards, app.Card{Suit: app.JokerB, Value: int(v)})
				}
			}
		}
	}

	prompt := logic.NewPrompt(cards, app.Card{Suit: app.A, Value: params.Joker})

	fmt.Println(prompt.GetResponse(cards))
	result := prompt.Calculate()

	response.Success(c, gin.H{
		"result": result,
	})
}

func (h Hand) RangeHand(c *gin.Context) {
	desk := logic.InitializeDeck()
	logic.ShuffleDeck(desk)
	cards := logic.DealCards(&desk, 13)

	for _, cc := range cards {
		fmt.Printf("{Suit: app.%s, Value: %d},\n", cc.Suit, cc.Value)
	}

	jokerV := rand.Intn(13) + 1
	prompt := logic.NewPrompt(cards, app.Card{Suit: app.A, Value: jokerV})
	result := prompt.Calculate()

	fmt.Println("jokerV:", jokerV)

	response.Success(c, gin.H{
		"myCards": prompt.GetResponse(cards)[0],
		"result":  result,
		"sysJoker": prompt.GetResponse([]app.Card{
			{Suit: app.A, Value: jokerV},
		}),
	})
}

func NewHand() *Hand {
	return &Hand{}
}

func RegHandRouter(r *gin.RouterGroup) {
	h := NewHand()
	r.POST("/hand/prompt", h.HandPrompt)
	r.GET("/hand/range", h.RangeHand)
}
