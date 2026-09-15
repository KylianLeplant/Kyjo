package game

import (
	"fmt"
	"math/rand"
)

type deck struct {
	cards []int
}

func newDeck() *deck {
	deck := &deck{
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
	deck.shuffle()
	return deck
}

func (d *deck) drawCard() (int, error) {
	if d.isEmpty() {
		return HiddenCardValue, fmt.Errorf("deck is empty")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

func (d *deck) isEmpty() bool {
	return len(d.cards) == 0
}

func (d *deck) drawCards(n int) ([]int, error) {
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

func (d *deck) getNbCards() int {
	return len(d.cards)
}

func (d *deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
