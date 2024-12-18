package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rummy-logic-v3/internal/logic"
	"rummy-logic-v3/pkg/app"
	"rummy-logic-v3/pkg/response"
	"strconv"
)

type Hand struct{}

func (h Hand) HandPrompt(c *gin.Context) {
	type Params struct {
		Cards []int `json:"cards"`
		Joker int   `json:"joker"`
	}

	var params Params

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err)
		return
	}

	var cards []app.Card
	for _, rankRaw := range params.Cards {
		rank := fmt.Sprintf("%x", rankRaw)

		var first, second string
		if len(rank) > 0 {
			first = string(rank[0]) // 第一个字符
		}
		if len(rank) > 1 {
			second = string(rank[1]) // 第二个字符
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
		default:
			cards = append(cards, app.Card{Suit: app.D, Value: rankRaw})
		}
	}

	prompt := logic.NewPrompt(cards, app.Card{Suit: app.A, Value: params.Joker})
	result := prompt.Calculate()

	response.Success(c, gin.H{
		"result": result,
	})
}

func NewHand() *Hand {
	return &Hand{}
}

func RegHandRouter(r *gin.RouterGroup) {
	h := NewHand()
	r.POST("/hand/prompt", h.HandPrompt)
}
