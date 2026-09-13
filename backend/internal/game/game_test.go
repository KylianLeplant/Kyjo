package game

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame(2)

	if len(game.grids) != 2 {
		t.Fatalf("game has %d grids, want 2", len(game.grids))
	}
	if got := game.deck.getNbCards(); got != 126 {
		t.Fatalf("deck contains %d cards after setup, want 126", got)
	}

	for player := range game.grids {
		values, discovered := game.GetGridState(player)
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
	game := NewGame(2)

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
