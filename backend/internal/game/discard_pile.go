package game

import "fmt"

type discardPile struct {
	cards []int
}

func newDiscardPile() *discardPile {
	return &discardPile{
		cards: []int{},
	}
}

func (d *discardPile) addCard(card int) error {
	if !isValidCardValue(card) {
		return fmt.Errorf("invalid card value: %d (must be between %d and %d)", card, MinCardValue, MaxCardValue)
	}
	d.cards = append(d.cards, card)
	return nil
}

func (d *discardPile) isEmpty() bool {
	return len(d.cards) == 0
}

func (d *discardPile) takeTopCard() (int, error) {
	if d.isEmpty() {
		return HiddenCardValue, fmt.Errorf("discard pile is empty")
	}
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card, nil
}

func (d *discardPile) takeAllCards() []int {
	cards := d.cards
	d.cards = []int{}
	return cards
}

func (d *discardPile) getTopCard() (int, error) {
	if d.isEmpty() {
		return HiddenCardValue, fmt.Errorf("discard pile is empty")
	}
	return d.cards[len(d.cards)-1], nil
}
