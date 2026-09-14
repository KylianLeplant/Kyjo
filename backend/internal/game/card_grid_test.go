package game

import "testing"

func TestCardGridState(t *testing.T) {
	grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	if err != nil {
		t.Fatal(err)
	}

	values, discovered := grid.GetGridState()
	if values[0][0] != HiddenCardValue || discovered[0][0] {
		t.Fatal("new cards should be hidden")
	}

	revealed, err := grid.Discover(0, 1)
	if err != nil || !revealed {
		t.Fatal("Discover() should reveal a hidden card")
	}
	revealed, err = grid.Discover(0, 1)
	if err != nil || revealed {
		t.Fatal("Discover() should not reveal an already discovered card")
	}

	values, discovered = grid.GetGridState()
	if values[0][1] != 2 || !discovered[0][1] {
		t.Fatalf("discovered card state is (%d, %t), want (2, true)", values[0][1], discovered[0][1])
	}
}

func TestCardGridStateReturnsCopies(t *testing.T) {
	grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	values, discovered := grid.GetGridState()
	values[0][0] = 99
	discovered[0][0] = true

	values, discovered = grid.GetGridState()
	if values[0][0] != HiddenCardValue || discovered[0][0] {
		t.Fatal("GetGridState() returned references to the grid state")
	}
}

func TestCardGridReplaceCard(t *testing.T) {
	grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if err := grid.ReplaceCard(1, 0, 9); err != nil {
		t.Fatal(err)
	}
	values, discovered := grid.GetGridState()
	if values[1][0] != 9 || !discovered[1][0] {
		t.Fatalf("replaced card state is (%d, %t), want (9, true)", values[1][0], discovered[1][0])
	}
}

func TestCardGridDiscoverRejectsInvalidIndices(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
	}{
		{name: "negative column", x: -1, y: 0},
		{name: "column too large", x: 2, y: 0},
		{name: "negative row", x: 0, y: -1},
		{name: "row too large", x: 0, y: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := grid.Discover(test.x, test.y); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestCardGridReplaceCardRejectsInvalidIndices(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
	}{
		{name: "negative column", x: -1, y: 0},
		{name: "column too large", x: 2, y: 0},
		{name: "negative row", x: 0, y: -1},
		{name: "row too large", x: 0, y: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
			if err != nil {
				t.Fatal(err)
			}
			if err := grid.ReplaceCard(test.x, test.y, 9); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestCardGridReplaceCardRejectsInvalidValue(t *testing.T) {
	for _, value := range []int{MinCardValue - 1, MaxCardValue + 1} {
		grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
		if err != nil {
			t.Fatal(err)
		}

		if err := grid.ReplaceCard(0, 0, value); err == nil {
			t.Fatalf("ReplaceCard() accepted invalid value %d", value)
		}
		values, discovered := grid.GetGridState()
		if values[0][0] != HiddenCardValue || discovered[0][0] {
			t.Fatalf("invalid replacement changed card state to (%d, %t)", values[0][0], discovered[0][0])
		}
	}
}

func TestCardGridDiscoverAll(t *testing.T) {
	grid, err := NewCardGrid([]int{1, 2, 3, 4}, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
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
	grid, err := NewCardGrid([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2)
	if err != nil {
		t.Fatal(err)
	}
	if err := grid.ReplaceCard(0, 1, 1); err != nil {
		t.Fatal(err)
	}
	if grid.nbCol == 3 {
		t.Fatalf("deletion not expected if all cards are not discovered")
	}
	if _, err := grid.Discover(0, 0); err != nil {
		t.Fatal(err)
	}
	if grid.nbCol != 3 {
		t.Fatalf("deletion expected after discovering the last card of a column if they are the same %d", grid.nbCol)
	}

	if _, err := grid.Discover(0, 0); err != nil {
		t.Fatal(err)
	}
	if err := grid.ReplaceCard(0, 1, 3); err != nil {
		t.Fatal(err)
	}
	if grid.nbCol != 2 {
		t.Fatalf("deletion expected after replacing the last card of a column if they are the same")
	}

	if _, err := grid.Discover(0, 0); err != nil {
		t.Fatal(err)
	}
	if err := grid.ReplaceCard(0, 1, 3); err != nil {
		t.Fatal(err)
	}
	if grid.nbCol == 1 {
		t.Fatalf("deletion not expected when we replace the last card of a column if all cards of the column are not the same")
	}

	if err := grid.ReplaceCard(1, 1, 3); err != nil {
		t.Fatal(err)
	}
	if _, err := grid.Discover(1, 0); err != nil {
		t.Fatal(err)
	}
	if grid.nbCol == 1 {
		t.Fatalf("deletion not expected when we discover the last card of a column if all cards of the column are not the same")
	}
}
