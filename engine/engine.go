package engine

import (

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
		//tmp := make([]string, len(values))
		//for i, val := range values {
		//	tmp[i] = fmt.Sprintf("%d", val)
		//}
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

	}
	return &result, nil
}

//MakePlan analyzes game and returns Plan of solutions (paths).
func (e *Engine) MakePlan() (*Plan, error) {
	var plan Plan
	return &plan, nil
}
