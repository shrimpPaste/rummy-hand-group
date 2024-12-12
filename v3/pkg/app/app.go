package app

const (
	A      string = "A"      // 黑桃
	B      string = "B"      // 红桃
	C      string = "C"      // 梅花
	D      string = "D"      // 方片
	JokerA string = "JokerA" // 大鬼
	JokerB string = "JokerB" // 小鬼
)

const (
	NotStatus int = 0 // 无状态
	HT        int = 1 // Head and Tail 缺少首尾的牌  2 3 缺少 1 4
	Bd        int = 2 // blocked // 卡隆牌 3 5 缺少 4
)

type Card struct {
	Suit  string
	Value int
}

type GapCard struct {
	Cards       []Card
	Status      int
	Score       int
	JokerUseNum int
	// 0: 无状态
	// 1: 1 2 4 牌型 应得到 1 2 4
	// 2: 3 5 7 牌型 应得到 5 7
}
