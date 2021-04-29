package engine

import (
	"fmt"
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
	emptyCells := e.game.EmptyCells()
	//index := rand.Intn(len(emptyCells))
	result := false
	for _, cellID := range emptyCells {
		cell, err := structures.NewSolutionCell(cellID, 0)
		if err != nil {
			return nil, err
		}
		values, err := e.game.CellFreeValues(cell)
		if err != nil {
			return nil, err
		}
		tmp := make([]string, len(values))
		for i, val := range values {
			tmp[i] = fmt.Sprintf("%d", val)
		}
		//sValues := strings.Join(tmp, ",")
		//fmt.Printf("cell: id=%s, values=%s\n", cellId, sValues)
		if len(values) == 1 {
			cell.SetValue(values[0])
			e.game.AddCell(cell)
			result = true
			return &result, nil
		}
	}
	return &result, nil
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

	}
	return &result, nil
}

//MakePlan analyzes game and returns Plan of solutions (paths).
func (e *Engine) MakePlan() (*Plan, error) {
	var plan Plan
	return &plan, nil
}
