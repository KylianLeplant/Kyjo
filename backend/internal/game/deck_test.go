package game

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if got := deck.getNbCards(); got != 150 {
		t.Fatalf("NewDeck() contains %d cards, want 150", got)
	}

	counts := make(map[int]int)
	for _, card := range deck.cards {
		counts[card]++
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

	for card, want := range expectedCounts {
		if got := counts[card]; got != want {
			t.Errorf("NewDeck() contains %d cards with value %d, want %d", got, card, want)
		}
	}
}

func TestDeckDrawCard(t *testing.T) {
	deck := NewDeck()

	card, err := deck.DrawCard()
	if card == HiddenCardValue {
		t.Fatal("DrawCard() returned the empty-deck value")
	}
	if err != nil {
		t.Fatalf("DrawCard() returned an error: %v", err)
	}
	if card < MinCardValue|| card > MaxCardValue {
		t.Fatalf("DrawCard() returned an invalid card value: %d", card)
	}
	if got := deck.getNbCards(); got != 149 {
		t.Fatalf("deck contains %d cards after a draw, want 149", got)
	}
}

func TestDeckDrawCards(t *testing.T) {
	deck := NewDeck()

	cards, err := deck.DrawCards(12)
	if err != nil {
		t.Fatalf("DrawCards(12) returned an error: %v", err)
	}
	if len(cards) != 12 {
		t.Fatalf("DrawCards(12) returned %d cards, want 12", len(cards))
	}
	for _, card := range cards {
		if card == HiddenCardValue {
			t.Fatal("DrawCards() returned the empty-deck value")
		}
		if card < MinCardValue || card > MaxCardValue {
			t.Fatalf("DrawCards() returned an invalid card value: %d", card)
		}
	}
	if got := deck.getNbCards(); got != 138 {
		t.Fatalf("deck contains %d cards after drawing 12, want 138", got)
	}
}

func TestEmptyDeck(t *testing.T) {
	deck := &Deck{}

	if card, err := deck.DrawCard(); card != HiddenCardValue || err == nil {
		t.Fatalf("DrawCard() on an empty deck returned %d, want %d and an error", card, HiddenCardValue)
	}
}

func TestDeckDrawCardsNotEnoughCards(t *testing.T) {
	deck := &Deck{cards: []int{1, 2}}

	_, err := deck.DrawCards(3)
	if err == nil {
		t.Fatalf("DrawCards(3) should have returned an error")
	}
	if got := deck.getNbCards(); got != 2 {
		t.Fatalf("deck contains %d cards after a failed draw, want 2", got)
	}
}
