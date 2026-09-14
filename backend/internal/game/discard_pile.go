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

func (d *DiscardPile) AddCard(card int) {
	d.cards = append(d.cards, card)
}

func (d *DiscardPile) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *DiscardPile) TakeTopCard() (int, error) {
	if d.IsEmpty(){
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