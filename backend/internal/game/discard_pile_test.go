package game

import "testing"

func TestDiscardPile(t *testing.T) {
	discard := NewDiscard()

	if card, err := discard.TakeTopCard(); err != nil || card != 13 {
		t.Fatalf("TakeTopCard() on an empty pile returned %d, want 13 and an error", card)
	}

	err := discard.AddCard(7)
	if err != nil {
		t.Fatalf("AddCard(7) returned an unexpected error: %v", err)
	}
	err = discard.AddCard(3)
	if err != nil {
		t.Fatalf("AddCard(3) returned an unexpected error: %v", err)
	}
	card, err := discard.TakeTopCard();
	if card != 7 {
		t.Fatalf("first discarded card is %d, want 7", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
	card, err = discard.TakeTopCard();
	if card != 3 {
		t.Fatalf("second discarded card is %d, want 3", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
}
