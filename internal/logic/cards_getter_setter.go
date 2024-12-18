package logic

import "rummy-logic-v3/pkg/app"

func (h *Hand) SetCards(cards []app.Card) {
	h.cards = cards
}

func (h *Hand) Cards() []app.Card {
	return h.cards
}

func (h *Hand) SetWildJoker(card app.Card) {
	h.wild = card
}

func (h *Hand) GetWildJoker() app.Card {
	return h.wild
}
