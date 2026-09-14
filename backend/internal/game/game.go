package game

import (
	"fmt"
	"io"
)

const (
	MinPlayerCount = 2
	MaxPlayerCount = 8
	CardGridWidth  = 4
	CardGridHeight = 3
)

func isValidPlayerCount(count int) bool {
	return count >= MinPlayerCount && count <= MaxPlayerCount
}

type Game struct {
	deck          *deck
	discard       *discardPile
	grids         []*cardGrid
	currentPlayer int
}

func NewGame(nbPlayers int) (*Game, error) {
	if !isValidPlayerCount(nbPlayers) {
		return nil, fmt.Errorf("invalid number of players: %d (must be between 2 and 8)", nbPlayers)
	}
	game := &Game{
		deck:    newDeck(),
		discard: newDiscardPile(),
		grids:   make([]*cardGrid, nbPlayers),
	}
	for i := 0; i < nbPlayers; i++ {
		cards, err := game.deck.drawCards(CardGridWidth * CardGridHeight)
		if err != nil {
			return nil, err
		}
		game.grids[i], err = newCardGrid(cards, CardGridWidth, CardGridHeight)
		if err != nil {
			return nil, err
		}
	}

	card, err := game.deck.drawCard()
	if err != nil {
		return nil, err
	}

	err = game.discard.addCard(card)
	if err != nil {
		return nil, err
	}

	game.currentPlayer = 0
	return game, nil
}

func (g *Game) GetGridState(playerIndex int) ([][]int, [][]bool, error) {
	grid, err := g.gridForPlayer(playerIndex)
	if err != nil {
		return nil, nil, err
	}
	values, discovered := grid.getGridState()
	return values, discovered, nil
}

// Discover reveals a card for the specified player.
func (g *Game) Discover(playerIndex, x, y int) (bool, error) {
	if playerIndex != g.currentPlayer {
		return false, fmt.Errorf("it's not player %d's turn", playerIndex)
	}
	grid, err := g.gridForPlayer(playerIndex)
	if err != nil {
		return false, err
	}
	return grid.discoverCard(x, y)
}

// ReplaceCard replaces a card for the specified player.
func (g *Game) ReplaceCard(playerIndex, x, y, newValue int) error {
	if playerIndex != g.currentPlayer {
		return fmt.Errorf("it's not player %d's turn", playerIndex)
	}
	grid, err := g.gridForPlayer(playerIndex)
	if err != nil {
		return err
	}
	return grid.replaceCard(x, y, newValue)
}

func (g *Game) PrintGrid(w io.Writer, playerIndex int) error {
	grid, err := g.gridForPlayer(playerIndex)
	if err != nil {
		return err
	}
	values, _ := grid.getGridState()
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if _, err := fmt.Fprint(w, values[i][j], " "); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) gridForPlayer(playerIndex int) (*cardGrid, error) {
	if playerIndex < 0 || playerIndex >= len(g.grids) {
		return nil, fmt.Errorf("invalid player index: %d", playerIndex)
	}
	return g.grids[playerIndex], nil
}
