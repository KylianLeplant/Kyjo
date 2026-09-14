package game

import "testing"

func TestDiscardPile(t *testing.T) {
	discard := newDiscardPile()
	card, err := discard.takeTopCard()
	if card != HiddenCardValue {
		t.Fatalf("TakeTopCard() on an empty pile returned %d, want %d", card, HiddenCardValue)
	}
	if err == nil {
		t.Fatalf("TakeTopCard() on an empty pile returned no error")
	}

	//ajout de cartes
	err = discard.addCard(7)
	if err != nil {
		t.Fatalf("AddCard(7) returned an unexpected error: %v", err)
	}
	err = discard.addCard(3)
	if err != nil {
		t.Fatalf("AddCard(3) returned an unexpected error: %v", err)
	}

	//récupération de la carte du dessus
	card, err = discard.takeTopCard()
	if card != 3 {
		t.Fatalf("first discarded card is %d, want 3", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
	card, err = discard.takeTopCard()
	if card != 7 {
		t.Fatalf("second discarded card is %d, want 7", card)
	}
	if err != nil {
		t.Fatalf("TakeTopCard() returned an unexpected error: %v", err)
	}
}

func TestDiscardPileGetTopCard(t *testing.T) {
	discard := newDiscardPile()

	if _, err := discard.getTopCard(); err == nil {
		t.Fatal("GetTopCard() on an empty pile returned no error")
	}
	if err := discard.addCard(7); err != nil {
		t.Fatal(err)
	}
	if err := discard.addCard(3); err != nil {
		t.Fatal(err)
	}

	card, err := discard.getTopCard()
	if err != nil || card != 3 {
		t.Fatalf("GetTopCard() returned (%d, %v), want (3, nil)", card, err)
	}
	if card, err := discard.takeTopCard(); err != nil || card != 3 {
		t.Fatalf("GetTopCard() removed or changed the top card: TakeTopCard() returned (%d, %v)", card, err)
	}
}

func TestDiscardPileTakeAllCards(t *testing.T) {
	discard := newDiscardPile()
	for _, card := range []int{7, 3, 1} {
		if err := discard.addCard(card); err != nil {
			t.Fatal(err)
		}
	}

	cards := discard.takeAllCards()
	if len(cards) != 3 || cards[0] != 7 || cards[1] != 3 || cards[2] != 1 {
		t.Fatalf("TakeAllCards() returned %v, want [7 3 1]", cards)
	}
	if !discard.isEmpty() {
		t.Fatal("discard pile is not empty after TakeAllCards()")
	}
	if cards := discard.takeAllCards(); len(cards) != 0 {
		t.Fatalf("second TakeAllCards() returned %v, want an empty slice", cards)
	}
}

func TestDiscardPileRejectsInvalidCard(t *testing.T) {
	for _, card := range []int{MinCardValue - 1, MaxCardValue + 1} {
		discard := newDiscardPile()

		if err := discard.addCard(card); err == nil {
			t.Fatalf("AddCard(%d) returned no error", card)
		}
		if !discard.isEmpty() {
			t.Fatalf("discard pile contains card %d after a failed AddCard", card)
		}
	}
}
