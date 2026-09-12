package game

import "testing"

func TestCardGridState(t *testing.T) {
	grid := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)

	values, discovered := grid.GetGridState()
	if values[0][0] != 13 || discovered[0][0] {
		t.Fatal("new cards should be hidden")
	}

	if !grid.Discover(0, 1) {
		t.Fatal("Discover() should reveal a hidden card")
	}
	if grid.Discover(0, 1) {
		t.Fatal("Discover() should not reveal an already discovered card")
	}

	values, discovered = grid.GetGridState()
	if values[0][1] != 2 || !discovered[0][1] {
		t.Fatalf("discovered card state is (%d, %t), want (2, true)", values[0][1], discovered[0][1])
	}
}

func TestCardGridStateReturnsCopies(t *testing.T) {
	grid := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	values, discovered := grid.GetGridState()
	values[0][0] = 99
	discovered[0][0] = true

	values, discovered = grid.GetGridState()
	if values[0][0] != 13 || discovered[0][0] {
		t.Fatal("GetGridState() returned references to the grid state")
	}
}

func TestCardGridReplaceCard(t *testing.T) {
	grid := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)

	grid.ReplaceCard(1, 0, 9)
	values, discovered := grid.GetGridState()
	if values[1][0] != 9 || !discovered[1][0] {
		t.Fatalf("replaced card state is (%d, %t), want (9, true)", values[1][0], discovered[1][0])
	}
}

func TestCardGridDiscoverAll(t *testing.T) {
	grid := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	grid.DiscoverAll()

	values, discovered := grid.GetGridState()
	for x := range values {
		for y := range values[x] {
			if !discovered[x][y] {
				t.Fatalf("card (%d, %d) is not discovered", x, y)
			}
		}
	}
}

func TestCardGridDeleteColumn(t *testing.T) {
	grid := NewCardGrid([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2)
	grid.ReplaceCard(0, 1, 1)
	if grid.nbCol == 3 {
		t.Fatalf("deletion not expected if all cards are not discovered")
	}
	grid.Discover(0, 0)
	if grid.nbCol != 3 {
		t.Fatalf("deletion expected after discovering the last card of a column if they are the same %d", grid.nbCol)
	}

	grid.Discover(0, 0)
	grid.ReplaceCard(0, 1, 3)
	if grid.nbCol != 2 {
		t.Fatalf("deletion expected after replacing the last card of a column if they are the same")
	}

	grid.Discover(0, 0)
	grid.ReplaceCard(0, 1, 3)
	if grid.nbCol == 1 {
		t.Fatalf("deletion not expected when we replace the last card of a column if all cards of the column are not the same")
	}

	grid.ReplaceCard(1, 1, 3)
	grid.Discover(1, 0)
	if grid.nbCol == 1 {
		t.Fatalf("deletion not expected when we discover the last card of a column if all cards of the column are not the same")
	}
}
