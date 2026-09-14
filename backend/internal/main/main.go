package main

import (
	"backend/internal/game"
	"fmt"
)

func main() {
	currentGame, err := game.NewGame(2)
	if err != nil {
		fmt.Println("Unable to create game:", err)
		return
	}
	fmt.Println("Initial grid state for player 0:")
	currentGame.PrintGrid(0)

}
