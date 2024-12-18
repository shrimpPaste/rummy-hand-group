package main

import (
	"fmt"
	"rummy-logic-v3/internal/logic"
	"testing"
)

func TestLenIsTrue(t *testing.T) {
	hand := logic.NewHand()

	for i := range 1000000 {
		result := hand.ToTest()
		pure := result["pure"]
		pureWithJoker := result["pureWithJoker"]
		set := result["set"]
		setWithJoker := result["setWithJoker"]
		joker := result["joker"]
		invalid := result["invalid"]
		num := 0
		if len(pure) > 0 && pure[0] != 0 {
			num += len(pure)
		}
		if len(pureWithJoker) > 0 && pureWithJoker[0] != 0 {
			num += len(pureWithJoker)
		}
		if len(set) > 0 && set[0] != 0 {
			num += len(set)
		}
		if len(setWithJoker) > 0 && setWithJoker[0] != 0 {
			num += len(setWithJoker)
		}
		if len(joker) > 0 && joker[0] != 0 {
			num += len(joker)
		}

		if len(invalid) > 0 && invalid[0] != 0 {
			num += len(invalid)
		}

		if num != len(result["myCards"]) {
			fmt.Println("测试失败，测试次数为", i, num, len(result["myCards"]))
			for _, cc := range hand.GetCards() {
				fmt.Printf("{Suit: app.%s, Value: %d},\n", cc.Suit, cc.Value)
			}
			fmt.Println("Joker值是", hand.GetWildJoker().Value)
		}
	}
}
