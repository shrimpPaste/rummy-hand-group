package stright

import (
	"reflect"
	"sort"
	"testing"
)

// TestStraight 手牌是 1 1 2 3 4
func TestStraight(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: A, Value: 1}, Card{Suit: A, Value: 1}, Card{Suit: A, Value: 2}, Card{Suit: A, Value: 3}, Card{Suit: A, Value: 4},
	}

	wantValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 1}, Card{Suit: A, Value: 2}, Card{Suit: A, Value: 3}, Card{Suit: A, Value: 4},
		},
	}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 1},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是 1 2 3 4
	// 无效牌是1
	valid, invalid := groupCards(player, 5)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

// TestStraight2 手牌是 1 1 2 4 5
func TestStraight2(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: A, Value: 1}, Card{Suit: A, Value: 1}, Card{Suit: A, Value: 2}, Card{Suit: A, Value: 4}, Card{Suit: A, Value: 5},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 1}, Card{Suit: A, Value: 1}, Card{Suit: A, Value: 2}, Card{Suit: A, Value: 4}, Card{Suit: A, Value: 5},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

// TestStraight3 手牌是 1 1 2 4 5
//func TestStraight3(t *testing.T) {
//	// 找顺子
//
//	// 手牌 1 1 2 3 4
//	player := Hand{
//		Card{Suit: A, Value: 2}, Card{Suit: A, Value: 8},
//		Card{Suit: B, Value: 5}, Card{Suit: B, Value: 6}, Card{Suit: B, Value: 8}, Card{Suit: B, Value: 13},
//		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 7}, Card{Suit: C, Value: 8}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
//		Card{Suit: D, Value: 2}, Card{Suit: D, Value: 7},
//	}
//
//	wantValid := Hand{
//		Card{Suit: C, Value: 7}, Card{Suit: C, Value: 8}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
//	}
//	wantInValid := Hand{
//		Card{Suit: A, Value: 2}, Card{Suit: A, Value: 8},
//		Card{Suit: B, Value: 5}, Card{Suit: B, Value: 6}, Card{Suit: B, Value: 8}, Card{Suit: B, Value: 13}, Card{Suit: C, Value: 1},
//		Card{Suit: D, Value: 2}, Card{Suit: D, Value: 7},
//	}
//	// 癞子为5的时候，
//	// 有效牌希望结果是
//	// 无效牌是 1 1 2 4 5
//	valid, invalid := groupCards(player, 6)
//
//	if !reflect.DeepEqual(wantValid, valid) {
//		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, invalid)
//	}
//
//	if !reflect.DeepEqual(wantInValid, invalid) {
//		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
//	}
//}

