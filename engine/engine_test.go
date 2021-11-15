package engine

import (
	//"fmt"

	"reflect"
	"testing"

	"github.com/chytilp/sudoku/structures"
)

const game1 string = `..5|.1.|6..
    39.|7..|.1.
    4.7|9..|.28
    97.|8..|.4.
    5..|2.6|..1
    .8.|..1|.73
    26.|..8|7.4
    .5.|..7|.39
    ..8|.4.|1..`

func TestEngineShouldFinishGame(t *testing.T) {
	g, err := structures.NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	engine := NewEngine(g)
	var counter int
	var stepOk *bool
	for !engine.IsFinished() {
		stepOk, err = engine.MakeStep()
		if err != nil {
			t.Errorf("Engine.MakeStep should pass, but err: %v", err)
		}
		if !*stepOk {
			t.Error("Engine.MakeStep next step was not found.")
		}
		//engine.PrintStatus()
		counter++
		if counter > 100 {
			break
		}
	}
	checkGameIsFinished(t, engine)
}

func TestEngineShouldFindNextStepCandidates(t *testing.T) {
	g, err := structures.NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	engine := NewEngine(g)
	candidates, err := engine.nextStepCandidates()
	if err != nil {
		t.Errorf("Game should find next step candidates, but err: %v", err)
	}
	expected := 45
	if len(candidates) != expected {
		t.Errorf("Game next step candidates count: %d, but expected: %d", len(candidates),
			expected)
	}
	tests := []struct {
		cellID       string
		expectedVals []uint8
	}{
		{"a1", []uint8{8}},
		{"i1", []uint8{7}},
		{"g6", []uint8{2, 5, 9}},
	}
	for _, test := range tests {
		vals, ok := candidates[test.cellID]
		if !ok {
			t.Errorf("Game candidate for cell: %s was not found.", test.cellID)
		}
		if !reflect.DeepEqual(vals, test.expectedVals) {
			t.Errorf("Game candidate values for cell: %s: %v, but expected: %v", test.cellID,
				vals, test.expectedVals)
		}
	}
	/*for k, v := range candidates {
		fmt.Printf("k=%s,v=%d\n", k, v)
	}*/
}

func TestEngineShouldFindBestCandidates(t *testing.T) {
	g, err := structures.NewGameFromString(game1)
	if err != nil {
		t.Errorf("Game should be succesfully created, but err: %v", err)
	}
	engine := NewEngine(g)
	bestCandidates, err := engine.SelectBestCandidates()
	if err != nil {
		t.Errorf("Game should find best candidates, but err: %v", err)
	}
	expected := map[string][]uint8{
		"i1": []uint8{7},
		"b3": []uint8{1},
		"h1": []uint8{9},
		"h7": []uint8{5},
		"i2": []uint8{5},
		"a1": []uint8{8},
		"a6": []uint8{6},
		"b1": []uint8{2},
		"b9": []uint8{3},
		"a8": []uint8{1},
		"a9": []uint8{7},
	}
	same := reflect.DeepEqual(bestCandidates, expected)
	if !same {
		t.Errorf("Best candidates are %v, but expected was %v.", bestCandidates,
			expected)
	}
}

func checkGameIsFinished(t *testing.T, engine *Engine) {
	rowOk, colOk, squareOk, err := engine.game.Validate()
	if err != nil {
		t.Errorf("Game validation should pass, but returns err: %v.", err)
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
	if !engine.IsFinished() {
		t.Error("Game should be finished.")
	}
	emptyCells := engine.game.EmptyCellCount()
	if emptyCells != 0 {
		t.Errorf("Game have %d empty cells, but expected is 0.", emptyCells)
	}
}
