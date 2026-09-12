package main

import (
	"backend/internal/game"
	"fmt"
)

func main() {
	game := game.NewGame(2)
	fmt.Println("Initial grid state for player 0:")
	game.PrintGrid(0)


}