func TestStraight4(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 7},
		Card{Suit: C, Value: 8}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 7}, Card{Suit: C, Value: 8}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight5(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3},
		Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 13}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 1},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight6(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4},
		Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 13}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 1},
			Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight7(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5},
		Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 13}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 1},
			Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight8(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 1}, Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5}, Card{Suit: C, Value: 5},
		Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 13}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 1},
			Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 5}, Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10},
		},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight9(t *testing.T) {
	// 找顺子

	// 手牌 1 1 2 3 4
	player := Hand{
		Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5}, Card{Suit: C, Value: 5},
		Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 11}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 2}, Card{Suit: C, Value: 3}, Card{Suit: C, Value: 4}, Card{Suit: C, Value: 5},
			Card{Suit: C, Value: 9}, Card{Suit: C, Value: 10}, Card{Suit: C, Value: 11}, Card{Suit: C, Value: 12}, Card{Suit: C, Value: 13},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {Card{Suit: C, Value: 5}},
	}
	// 癞子为5的时候，
	// 有效牌希望结果是
	// 无效牌是 1 1 2 4 5
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight10(t *testing.T) {
	// 手牌 1
	player := Hand{
		Card{Suit: C, Value: 1},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		C: {Card{Suit: C, Value: 1}},
	}

	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight11(t *testing.T) {
	// 手牌 1 2
	player := Hand{
		Card{Suit: C, Value: 1},
		Card{Suit: C, Value: 2},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 1},
			Card{Suit: C, Value: 2},
		},
	}

	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight12(t *testing.T) {
	// 手牌 2 4 5
	player := Hand{
		Card{Suit: C, Value: 2},
		Card{Suit: C, Value: 4},
		Card{Suit: C, Value: 5},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 2},
			Card{Suit: C, Value: 4},
			Card{Suit: C, Value: 5},
		},
	}

	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight13(t *testing.T) {
	// 手牌 3 3 4 5
	player := Hand{
		Card{Suit: C, Value: 3},
		Card{Suit: C, Value: 3},
		Card{Suit: C, Value: 4},
		Card{Suit: C, Value: 5},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 3},
			Card{Suit: C, Value: 4},
			Card{Suit: C, Value: 5},
		},
	}
	wantInValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 3},
		},
	}

	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight14(t *testing.T) {
	// 手牌 1 2 3 4 5, 癞子为 6
	player := Hand{
		Card{Suit: C, Value: 1},
		Card{Suit: C, Value: 2},
		Card{Suit: C, Value: 3},
		Card{Suit: C, Value: 4},
		Card{Suit: C, Value: 5},
	}

	// 用癞子6来补全顺子
	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 1},
			Card{Suit: C, Value: 2},
			Card{Suit: C, Value: 3},
			Card{Suit: C, Value: 4},
			Card{Suit: C, Value: 5},
		},
	}
	wantInValid := map[Suit]Hand{}

	// 此处假设癞子是6
	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func TestStraight15(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: C, Value: 1},
		Card{Suit: D, Value: 2},
		Card{Suit: A, Value: 3},
		Card{Suit: D, Value: 4},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 3},
		},
		C: {
			Card{Suit: C, Value: 1},
		},
		D: {
			Card{Suit: D, Value: 2},
			Card{Suit: D, Value: 4},
		},
	}

	valid, invalid := groupCards(player, 6)

	if !reflect.DeepEqual(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !reflect.DeepEqual(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

// 找刻子
func TestStraight16(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: C, Value: 1},
		Card{Suit: D, Value: 2},
		Card{Suit: D, Value: 4},
		Card{Suit: D, Value: 1},

		Card{Suit: A, Value: 3},
		Card{Suit: A, Value: 5},
		Card{Suit: A, Value: 4},

		Card{Suit: A, Value: 1},
	}

	wantValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 3},
			Card{Suit: A, Value: 4},
			Card{Suit: A, Value: 5},

			Card{Suit: C, Value: 1},
			Card{Suit: D, Value: 1},
			Card{Suit: A, Value: 1},
		},
	}
	wantInValid := map[Suit]Hand{
		A: nil,
		C: nil,
		D: {
			Card{Suit: D, Value: 2},
			Card{Suit: D, Value: 4},
		},
	}

	valid, invalid := groupCards(player, 6)

	valid, invalid = find111Card(valid, invalid)

	if !compareMaps(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, 得到:%v", wantInValid, invalid)
	}
}

func compareMaps(want, got map[Suit]Hand) bool {
	// 检查 map 长度是否相等
	if len(want) != len(got) {
		return false
	}

	// 遍历 want 中的键，比较每个键的值
	for key, wantHand := range want {
		gotHand, exists := got[key]
		if !exists {
			return false
		}

		// 创建两个切片来存储排序后的值
		wantValues := make([]int, len(wantHand))
		gotValues := make([]int, len(gotHand))

		// 将手牌中的值复制到切片中
		for i, card := range wantHand {
			wantValues[i] = card.Value
		}
		for i, card := range gotHand {
			gotValues[i] = card.Value
		}

		// 对值进行排序
		sort.Ints(wantValues)
		sort.Ints(gotValues)

		// 比较排序后的值切片是否相等
		if !reflect.DeepEqual(wantValues, gotValues) {
			return false
		}
	}

	return true
}

