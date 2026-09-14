package game

import "testing"

func TestDiscardPile(t *testing.T) {
	discard := NewDiscardPile()
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

func TestDiscardPileGetTopCard(t *testing.T) {
	discard := NewDiscardPile()

	if _, err := discard.GetTopCard(); err == nil {
		t.Fatal("GetTopCard() on an empty pile returned no error")
	}
	if err := discard.AddCard(7); err != nil {
		t.Fatal(err)
	}
	if err := discard.AddCard(3); err != nil {
		t.Fatal(err)
	}

	card, err := discard.GetTopCard()
	if err != nil || card != 3 {
		t.Fatalf("GetTopCard() returned (%d, %v), want (3, nil)", card, err)
	}
	if card, err := discard.TakeTopCard(); err != nil || card != 3 {
		t.Fatalf("GetTopCard() removed or changed the top card: TakeTopCard() returned (%d, %v)", card, err)
	}
}

func TestDiscardPileTakeAllCards(t *testing.T) {
	discard := NewDiscardPile()
	for _, card := range []int{7, 3, 1} {
		if err := discard.AddCard(card); err != nil {
			t.Fatal(err)
		}
	}

	cards := discard.TakeAllCards()
	if len(cards) != 3 || cards[0] != 7 || cards[1] != 3 || cards[2] != 1 {
		t.Fatalf("TakeAllCards() returned %v, want [7 3 1]", cards)
	}
	if !discard.IsEmpty() {
		t.Fatal("discard pile is not empty after TakeAllCards()")
	}
	if cards := discard.TakeAllCards(); len(cards) != 0 {
		t.Fatalf("second TakeAllCards() returned %v, want an empty slice", cards)
	}
}

func TestDiscardPileRejectsInvalidCard(t *testing.T) {
	for _, card := range []int{MinCardValue - 1, MaxCardValue + 1} {
		discard := NewDiscardPile()

		if err := discard.AddCard(card); err == nil {
			t.Fatalf("AddCard(%d) returned no error", card)
		}
		if !discard.IsEmpty() {
			t.Fatalf("discard pile contains card %d after a failed AddCard", card)
		}
	}
}
