package engine

import (
	//"fmt"
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
	engine := Engine{game: g}
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
	checkGameIsFinished(t, &engine)
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
