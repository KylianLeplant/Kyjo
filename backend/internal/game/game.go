package game

import (
	"fmt"
	"io"
)

type TurnPhase int

const (
	MinPlayerCount = 2
	MaxPlayerCount = 8
	CardGridWidth  = 4
	CardGridHeight = 3
)

const (
	PhaseReady TurnPhase = iota
	PhaseAwaitingDrawDecision
	PhaseAwaitingReveal
)

func isValidPlayerCount(count int) bool {
	return count >= MinPlayerCount && count <= MaxPlayerCount
}

type pendingTurn struct {
	playerIndex int
	drawnCard   int
	phase       TurnPhase
}

type Game struct {
	deck        *deck
	discard     *discardPile
	grids       []*cardGrid
	pendingTurn *pendingTurn
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

	game.pendingTurn = &pendingTurn{
		playerIndex: 0,
		phase:       PhaseReady,
		drawnCard:   HiddenCardValue,
	}
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

func (g *Game) PrintGrid(w io.Writer, playerIndex int) error {
	grid, err := g.gridForPlayer(playerIndex)
	if err != nil {
		return err
	}
	values, _ := grid.getGridState()
	for j := 0; j < len(values[0]); j++ {
		for i := 0; i < len(values); i++ {
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

func (g *Game) GetCurrentPlayer() int {
	return g.pendingTurn.playerIndex
}

func (g *Game) nextTurn() {
	g.pendingTurn = &pendingTurn{
		playerIndex: (g.pendingTurn.playerIndex + 1) % len(g.grids),
		phase:       PhaseReady,
		drawnCard:   HiddenCardValue,
	}
}

func (g *Game) GetDiscardTopCard() (int, error) {
	return g.discard.getTopCard()
}

func (g *Game) DrawCard(player int) (int, error) {
	if g.pendingTurn == nil {
		return HiddenCardValue, fmt.Errorf("no pending turn")
	}
	if player != g.pendingTurn.playerIndex {
		return HiddenCardValue, fmt.Errorf("it's not player %d's turn", player)
	}
	if g.pendingTurn.phase != PhaseReady {
		return HiddenCardValue, fmt.Errorf("not in the right phase to draw a card")
	}
	card, err := g.deck.drawCard()
	if err != nil {
		return HiddenCardValue, err
	}
	g.pendingTurn = &pendingTurn{
		playerIndex: player,
		drawnCard:   card,
		phase:       PhaseAwaitingDrawDecision,
	}
	return card, nil
}

func (g *Game) TakeDiscardCard(player int) (int, error) {
	if g.pendingTurn == nil {
		return HiddenCardValue, fmt.Errorf("no pending turn")
	}
	if player != g.pendingTurn.playerIndex {
		return HiddenCardValue, fmt.Errorf("it's not player %d's turn", player)
	}
	if g.pendingTurn.phase != PhaseReady {
		return HiddenCardValue, fmt.Errorf("not in the right phase to take a discard card")
	}
	card, err := g.discard.takeTopCard()
	if err != nil {
		return HiddenCardValue, err
	}
	g.pendingTurn = &pendingTurn{
		playerIndex: player,
		drawnCard:   card,
		phase:       PhaseAwaitingDrawDecision,
	}
	return card, nil
}

func (g *Game) DiscardCard(player int) error {
	if g.pendingTurn == nil {
		return fmt.Errorf("Error : no pending turn")
	}
	if player != g.pendingTurn.playerIndex {
		return fmt.Errorf("it's not player %d's turn", player)
	}
	if g.pendingTurn.phase != PhaseAwaitingDrawDecision {
		return fmt.Errorf("Error : not in the right phase to discard")
	}
	err := g.discard.addCard(g.pendingTurn.drawnCard)
	if err != nil {
		return err
	}
	g.pendingTurn.phase = PhaseAwaitingReveal
	return nil
}

func (g *Game) ReplaceCard(player int, x int, y int) error {
	if g.pendingTurn == nil {
		return fmt.Errorf("Error : no pending turn")
	}
	if player != g.pendingTurn.playerIndex {
		return fmt.Errorf("it's not player %d's turn", player)
	}
	if g.pendingTurn.phase != PhaseAwaitingDrawDecision {
		return fmt.Errorf("Error : not in the right phase to replace the card.")
	}
	grid, err := g.gridForPlayer(player)
	if err != nil {
		return err
	}
	lastCard, err := grid.replaceCard(x, y, g.pendingTurn.drawnCard)
	if err != nil {
		return err
	}
	err = g.discard.addCard(lastCard)
	if err != nil {
		return err
	}
	g.nextTurn()
	return nil
}

func (g *Game) RevealCard(player int, x int, y int) (bool, error) {
	if g.pendingTurn == nil {
		return false, fmt.Errorf("Error : no pending turn")
	}
	if player != g.pendingTurn.playerIndex {
		return false, fmt.Errorf("it's not player %d's turn", player)
	}
	if g.pendingTurn.phase != PhaseAwaitingReveal {
		return false, fmt.Errorf("Error : not in the right phase to reveal a card.")
	}
	grid, err := g.gridForPlayer(player)
	if err != nil {
		return false, err
	}
	match, err := grid.discoverCard(x, y)
	if err != nil {
		return false, err
	}
	g.nextTurn()
	return match, nil
}