// 找刻子
func TestStraight17(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: D, Value: 2},
		Card{Suit: D, Value: 7},

		Card{Suit: C, Value: 2},
		Card{Suit: C, Value: 4},
		Card{Suit: C, Value: 7},
		Card{Suit: C, Value: 11},
		Card{Suit: C, Value: 13},

		Card{Suit: B, Value: 5},
		Card{Suit: B, Value: 6},
		Card{Suit: B, Value: 8},
		Card{Suit: B, Value: 13},

		Card{Suit: A, Value: 2},
		Card{Suit: A, Value: 8},
	}

	wantValid := map[Suit]Hand{}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 2},
			Card{Suit: A, Value: 8},
		},
		B: {
			Card{Suit: B, Value: 5},
			Card{Suit: B, Value: 6},
			Card{Suit: B, Value: 8},
			Card{Suit: B, Value: 13},
		},
		C: {
			Card{Suit: C, Value: 2},
			Card{Suit: C, Value: 4},
			Card{Suit: C, Value: 7},
			Card{Suit: C, Value: 11},
			Card{Suit: C, Value: 13},
		},
		D: {
			Card{Suit: D, Value: 2},
			Card{Suit: D, Value: 7},
		},
	}

	valid, invalid := groupCards(player, 0)

	hasStraight := 0
	for _, hand := range valid {
		if len(hand) >= 3 {
			hasStraight += 1
		}
	}
	if hasStraight >= 2 {
		valid, invalid = find111Card(valid, invalid)
	}

	if !compareMaps(wantValid, valid) {
		t.Errorf("希望得到正确牌组:%v, \n 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("希望得到废弃牌组:%v, \n 得到:%v", wantInValid, invalid)
	}
}

func TestStraight18(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: D, Value: 2},
		Card{Suit: D, Value: 8},

		Card{Suit: C, Value: 2},
		Card{Suit: C, Value: 7},
		Card{Suit: C, Value: 8},
		Card{Suit: C, Value: 9},
		Card{Suit: C, Value: 10},
		Card{Suit: C, Value: 1},

		Card{Suit: B, Value: 4},
		Card{Suit: B, Value: 6},
		Card{Suit: B, Value: 13},
		Card{Suit: B, Value: 1},

		Card{Suit: A, Value: 1},
	}

	wantValid := map[Suit]Hand{
		C: {
			Card{Suit: C, Value: 7},
			Card{Suit: C, Value: 8},
			Card{Suit: C, Value: 9},
			Card{Suit: C, Value: 10},
		},
	}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 1},
		},
		B: {
			Card{Suit: B, Value: 4},
			Card{Suit: B, Value: 6},
			Card{Suit: B, Value: 13},
			Card{Suit: B, Value: 1},
		},
		C: {
			Card{Suit: C, Value: 2},
			Card{Suit: C, Value: 1},
		},
		D: {
			Card{Suit: D, Value: 2},
			Card{Suit: D, Value: 8},
		},
	}

	valid, invalid := groupCards(player, 9)

	// 判断如果没有两个顺子，则不找刻子
	hasStraight := 0
	for _, hand := range valid {
		if len(hand) >= 3 {
			hasStraight += 1
		}
	}
	if hasStraight >= 2 {
		valid, invalid = find111Card(valid, invalid)
	}

	if !compareMaps(wantValid, valid) {
		t.Errorf("\n 希望得到正确牌组:%v, \n 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("\n 希望得到废弃牌组:%v, \n 得到:%v", wantInValid, invalid)
	}
}

// 找顺子和一个有癞子的顺子
func TestStraight19(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: D, Value: 2},
		Card{Suit: D, Value: 4},
		Card{Suit: D, Value: 5},
		Card{Suit: D, Value: 6},
		Card{Suit: D, Value: 7},

		Card{Suit: C, Value: 8},
		Card{Suit: C, Value: 7},
		Card{Suit: C, Value: 6},
		Card{Suit: C, Value: 1},

		Card{Suit: B, Value: 3},
		Card{Suit: B, Value: 7},
		Card{Suit: B, Value: 13},
		//Card{Suit: B, Value: 2},
		Card{Suit: B, Value: 1},

		Card{Suit: A, Value: 2},
		Card{Suit: A, Value: 12},
	}

	wantValid := map[Suit]Hand{
		D: {
			Card{Suit: D, Value: 4},
			Card{Suit: D, Value: 5},
			Card{Suit: D, Value: 6},
			Card{Suit: D, Value: 7},
		},
		B: {
			Card{Suit: D, Value: 2},
			Card{Suit: B, Value: 13},
			Card{Suit: B, Value: 1},
		},
		C: {
			Card{Suit: C, Value: 8},
			Card{Suit: C, Value: 7},
			Card{Suit: C, Value: 6},
		},
	}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 12},
		},
		B: {
			Card{Suit: B, Value: 3},
			Card{Suit: B, Value: 7},
		},
		C: {
			Card{Suit: C, Value: 1},
		},
		D: {},
	}

	valid, invalid := groupCards(player, 2)

	// 判断如果没有两个顺子，则不找刻子
	hasStraight := 0
	for _, hand := range valid {
		if len(hand) >= 3 {
			hasStraight += 1
		}
	}
	if hasStraight >= 2 {
		valid, invalid = find111Card(valid, invalid)
	}

	if !compareMaps(wantValid, valid) {
		t.Errorf("\n 希望得到正确牌组:%v, \n 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("\n 希望得到废弃牌组:%v, \n 实际得到有用牌组:%v", wantInValid, invalid)
	}
}

