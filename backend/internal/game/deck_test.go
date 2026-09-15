package game

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

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
	deck := newDeck()

	card, err := deck.drawCard()
	if card == HiddenCardValue {
		t.Fatal("DrawCard() returned the empty-deck value")
	}
	if err != nil {
		t.Fatalf("DrawCard() returned an error: %v", err)
	}
	if card < MinCardValue || card > MaxCardValue {
		t.Fatalf("DrawCard() returned an invalid card value: %d", card)
	}
	if got := deck.getNbCards(); got != 149 {
		t.Fatalf("deck contains %d cards after a draw, want 149", got)
	}
}

func TestDeckDrawCards(t *testing.T) {
	deck := newDeck()

	cards, err := deck.drawCards(12)
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
	deck := &deck{}

	if card, err := deck.drawCard(); card != HiddenCardValue || err == nil {
		t.Fatalf("DrawCard() on an empty deck returned %d, want %d and an error", card, HiddenCardValue)
	}
}

func TestDeckDrawCardsNotEnoughCards(t *testing.T) {
	deck := &deck{cards: []int{1, 2}}

	_, err := deck.drawCards(3)
	if err == nil {
		t.Fatalf("DrawCards(3) should have returned an error")
	}
	if got := deck.getNbCards(); got != 2 {
		t.Fatalf("deck contains %d cards after a failed draw, want 2", got)
	}
}

func TestDeckDrawCardsRejectsNonPositiveCount(t *testing.T) {
	for _, count := range []int{0, -1} {
		deck := newDeck()

		if _, err := deck.drawCards(count); err == nil {
			t.Fatalf("DrawCards(%d) returned no error", count)
		}
		if got := deck.getNbCards(); got != 150 {
			t.Fatalf("deck contains %d cards after DrawCards(%d), want 150", got, count)
		}
	}
}

func TestDeckIsEmpty(t *testing.T) {
	deck := newDeck()

	if deck.isEmpty() {
		t.Fatal("NewDeck() returned an empty deck")
	}

	// Draw all cards
	for i := 0; i < 150; i++ {
		if _, err := deck.drawCard(); err != nil {
			t.Fatalf("DrawCard() returned an error on draw %d: %v", i, err)
		}
	}

	if !deck.isEmpty() {
		t.Fatal("Deck should be empty after drawing all cards")
	}
}

func TestDeckGetNbCards(t *testing.T) {
	deck := newDeck()

	if got := deck.getNbCards(); got != 150 {
		t.Fatalf("NewDeck() contains %d cards, want 150", got)
	}

	// Draw some cards
	for i := 0; i < 10; i++ {
		if _, err := deck.drawCard(); err != nil {
			t.Fatalf("DrawCard() returned an error on draw %d: %v", i, err)
		}
	}

	if got := deck.getNbCards(); got != 140 {
		t.Fatalf("Deck contains %d cards after drawing 10, want 140", got)
	}
}

func TestDeckShufflePreservesCards(t *testing.T) {
	deck := newDeck()

	originalCounts := make(map[int]int)
	for _, card := range deck.cards {
		originalCounts[card]++
	}

	deck.shuffle()

	if got := deck.getNbCards(); got != 150 {
		t.Fatalf("Shuffle() changed the number of cards to %d, want 150", got)
	}

	shuffledCounts := make(map[int]int)
	for _, card := range deck.cards {
		shuffledCounts[card]++
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
		if got := shuffledCounts[card]; got != want {
			t.Errorf("Shuffle() produced %d cards with value %d, want %d", got, card, want)
		}
	}

	for card, want := range originalCounts {
		if got := shuffledCounts[card]; got != want {
			t.Errorf("Shuffle() changed count for card %d from %d to %d", card, want, got)
		}
	}
}

func TestDeckShuffleDoesNotPanicOnEmptyDeck(t *testing.T) {
	deck := &deck{}

	deck.shuffle()

	if got := deck.getNbCards(); got != 0 {
		t.Fatalf("Shuffle() on empty deck produced %d cards, want 0", got)
	}
}

func TestDeckShuffleDoesNotPanicOnSingleCard(t *testing.T) {
	deck := &deck{cards: []int{7}}

	deck.shuffle()

	if got := deck.getNbCards(); got != 1 {
		t.Fatalf("Shuffle() on single-card deck produced %d cards, want 1", got)
	}
	if deck.cards[0] != 7 {
		t.Fatalf("Shuffle() on single-card deck changed the card value to %d, want 7", deck.cards[0])
	}
}

