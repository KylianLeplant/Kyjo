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
	deck          *Deck
	discard       *DiscardPile
	grids         []*CardGrid
	currentPlayer int
}

func NewGame(nbPlayers int) (*Game, error) {
	if !isValidPlayerCount(nbPlayers) {
		return nil, fmt.Errorf("invalid number of players: %d (must be between 2 and 8)", nbPlayers)
	}
	game := &Game{
		deck:    NewDeck(),
		discard: NewDiscardPile(),
		grids:   make([]*CardGrid, nbPlayers),
	}
	for i := 0; i < nbPlayers; i++ {
		cards, err := game.deck.DrawCards(CardGridWidth * CardGridHeight)
		if err != nil {
			return nil, err
		}
		game.grids[i], err = NewCardGrid(cards, CardGridWidth, CardGridHeight)
		if err != nil {
			return nil, err
		}
	}

	card, err := game.deck.DrawCard()
	if err != nil {
		return nil, err
	}

	err = game.discard.AddCard(card)
	if err != nil {
		return nil, err
	}

	game.currentPlayer = 0
	return game, nil
}

func (g *Game) GetGridState(playerIndex int) ([][]int, [][]bool, error) {
	if playerIndex < 0 || playerIndex >= len(g.grids) {
		return nil, nil, fmt.Errorf("invalid player index: %d", playerIndex)
	}
	values, discovered := g.grids[playerIndex].GetGridState()
	return values, discovered, nil
}

func (g *Game) PrintGrid(w io.Writer, playerIndex int) error {
	if playerIndex < 0 || playerIndex >= len(g.grids) {
		return fmt.Errorf("invalid player index: %d", playerIndex)
	}
	values, _ := g.grids[playerIndex].GetGridState()
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
