package structures

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//Messages for game errors.
const (
	ErrParseCellsNoThreePartsMsg string = "No 3 parts devided by | in row: %s."
	ErrParseCellsNoThreeCharsMsg string = "Row part has not 3 characters: %s."
	ErrDuplicatedCellInGameMsg   string = "Cell id=%s is already presented in game."
	ErrCellWasNotFoundMsg        string = "Cell id=%s was not found in game."
)

//Errors for game object.
var (
	ErrParseCellsNoNineRows                      error = errors.New("No 9 rows in game input")
	ErrOnlyOneFromRowColumnSquareMustBeSpecified error = errors.New("Only one argument from (row, column, square) should be > 0")
	ErrOnlyOneArgumentShouldBeTrue               error = errors.New("Only one argument from (row, column, square) should be true")
)

func valueFoundInSlice(slice []uint8, value uint8) bool {
	index := sort.Search(len(slice), func(i int) bool { return slice[i] >= value })
	found := index < len(slice) && slice[index] == value
	return found
}

func parseCells(textCells string) ([]*cell, error) {
	rows := strings.Split(textCells, "\n")
	if len(rows) != 9 {
		return nil, ErrParseCellsNoNineRows
	}
	var err error
	var cells, tmp []*cell
	for index, row := range rows {
		tmp, err = parseRow(row, uint8(index+1))
		if err != nil {
			return nil, err
		}
		cells = append(cells, tmp...)
	}
	return cells, nil
}

func parseRow(textRow string, rowIndex uint8) ([]*cell, error) {
	parts := strings.Split(strings.TrimSpace(textRow), "|")
	if len(parts) != 3 {
		return nil, fmt.Errorf(ErrParseCellsNoThreePartsMsg, textRow)
	}
	var cells, tmp []*cell
	var err error
	cells, err = processRowPart(parts[0], rowIndex, "abc")
	if err != nil {
		return nil, err
	}
	tmp, err = processRowPart(parts[1], rowIndex, "def")
	if err != nil {
		return nil, err
	}
	cells = append(cells, tmp...)
	tmp, err = processRowPart(parts[2], rowIndex, "ghi")
	if err != nil {
		return nil, err
	}
	cells = append(cells, tmp...)
	return cells, nil
}

func processRowPart(part string, rowIndex uint8, columns string) ([]*cell, error) {
	if len(part) != 3 {
		return nil, fmt.Errorf(ErrParseCellsNoThreeCharsMsg, part)
	}
	var cells []*cell
	var c *cell
	var value uint64
	var err error
	for index, char := range part {
		if string(char) != EmptyCellTextValue {
			value, err = strconv.ParseUint(string(char), 10, 8)
			if err != nil {
				return nil, err
			}
			c, err = NewCell(string(columns[index])+fmt.Sprintf("%d", rowIndex), uint8(value))
			if err != nil {
				return nil, err
			}
			cells = append(cells, c)
		}
	}
	return cells, nil
}

//NewGameFromString creates Game object from string representation.
func NewGameFromString(textCells string) (*Game, error) {
	cells, err := parseCells(textCells)
	if err != nil {
		return nil, err
	}
	game, err := NewGameFromCells(cells)
	if err != nil {
		return nil, err
	}
	return game, nil
}

//NewGameFromCells creates Game object from slice of cells.
func NewGameFromCells(cells []*cell) (*Game, error) {
	g := Game{}
	g.cells = make(map[string]*cell)
	var err error
	for _, c := range cells {
		err = g.AddCell(c)
		if err != nil {
			return nil, err
		}
	}
	return &g, nil
}

//Game struct represnts one sudoku game.
type Game struct {
	cells         map[string]*cell
	solutionSteps []string
}

//AddCell method add new cell to the game.
func (g *Game) AddCell(c *cell) error {
	_, ok := g.cells[c.Id]
	if ok {
		return fmt.Errorf(ErrDuplicatedCellInGameMsg, c.Id)
	}
	g.cells[c.Id] = c
	if c.SolutionCell() {
		g.solutionSteps = append(g.solutionSteps, c.Id)
	}
	return nil
}

//AddStringCell method add new cell (from string representation) to the game.
func (g *Game) AddStringCell(textCell string) error {
	c, err := NewCellFromString(textCell)
	if err != nil {
		return err
	}
	err = g.AddCell(c)
	if err != nil {
		return err
	}
	return nil
}

//Cell returns cell from game by id in parameter.
func (g *Game) Cell(id string) (*cell, error) {
	c, ok := g.cells[id]
	if !ok {
		return nil, fmt.Errorf(ErrCellWasNotFoundMsg, id)
	}
	return c, nil
}

//EmptyCellCount returns count of empty cells in the game.
func (g *Game) EmptyCellCount() uint8 {
	return uint8(81 - g.FilledCellCount())
}

//FilledCellCount returns count of filled cells in the game.
func (g *Game) FilledCellCount() uint8 {
	return uint8(len(g.cells))
}

