package game

import (
	"math/rand"
)

type Deck struct {
	cards []int
}

func NewDeck() *Deck {
	deck := &Deck{
		cards: make([]int, 150),
	}
	for i := 0; i < 5; i++ {
		deck.cards[i] = -2
	}
	for i := 5; i < 20; i++ {
		deck.cards[i] = 0
	}
	tab := []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i, value := range tab {
		for j := 0; j < 10; j++ {
			deck.cards[20+i*10+j] = value
		}
	}
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
	return deck
}

func (d *Deck) DrawCard() int {
	if d.IsEmpty() {
		return 13 // 13 indicates that the deck is empty
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *Deck) DrawCards(n int) []int {
	if n > len(d.cards) {
		return nil // Not enough cards to draw
	}
	drawnCards := d.cards[:n]
	d.cards = d.cards[n:]
	return drawnCards
}

func (d *Deck) getNbCards() int {
	return len(d.cards)
}

