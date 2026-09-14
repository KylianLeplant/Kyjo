package main

import (
	"backend/internal/game"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	currentGame, err := game.NewGame(2)
	if err != nil {
		fmt.Println("Unable to create game:", err)
		return
	}

	turns := 0
	maxTurns := 10

	for turns < maxTurns {
		player := currentGame.GetCurrentPlayer()

		fmt.Printf("\n=== Player %d's turn ===\n", player)

		topCard, err := currentGame.GetDiscardTopCard()
		if err != nil {
			fmt.Println("Discard error:", err)
			return
		}
		fmt.Printf("Discard top card: %d\n", topCard)

		fmt.Printf("Player %d grid:\n", player)
		if err := currentGame.PrintGrid(os.Stdout, player); err != nil {
			fmt.Println("Print grid error:", err)
			return
		}

		source := prompt(scanner, "Draw from (d)eck or take from (t)op of discard? ")

		var drawnCard int
		switch strings.ToLower(source) {
		case "d", "deck":
			fmt.Println("Press Enter to draw a card...")
			scanner.Scan()
			card, err := currentGame.DrawCard(player)
			if err != nil {
				fmt.Println("Draw error:", err)
				return
			}
			drawnCard = card
		case "t", "top", "discard":
			card, err := currentGame.TakeDiscardCard(player)
			if err != nil {
				fmt.Println("Take discard error:", err)
				return
			}
			drawnCard = card
		default:
			fmt.Println("Invalid choice, turn skipped.")
			turns++
			continue
		}
		fmt.Printf("You drew: %d\n", drawnCard)

		choice := prompt(scanner, "Choose: (p)lace or (d)iscard? ")
		switch strings.ToLower(choice) {
		case "p", "place":
			x := readInt(scanner, "Column (x): ")
			y := readInt(scanner, "Row (y): ")
			if err := currentGame.ReplaceCard(player, x, y); err != nil {
				fmt.Println("Replace error:", err)
				return
			}
			fmt.Println("Card placed.")
		case "d", "discard":
			if err := currentGame.DiscardCard(player); err != nil {
				fmt.Println("Discard error:", err)
				return
			}
			fmt.Println("Card discarded. You must reveal a card.")
			x := readInt(scanner, "Column (x): ")
			y := readInt(scanner, "Row (y): ")
			if _, err := currentGame.RevealCard(player, x, y); err != nil {
				fmt.Println("Reveal error:", err)
				return
			}
			fmt.Println("Card revealed.")
		default:
			fmt.Println("Invalid choice, turn skipped.")
		}

		turns++
	}

	fmt.Println("\nGame finished after", maxTurns, "turns.")
}

func prompt(scanner *bufio.Scanner, message string) string {
	fmt.Print(message)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readInt(scanner *bufio.Scanner, message string) int {
	for {
		fmt.Print(message)
		if !scanner.Scan() {
			return 0
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Println("Invalid number, try again.")
	}
}
