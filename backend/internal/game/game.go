package game

type Game struct {
	deck    *Deck
	discard *DiscardPile
	grids   []*CardGrid
}

func NewGame(nbPlayers int) *Game {
	game := &Game{
		deck:    NewDeck(),
		discard: NewDiscard(),
		grids:   make([]*CardGrid, nbPlayers),
	}
	for i := 0; i < nbPlayers; i++ {
		game.grids[i] = NewCardGrid(game.deck.DrawCards(12), 4, 3)
	}
	return game
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