// 找顺子和一个有癞子的顺子
func TestStraight20(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: D, Value: 1},
		Card{Suit: D, Value: 3},
		Card{Suit: D, Value: 4},
		Card{Suit: D, Value: 7},

		Card{Suit: C, Value: 5},
		Card{Suit: C, Value: 7},
		Card{Suit: C, Value: 6},

		Card{Suit: B, Value: 8},
		Card{Suit: B, Value: 11},
		Card{Suit: B, Value: 10},

		Card{Suit: A, Value: 2},
		Card{Suit: A, Value: 12},
		Card{Suit: A, Value: 1},
	}

	wantValid := map[Suit]Hand{
		B: {
			Card{Suit: B, Value: 8},
			Card{Suit: B, Value: 11},
			Card{Suit: B, Value: 10},
		},
		C: {
			Card{Suit: C, Value: 5},
			Card{Suit: C, Value: 7},
			Card{Suit: C, Value: 6},
		},
	}
	wantInValid := map[Suit]Hand{
		A: {
			Card{Suit: A, Value: 2},
			Card{Suit: A, Value: 12},
			Card{Suit: A, Value: 1},
		},
		B: {},
		D: {
			Card{Suit: D, Value: 1},
			Card{Suit: D, Value: 3},
			Card{Suit: D, Value: 4},
			Card{Suit: D, Value: 7},
		},
	}

	valid, invalid := groupCards(player, 8)

	// 判断如果没有两个顺子，则不找刻子
	hasStraight := 0
	for _, hand := range valid {
		if len(hand) >= 3 {
			hasStraight += 1
		}
	}
	if hasStraight >= 2 {
		valid, invalid = find111Card(valid, invalid)
	}

	if !compareMaps(wantValid, valid) {
		t.Errorf("\n 希望得到正确牌组:%v, \n 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("\n 希望得到废弃牌组:%v, \n 实际得到有用牌组:%v", wantInValid, invalid)
	}
}

// 找顺子和一个有癞子的顺子
func TestStraight21(t *testing.T) {
	// 手牌含有不同花色的牌
	player := Hand{
		Card{Suit: D, Value: 1},
		Card{Suit: D, Value: 3},
		Card{Suit: D, Value: 4},
		Card{Suit: D, Value: 5},
		Card{Suit: D, Value: 6},
	}

	wantValid := map[Suit]Hand{
		D: {
			Card{Suit: D, Value: 1},
			Card{Suit: D, Value: 3},
			Card{Suit: D, Value: 4},
			Card{Suit: D, Value: 5},
			Card{Suit: D, Value: 6},
		},
	}
	wantInValid := map[Suit]Hand{}

	valid, invalid := groupCards(player, 6)

	// 判断如果没有两个顺子，则不找刻子
	hasStraight := 0
	for _, hand := range valid {
		if len(hand) >= 3 {
			hasStraight += 1
		}
	}
	if hasStraight >= 2 {
		valid, invalid = find111Card(valid, invalid)
	}

	if !compareMaps(wantValid, valid) {
		t.Errorf("\n 希望得到正确牌组:%v, \n 得到:%v", wantValid, valid)
	}

	if !compareMaps(wantInValid, invalid) {
		t.Errorf("\n 希望得到废弃牌组:%v, \n 实际得到有用牌组:%v", wantInValid, invalid)
	}
}
