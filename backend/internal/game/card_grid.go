package game

import (
	"fmt"
)

// CardGrid stores card values and their discovery state by column and row.
type CardGrid struct {
	values     [][]int
	discovered [][]bool
	nbCol      int
	nbRow      int
}

// NewCardGrid creates a card grid from a flat list of values.
func NewCardGrid(values []int, nbCol int, nbRow int) *CardGrid {
	grid := &CardGrid{
		values:     make([][]int, nbCol),
		discovered: make([][]bool, nbCol),
		nbCol:      nbCol,
		nbRow:      nbRow,
	}
	for i := 0; i < nbCol; i++ {
		grid.values[i] = make([]int, nbRow)
		grid.discovered[i] = make([]bool, nbRow)
		for j := 0; j < nbRow; j++ {
			grid.values[i][j] = values[i*nbRow+j]
			grid.discovered[i][j] = false
		}
	}
	return grid
}

// GetGridState returns visible values and discovery flags as independent copies.
// Hidden cards are represented by the value 13.
func (cg *CardGrid) GetGridState() ([][]int, [][]bool) {
	valuesCopy := make([][]int, len(cg.values))
	discoveredCopy := make([][]bool, len(cg.values))

	for i := range cg.values {
		valuesCopy[i] = make([]int, len(cg.values[i]))
		discoveredCopy[i] = make([]bool, len(cg.values[i]))
		for j := range cg.values[i] {
			discoveredCopy[i][j] = cg.discovered[i][j]
			if cg.discovered[i][j] {
				valuesCopy[i][j] = cg.values[i][j]
			} else {
				valuesCopy[i][j] = 13 // 13 indicates a hidden card
			}
		}
	}
	return valuesCopy, discoveredCopy
}

// Discover reveals a card and removes its column when it forms a match.
func (cg *CardGrid) Discover(x int, y int) (bool, error) {
	if x < 0 || x >= len(cg.values) || y < 0 || y >= len(cg.values[x]) {
		return false, fmt.Errorf("invalid card position (%d, %d)", x, y)
	}
	if !cg.discovered[x][y] {
		cg.discovered[x][y] = true
		if cg.checkColumnMatch(x) {
			cg.deleteColumn(x)
		}
		return true, nil
	}
	return false, nil
}

// ReplaceCard changes a card value and marks the card as discovered.
// If the column forms a match after the replacement, it is removed.
func (cg *CardGrid) ReplaceCard(x int, y int, newValue int) error {
	if x < 0 || x >= len(cg.values) || y < 0 || y >= len(cg.values[x]) {
		return fmt.Errorf("invalid card position (%d, %d)", x, y)
	}
	cg.values[x][y] = newValue
	cg.discovered[x][y] = true
	if cg.checkColumnMatch(x) {
		cg.deleteColumn(x)
	}
	return nil
}

// DiscoverAll reveals every card in the grid.
func (cg *CardGrid) DiscoverAll() {
	for i := 0; i < cg.nbCol; i++ {
		for j := 0; j < cg.nbRow; j++ {
			cg.discovered[i][j] = true
		}
	}
}

// checkColumnMatch reports whether every card in a column is discovered and equal.
func (cg *CardGrid) checkColumnMatch(colIndex int) bool {
	for y := 0; y < len(cg.values[colIndex]); y++ {
		if !cg.discovered[colIndex][y] {
			return false
		}
		if y > 0 && cg.values[colIndex][y] != cg.values[colIndex][y-1] {
			return false
		}
	}
	return true
}

// deleteColumn duplicates the column values after a successful match.
func (cg *CardGrid) deleteColumn(colIndex int) {
	cg.nbCol--
	cg.values = append(cg.values[:colIndex], cg.values[colIndex+1:]...)
	cg.discovered = append(cg.discovered[:colIndex], cg.discovered[colIndex+1:]...)
}

func (cg *CardGrid) printTestGrid() {
	for i := 0; i < cg.nbRow; i++ {
		for j := 0; j < cg.nbCol; j++ {
			fmt.Printf("%d ", cg.values[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < cg.nbRow; i++ {
		for j := 0; j < cg.nbCol; j++ {
			fmt.Printf("%t ", cg.discovered[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
}
