package structures

import (
	"reflect"
	"testing"
)

const game1 string = `8..|94.|..5
    ...|.5.|2..
    1.9|6.2|...
    5.1|...|..4
    46.|...|.53
    2..|...|8.1
    ...|4.9|1.7
    ..4|.6.|...
    9..|.17|..6`

func TestGameObjectCorrectlyCreated(t *testing.T) {
	g, err := NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	var expected uint8 = 30
	if g.FilledCellCount() != expected {
		t.Errorf("Game have filled cells: %d, but expected number: %d", g.FilledCellCount(), expected)
	}
	expected = 51
	if g.EmptyCellCount() != expected {
		t.Errorf("Game has empty cells: %d, but expected number: %d", g.EmptyCellCount(), expected)
	}
	if g.GameVisual() == game1 {
		t.Errorf("Game looks like\n %s, but expected is\n %s", g.GameVisual(), game1)
	}
}

func TestGameCreateEmptyObject(t *testing.T) {
	var cells []*cell
	g, err := NewGameFromCells(cells)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	var expected uint8 = 0
	if g.FilledCellCount() != expected {
		t.Errorf("Game have filled cells: %d, but expected number: %d", g.FilledCellCount(), expected)
	}
	expected = 81
	if g.EmptyCellCount() != expected {
		t.Errorf("Game has empty cells: %d, but expected number: %d", g.EmptyCellCount(), expected)
	}
}

func TestGameObjectHappyValidation(t *testing.T) {
	g, err := NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	rowOk, colOk, squareOk, err := g.Validate()
	if err != nil {
		t.Errorf("Game validation should pass, but err: %v", err)
	}
	if !*rowOk {
		t.Errorf("Game row validation is: %t, but expected: %t", *rowOk, true)
	}
	if !*colOk {
		t.Errorf("Game column validation is: %t, but expected: %t", *colOk, true)
	}
	if !*squareOk {
		t.Errorf("Game square validation is: %t, but expected: %t", *squareOk, true)
	}
}

func TestGameObjectValidationWithErrors(t *testing.T) {
	tests := []struct {
		solutionCell string
		rowOk        bool
		colOk        bool
		squareOk     bool
		solutionStep string
	}{
		{"f1=8 x", false, true, true, "f1"},
		{"b2=6 x", true, false, true, "b2"},
		{"b2=8 x", true, true, false, "b2"},
	}

	var g *Game
	var err error
	for _, test := range tests {
		g, err = NewGameFromString(game1)
		if err != nil {
			t.Errorf("Game should be succesfully created, but err: %v", err)
		}
		err = g.AddStringCell(test.solutionCell)
		if err != nil {
			t.Errorf("Solution cell: %s should be added, but err: %v", test.solutionCell, err)
		}
		rowOk, colOk, squareOk, err := g.Validate()
		if err != nil {
			t.Errorf("Game validation should pass, but err: %v", err)
		}
		if *rowOk != test.rowOk {
			t.Errorf("Game row validation is: %t, but expected: %t", *rowOk, test.rowOk)
		}
		if *colOk != test.colOk {
			t.Errorf("Game column validation is: %t, but expected: %t", *colOk, test.colOk)
		}
		if *squareOk != test.squareOk {
			t.Errorf("Game square validation is: %t, but expected: %t", *squareOk, test.squareOk)
		}
		if g.solutionSteps[0] != test.solutionStep {
			t.Errorf("Game solution step is: %s, but expected: %s", g.solutionSteps[0], test.solutionStep)
		}
	}
}

func TestGameCellFreeValues(t *testing.T) {
	g, err := NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	id := "b1"
	c, err := NewSolutionCell(id, 0)
	if err != nil {
		t.Errorf("Cell id=%s should be created, but err: %v", id, err)
	}
	freeValues, err := g.CellFreeValues(c)
	if err != nil {
		t.Errorf("Cell id=%s free values should be retrurned, but err: %v", id, err)
	}
	expected := []uint8{2, 3, 7}
	if !reflect.DeepEqual(freeValues, expected) {
		t.Errorf("Cell id=%s free values are %v, but expected was: %v", id, freeValues, expected)
	}
}

func TestGameEmptyCells(t *testing.T) {
	g, err := NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	cells := g.EmptyCells()
	expectedCount := 51
	if len(cells) != expectedCount {
		t.Errorf("Game empty cells should be: %d", expectedCount)
	}
	expected := []string{"a2", "a7", "a8", "b1", "b2", "b3", "b4", "b6", "b7", "b8", "b9", "c1", "c2", "c5", "c6", "c7", "c9",
		"d2", "d4", "d5", "d6", "d8", "d9", "e3", "e4", "e5", "e6", "e7", "f1", "f2", "f4", "f5", "f6", "f8",
		"g1", "g3", "g4", "g5", "g8", "g9", "h1", "h2", "h3", "h4", "h6", "h7", "h8", "h9", "i2", "i3", "i8"}
	if !reflect.DeepEqual(cells, expected) {
		t.Errorf("Game empty cells are %v, but expected was: %v", cells, expected)
	}
}
