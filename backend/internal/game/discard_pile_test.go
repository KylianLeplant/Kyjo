package game

import "testing"

func TestDiscardPile(t *testing.T) {
	discard := NewDiscard()

	if card := discard.DiscardCard(); card != 13 {
		t.Fatalf("DiscardCard() on an empty pile returned %d, want 13", card)
	}

	discard.AddCard(7)
	discard.AddCard(3)
	if card := discard.DiscardCard(); card != 7 {
		t.Fatalf("first discarded card is %d, want 7", card)
	}
	if card := discard.DiscardCard(); card != 3 {
		t.Fatalf("second discarded card is %d, want 3", card)
	}
}
