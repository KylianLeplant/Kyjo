package game

import (
	"bytes"
	"testing"
)

func TestNewGame(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if len(game.grids) != 2 {
		t.Fatalf("game has %d grids, want 2", len(game.grids))
	}
	if got := game.deck.getNbCards(); got != 125 {
		t.Fatalf("deck contains %d cards after setup, want 125", got)
	}

	for player := range game.grids {
		values, discovered, err := game.GetGridState(player)
		if err != nil {
			t.Fatalf("GetGridState(%d) returned an unexpected error: %v", player, err)
		}
		if len(values) != 4 || len(discovered) != 4 {
			t.Fatalf("player %d grid has an unexpected number of columns", player)
		}
		for column := range values {
			if len(values[column]) != 3 || len(discovered[column]) != 3 {
				t.Fatalf("player %d grid has an unexpected column size", player)
			}
		}
	}
}

func TestGameAccessors(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if game.grids[0] == nil {
		t.Fatal("player 0 grid is nil")
	}
	if game.deck == nil {
		t.Fatal("game deck is nil")
	}
	if game.discard == nil {
		t.Fatal("game discard pile is nil")
	}
}

func TestNewGameRejectsInvalidPlayerCount(t *testing.T) {
	for _, players := range []int{0, 1, 9, -1} {
		game, err := NewGame(players)
		if err == nil {
			t.Fatalf("NewGame(%d) returned no error", players)
		}
		if game != nil {
			t.Fatalf("NewGame(%d) returned a game with an error", players)
		}
	}
}

func TestPrintGridWritesToWriter(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	var output bytes.Buffer
	if err := game.PrintGrid(&output, 0); err != nil {
		t.Fatalf("PrintGrid() returned an unexpected error: %v", err)
	}
	if output.Len() == 0 {
		t.Fatal("PrintGrid() wrote no output")
	}
}

func TestNewGameInitializesTurn(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if game.pendingTurn == nil {
		t.Fatal("NewGame() did not initialize the pending turn")
	}
	if game.GetCurrentPlayer() != 0 {
		t.Fatalf("initial player is %d, want 0", game.GetCurrentPlayer())
	}
	if game.pendingTurn.phase != PhaseReady {
		t.Fatalf("initial phase is %d, want PhaseReady", game.pendingTurn.phase)
	}
	if game.pendingTurn.drawnCard != HiddenCardValue {
		t.Fatalf("initial drawn card is %d, want HiddenCardValue", game.pendingTurn.drawnCard)
	}
}

func TestDrawCardCanOnlyBeCalledOncePerTurn(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := game.DrawCard(1); err == nil {
		t.Fatal("DrawCard() accepted a card draw from the wrong player")
	}
	card, err := game.DrawCard(0)
	if err != nil {
		t.Fatal(err)
	}
	if card < MinCardValue || card > MaxCardValue {
		t.Fatalf("DrawCard() returned invalid card value %d", card)
	}
	if got := game.deck.getNbCards(); got != 124 {
		t.Fatalf("deck contains %d cards after drawing, want 124", got)
	}

	if _, err := game.DrawCard(0); err == nil {
		t.Fatal("DrawCard() allowed a second draw before resolving the first one")
	}
	if got := game.deck.getNbCards(); got != 124 {
		t.Fatalf("deck contains %d cards after rejected second draw, want 124", got)
	}
}

func TestDiscardDrawnCardAndReveal(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	drawnCard, err := game.DrawCard(0)
	if err != nil {
		t.Fatal(err)
	}
	if err := game.DiscardCard(0); err != nil {
		t.Fatal(err)
	}
	if game.pendingTurn.phase != PhaseAwaitingReveal {
		t.Fatalf("phase after discard is %d, want PhaseAwaitingReveal", game.pendingTurn.phase)
	}
	if topCard, err := game.GetDiscardTopCard(); err != nil || topCard != drawnCard {
		t.Fatalf("discard top card is (%d, %v), want (%d, nil)", topCard, err, drawnCard)
	}

	if _, err := game.RevealCard(1, 0, 0); err == nil {
		t.Fatal("RevealCard() accepted the wrong player")
	}
	if _, err := game.RevealCard(0, 0, 0); err != nil {
		t.Fatal(err)
	}
	if game.GetCurrentPlayer() != 1 {
		t.Fatalf("current player after reveal is %d, want 1", game.GetCurrentPlayer())
	}
}

func TestReplaceDrawnCardAdvancesTurn(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	drawnCard, err := game.DrawCard(0)
	if err != nil {
		t.Fatal(err)
	}
	if err := game.ReplaceCard(0, 0, 0); err != nil {
		t.Fatal(err)
	}

	values, discovered, err := game.GetGridState(0)
	if err != nil {
		t.Fatal(err)
	}
	if values[0][0] != drawnCard || !discovered[0][0] {
		t.Fatalf("replaced card state is (%d, %t), want (%d, true)", values[0][0], discovered[0][0], drawnCard)
	}
	if game.GetCurrentPlayer() != 1 {
		t.Fatalf("current player after replacement is %d, want 1", game.GetCurrentPlayer())
	}
}

func TestActionsRejectInvalidPhase(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if err := game.DiscardCard(0); err == nil {
		t.Fatal("DiscardCard() succeeded before drawing a card")
	}
	if err := game.ReplaceCard(0, 0, 0); err == nil {
		t.Fatal("ReplaceCard() succeeded before drawing a card")
	}
	if _, err := game.RevealCard(0, 0, 0); err == nil {
		t.Fatal("RevealCard() succeeded before discarding a drawn card")
	}
}
