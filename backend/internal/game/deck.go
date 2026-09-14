package game

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	cards []int
}

func NewDeck() *Deck {
	deck := &Deck{
		cards: make([]int, 150),
	}
	expectedCounts := map[int]int{
		-2: 5,
		-1: 10,
		0:  15,
		1:  10,
		2:  10,
		3:  10,
		4:  10,
		5:  10,
		6:  10,
		7:  10,
		8:  10,
		9:  10,
		10: 10,
		11: 10,
		12: 10,
	}
	index := 0
	for card, count := range expectedCounts {
		for j := 0; j < count; j++ {
			deck.cards[index] = card
			index++
		}
	}
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
	return deck
}

func (d *Deck) DrawCard() (int, error) {
	if d.IsEmpty() {
		return HiddenCardValue, fmt.Errorf("deck is empty")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *Deck) DrawCards(n int) ([]int, error) {
	if n > len(d.cards) {
		return nil, fmt.Errorf("not enough cards to draw: have %d, want %d", len(d.cards), n)
	}
	if n <= 0 {
		return nil, fmt.Errorf("number of cards to draw must be positive: got %d", n)
	}
	drawnCards := d.cards[:n]
	d.cards = d.cards[n:]
	return drawnCards, nil
}

func (d *Deck) getNbCards() int {
	return len(d.cards)
}
