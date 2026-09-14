package game

import "testing"

func TestNewGame(t *testing.T) {
	game, err := NewGame(2)
	if err != nil {
		t.Fatal(err)
	}

	if len(game.grids) != 2 {
		t.Fatalf("game has %d grids, want 2", len(game.grids))
	}
	if got := game.deck.getNbCards(); got != 126 {
		t.Fatalf("deck contains %d cards after setup, want 126", got)
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
