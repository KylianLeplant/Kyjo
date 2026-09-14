package main

import (
	"fmt"
	"os"

	"backend/internal/game"
)

func main() {
	currentGame, err := game.NewGame(2)
	if err != nil {
		fmt.Println("Unable to create game:", err)
		return
	}
	fmt.Println("Initial grid state for player 0:")
	if err := currentGame.PrintGrid(os.Stdout, 0); err != nil {
		fmt.Println("Unable to print grid:", err)
	}
}
