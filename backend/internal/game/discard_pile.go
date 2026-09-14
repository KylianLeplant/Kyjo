package game

import "fmt"

type DiscardPile struct {
	cards []int
}

func NewDiscard() *DiscardPile {
	return &DiscardPile{
		cards: []int{},
	}
}

func (d *DiscardPile) AddCard(card int) error {
	if card < -2 || card > 12 {
		return fmt.Errorf("invalid card value: %d (must be between -2 and 12)", card)
	}
	d.cards = append(d.cards, card)
	return nil
}

func (d *DiscardPile) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *DiscardPile) TakeTopCard() (int, error) {
	if d.IsEmpty() {
		return 13, fmt.Errorf("discard pile is empty")
	}
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card, nil
}

func (d *DiscardPile) TakeAllCards() []int {
	cards := d.cards
	d.cards = []int{}
	return cards
}

func (d *DiscardPile) GetTopCard() (int, error) {
	if d.IsEmpty() {
		return 13, fmt.Errorf("discard pile is empty")
	}
	return d.cards[len(d.cards)-1], nil
}
