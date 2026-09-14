package game

import "testing"

func TestDiscardPile(t *testing.T) {
	discard := NewDiscard()
	card, err := discard.TakeTopCard()
	if card != HiddenCardValue {
		t.Fatalf("TakeTopCard() on an empty pile returned %d, want %d", card, HiddenCardValue)
	}
	if err == nil {
		t.Fatalf("TakeTopCard() on an empty pile returned no error")
	}

	//ajout de cartes
	err = discard.AddCard(7)
	if err != nil {
		t.Fatalf("AddCard(7) returned an unexpected error: %v", err)
	}
	err = discard.AddCard(3)
	if err != nil {
		t.Fatalf("AddCard(3) returned an unexpected error: %v", err)
	}

	//récupération de la carte du dessus
	card, err = discard.TakeTopCard()
	if card != 3 {
		t.Fatalf("first discarded card is %d, want 3", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
	card, err = discard.TakeTopCard()
	if card != 7 {
		t.Fatalf("second discarded card is %d, want 7", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
}
