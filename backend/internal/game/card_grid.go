package game

import (
	"fmt"
)

// cardGrid stores card values and their discovery state by column and row.
type cardGrid struct {
	values     [][]int
	discovered [][]bool
}

// newCardGrid creates a card grid from a flat list of values.
func newCardGrid(values []int, nbCol int, nbRow int) (*cardGrid, error) {
	if nbCol <= 0 || nbRow <= 0 {
		return nil, fmt.Errorf("grid dimensions must be positive")
	}
	if len(values) != nbCol*nbRow {
		return nil, fmt.Errorf("number of values does not match grid dimensions")
	}
	grid := &cardGrid{
		values:     make([][]int, nbCol),
		discovered: make([][]bool, nbCol),
	}
	for i := 0; i < nbCol; i++ {
		grid.values[i] = make([]int, nbRow)
		grid.discovered[i] = make([]bool, nbRow)
		for j := 0; j < nbRow; j++ {
			grid.values[i][j] = values[i*nbRow+j]
			grid.discovered[i][j] = false
		}
	}
	return grid, nil
}

// GetGridState returns visible values and discovery flags as independent copies.
// Hidden cards are represented by the value HiddenCardValue (13).
func (cg *cardGrid) getGridState() ([][]int, [][]bool) {
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
				valuesCopy[i][j] = HiddenCardValue
			}
		}
	}
	return valuesCopy, discoveredCopy
}

// discoverCard reveals a card and removes its column when it forms a match.
func (cg *cardGrid) discoverCard(x int, y int) (bool, error) {
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

// replaceCard changes a card value and marks the card as discovered.
// If the column forms a match after the replacement, it is removed.
func (cg *cardGrid) replaceCard(x int, y int, newValue int) error {
	if !isValidCardValue(newValue) {
		return fmt.Errorf("invalid card value: %d (must be between %d and %d)", newValue, MinCardValue, MaxCardValue)
	}
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

// discoverAll reveals every card in the grid.
func (cg *cardGrid) discoverAll() {
	for i := range cg.values {
		for j := range cg.values[i] {
			cg.discovered[i][j] = true
		}
	}
}

// checkColumnMatch reports whether every card in a column is discovered and equal.
func (cg *cardGrid) checkColumnMatch(colIndex int) bool {
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

// deleteColumn removes a column after a successful match.
func (cg *cardGrid) deleteColumn(colIndex int) {
	cg.values = append(cg.values[:colIndex], cg.values[colIndex+1:]...)
	cg.discovered = append(cg.discovered[:colIndex], cg.discovered[colIndex+1:]...)
}

func (cg *cardGrid) printTestGrid() {
	if len(cg.values) == 0 {
		return
	}
	for i := 0; i < len(cg.values[0]); i++ {
		for j := range cg.values {
			fmt.Printf("%d ", cg.values[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(cg.discovered[0]); i++ {
		for j := range cg.discovered {
			fmt.Printf("%t ", cg.discovered[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
}
