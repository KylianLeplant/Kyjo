package game

type DiscardPile struct {
	cards []int
}

func NewDiscard() *DiscardPile {
	return &DiscardPile{
		cards: []int{},
	}
}

func (d *DiscardPile) AddCard(card int) {
	d.cards = append(d.cards, card)
}

func (d *DiscardPile) DiscardCard() int {
	if len(d.cards) == 0 {
		return 13 // 13 indicates that the discard pile is empty
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}
