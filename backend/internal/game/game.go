package game

import "fmt"

type Game struct {
	deck    *Deck
	discard *DiscardPile
	grids   []*CardGrid
}

func NewGame(nbPlayers int) (*Game, error) {
	if nbPlayers < 2 || nbPlayers > 8 {
		return nil, fmt.Errorf("invalid number of players: %d (must be between 2 and 8)", nbPlayers)
	}
	game := &Game{
		deck:    NewDeck(),
		discard: NewDiscard(),
		grids:   make([]*CardGrid, nbPlayers),
	}
	for i := 0; i < nbPlayers; i++ {
		cards, err := game.deck.DrawCards(12)
		if err != nil {
			return nil, err
		}
		game.grids[i], err = NewCardGrid(cards, 4, 3)
		if err != nil {
			return nil, err
		}
	}
	return game, nil
}

func (g *Game) GetGridState(playerIndex int) ([][]int, [][]bool) {
	return g.grids[playerIndex].GetGridState()
}

func (g *Game) PrintGrid(playerIndex int) {
	values, _ := g.grids[playerIndex].GetGridState()
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			print(values[i][j], " ")
		}
		println()
	}
}
