package engine

import (
	"errors"
	//"math/rand"

	"github.com/chytilp/sudoku/structures"
)

//Engine struct represent engine for solving sudoku game.
type Engine struct {
	game *structures.Game
}

//PrintStatus dispalys solution cells of the game.
func (e *Engine) PrintStatus() {
	e.game.ShowSolutionCells()
}

//MakeStep creates one step in solving sudoku game.
func (e *Engine) MakeStep() (*bool, error) {
	candidates, err := e.SelectBestCandidates()
	if err != nil {
		return nil, err
	}
	result := false
	for cellID, values := range candidates {
		if len(values) == 1 {
			// simplest case, candidates has only best value.
			cell, _ := structures.NewSolutionCell(cellID, values[0])
			e.game.AddCell(cell)
			result = true
			return &result, nil
		}
		//TODO: not implemented yet - cases with len(values) != 1.
	}
	return &result, nil
}

//NextStepCandidates returns proposals for next step (cells and their values)
func (e *Engine) nextStepCandidates() (map[string][]uint8, error) {
	m := make(map[string][]uint8)
	emptyCells := e.game.EmptyCells()
	for _, cellID := range emptyCells {
		cell, err := structures.NewSolutionCell(cellID, 0)
		if err != nil {
			return nil, err
		}
		values, err := e.game.CellFreeValues(cell)
		if err != nil {
			return nil, err
		}
		m[cellID] = values
	}
	return m, nil
}

//SelectBestCandidates finds out candidates for next step and selects those
// with lowest number of proposed values.
func (e *Engine) SelectBestCandidates() (map[string][]uint8, error) {
	candidates, err := e.nextStepCandidates()
	if err != nil {
		return nil, err
	}
	min := 9
	bestCandidates := make(map[string][]uint8)
	for cellID, vals := range candidates {
		if len(vals) < min {
			min = len(vals)
			bestCandidates = make(map[string][]uint8)
		}
		if len(vals) <= min {
			bestCandidates[cellID] = vals
		}
	}
	return bestCandidates, nil
}

//IsFinished returns if game is finished or not.
func (e *Engine) IsFinished() bool {
	emptyCells := e.game.EmptyCellCount()
	//fmt.Printf("empty cells: %d\n", emptyCells)
	return emptyCells == 0
}

//Run method runs solving sudoku process until it finishes game.
func (e *Engine) Run() (*bool, error) {
	var counter int
	var result bool
	//var stepOk *bool
	for !e.IsFinished() {
		_, err := e.MakeStep()
		if err != nil {
			return nil, err
		}

		//e.PrintStatus()
		counter++
		if counter > 100 {
			break
		}
	}
	if !e.IsFinished() {
		return nil, errors.New("game was not finished in 100 steps")
	}
	return &result, nil
}

//MakePlan analyzes game and returns Plan of solutions (paths).
func (e *Engine) MakePlan() (*Plan, error) {
	var plan Plan
	//TODO: create plan
	return &plan, nil
}