//Validate method makes validation of the game and returns
// if problem is in rows, columns or squares.
func (g *Game) Validate() (*bool, *bool, *bool, error) {
	var err error
	invalidRows, err := g.validateRows()
	if err != nil {
		return nil, nil, nil, err
	}
	invalidColumns, err := g.validateColumns()
	if err != nil {
		return nil, nil, nil, err
	}
	invalidSquares, err := g.validateSquares()
	if err != nil {
		return nil, nil, nil, err
	}
	rowsOk := len(invalidRows) == 0
	columnsOk := len(invalidColumns) == 0
	squaresOk := len(invalidSquares) == 0
	return &rowsOk, &columnsOk, &squaresOk, nil
}

func (g *Game) validateEntities(rows bool, columns bool, squares bool) ([]uint8, error) {
	if (rows && columns) || (rows && squares) || (columns && squares) {
		return nil, ErrOnlyOneArgumentShouldBeTrue
	}
	var invalid []uint8
	entities := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var partResult *bool
	var err error
	var rowIdx, colIdx, squareIdx uint8
	for _, entity := range entities {
		rowIdx, colIdx, squareIdx = 0, 0, 0
		if rows {
			rowIdx = entity
		} else if columns {
			colIdx = entity
		} else if squares {
			squareIdx = entity
		}
		partResult, err = g.specifiedCellsHaveDuplicities(rowIdx, colIdx, squareIdx)
		if err != nil {
			return nil, err
		}
		if *partResult {
			invalid = append(invalid, entity)
		}
	}
	return invalid, nil
}

func (g *Game) validateRows() ([]uint8, error) {
	invalid, err := g.validateEntities(true, false, false)
	return invalid, err
}

func (g *Game) validateColumns() ([]uint8, error) {
	invalid, err := g.validateEntities(false, true, false)
	return invalid, err
}

func (g *Game) validateSquares() ([]uint8, error) {
	invalid, err := g.validateEntities(false, false, true)
	return invalid, err
}

func (g *Game) specifiedCellsHaveDuplicities(row uint8, column uint8, square uint8) (*bool, error) {
	if (row > 0 && column > 0) || (row > 0 && square > 0) || (column > 0 && square > 0) {
		return nil, ErrOnlyOneFromRowColumnSquareMustBeSpecified
	}
	valueOccurences := make(map[uint8]uint8)
	result := false
	values, err := g.specifiedCellsValues(row, column, square)
	if err != nil {
		return nil, err
	}
	for _, val := range values {
		if valueOccurences[val] > 0 {
			result = true
			return &result, nil
		}
		valueOccurences[val]++
	}
	return &result, nil
}

func (g *Game) specifiedCellsValues(row uint8, column uint8, square uint8) ([]uint8, error) {
	if (row > 0 && column > 0) || (row > 0 && square > 0) || (column > 0 && square > 0) {
		return nil, ErrOnlyOneFromRowColumnSquareMustBeSpecified
	}
	var result []uint8
	for _, c := range g.cells {
		if (row > 0 && c.Row() == row) || (column > 0 && c.Column() == column) ||
			(square > 0 && c.Square() == square) {
			result = append(result, c.Value())
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result, nil
}

//CellFreeValues returns for the cell which values can have yet.
func (g *Game) CellFreeValues(cell *cell) ([]uint8, error) {
	rowValues, err := g.specifiedCellsValues(cell.Row(), 0, 0)
	if err != nil {
		return nil, err
	}
	colValues, err := g.specifiedCellsValues(0, cell.Column(), 0)
	if err != nil {
		return nil, err
	}
	squareValues, err := g.specifiedCellsValues(0, 0, cell.Square())
	if err != nil {
		return nil, err
	}
	var result []uint8
	var rFound, cFound, sFound bool
	for v := 1; v < 10; v++ {
		rFound = valueFoundInSlice(rowValues, uint8(v))
		cFound = valueFoundInSlice(colValues, uint8(v))
		sFound = valueFoundInSlice(squareValues, uint8(v))
		if !rFound && !cFound && !sFound {
			result = append(result, uint8(v))
		}
	}
	return result, nil
}

func (g *Game) findCellValue(rowIdx uint8, colIdx uint8) string {
	for _, c := range g.cells {
		if c.Row() == rowIdx && c.Column() == colIdx {
			return c.TextValue()
		}
	}
	return EmptyCellTextValue
}

//GameVisual returns visual representation of the game.
func (g *Game) GameVisual() string {
	var visual string
	for r := 1; r < 10; r++ {
		line := ""
		for c := 1; c < 10; c++ {
			line += g.findCellValue(uint8(r), uint8(c))
		}
		visual += line[:3] + "|" + line[3:6] + "|" + line[6:] + "\n"
	}
	return visual
}

//EmptyCells returns slice of empty cells in the game.
func (g *Game) EmptyCells() []string {
	columns := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	var ok bool
	var result []string
	var id string
	for r := 1; r < 10; r++ {
		for _, c := range columns {
			id = string(c) + fmt.Sprintf("%d", r)
			_, ok = g.cells[id]
			if !ok {
				result = append(result, id)
			}
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

//ShowSolutionCells prints all solution cells of the game.
func (g *Game) ShowSolutionCells() {
	for _, c := range g.cells {
		if c.SolutionCell() {
			fmt.Printf("%s\n", c)
		}
	}
}
