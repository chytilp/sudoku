package structures

import (
	"fmt"
	"strconv"
	"strings"
)

var convertTable = map[string]uint8{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9}

//Error messages for cell object.
var (
	ErrInvalidCellIDFormatMsg = "Invalid format of cell id: %s. Allowed a1-i9."
	ErrInvalidValueMsg        = "Value should be number 0-9. Given number: %d"
)

//Empty cell representations.
const (
	EmptyCellValue     uint8  = 0
	EmptyCellTextValue string = "."
)

// NewCell creates cell object with solutionCell=false.
func NewCell(id string, value uint8) (*cell, error) {
	return createCell(id, value, false)
}

// NewCellFromString creates cell object from string representation.
func NewCellFromString(text string) (*cell, error) {
	// format a1=5 o resp. a1=5 x
	id := string(text[:2])
	tmp, err := strconv.ParseUint(string(text[3]), 10, 8)
	if err != nil {
		return nil, err
	}
	value := uint8(tmp)
	if string(text[5]) == "x" {
		return NewSolutionCell(id, value)
	} else {
		return NewCell(id, value)
	}
}

// NewSolutionCell creates cell object with solutionCell=true.
func NewSolutionCell(id string, value uint8) (*cell, error) {
	return createCell(id, value, true)
}

func createCell(id string, value uint8, solution bool) (*cell, error) {
	id, err := validateIDFormat(id)
	if err != nil {
		return nil, err
	}
	_, err = validateColumn(id)
	if err != nil {
		return nil, err
	}
	_, err = validateRow(id)
	if err != nil {
		return nil, err
	}
	c := cell{Id: id, solutionCell: solution}
	c.SetValue(value)
	return &c, nil
}

func validateIDFormat(id string) (string, error) {
	id = strings.TrimSpace(id)
	if len(id) != 2 {
		return "", fmt.Errorf(ErrInvalidCellIDFormatMsg, id)
	}
	return id, nil
}

func validateColumn(id string) (uint8, error) {
	idPart := strings.ToLower(string(id[0]))
	column, ok := convertTable[idPart]
	if !ok {
		return 0, fmt.Errorf(ErrInvalidCellIDFormatMsg, id)
	}
	return column, nil
}

func validateRow(id string) (uint8, error) {
	idPart := string(id[1])
	row, err := strconv.ParseUint(idPart, 10, 8)
	if err != nil {
		return 0, fmt.Errorf(ErrInvalidCellIDFormatMsg, id)
	}
	if row < 1 || row > 9 {
		return 0, fmt.Errorf(ErrInvalidCellIDFormatMsg, id)
	}
	return uint8(row), nil
}

type cell struct {
	Id           string
	value        *uint8
	solutionCell bool
}

//Value returns cell value.
func (c *cell) Value() uint8 {
	if c.value == nil {
		return EmptyCellValue
	}
	return *c.value
}

//TextValue returns cell value in text format.
func (c *cell) TextValue() string {
	value := c.Value()
	if value == EmptyCellValue {
		return ""
	}
	return string(value)
}

//SetValue can validate and set cell value.
func (c *cell) SetValue(value uint8) error {
	if value < 0 || value > 9 {
		return fmt.Errorf(ErrInvalidValueMsg, value)
	}
	if c.value == nil {
		var v uint8
		v = value
		c.value = &v
	} else {
		*c.value = value
	}
	return nil
}

//Row returns cell row index.
func (c *cell) Row() uint8 {
	row, _ := validateRow(c.Id)
	return row
}

//TextColumn returns cell column index in text format ex.a,b,c...
func (c *cell) TextColumn() string {
	idPart := strings.ToLower(string(c.Id[0]))
	return idPart
}

//Column returns cell column index.
func (c *cell) Column() uint8 {
	column, _ := validateColumn(c.Id)
	return column
}

//SolutionCell returns if cell is solution cell or not.
func (c *cell) SolutionCell() bool {
	return c.solutionCell
}

//Square returns cell square index.
func (c *cell) Square() uint8 {
	r := c.Row()
	s := c.Column()
	if r >= 1 && r <= 3 {
		if s >= 1 && s <= 3 {
			return 1
		}
		if s >= 4 && s <= 6 {
			return 2
		}
		if s >= 7 && s <= 9 {
			return 3
		}
	} else if r >= 4 && r <= 6 {
		if s >= 1 && s <= 3 {
			return 4
		}
		if s >= 4 && s <= 6 {
			return 5
		}
		if s >= 7 && s <= 9 {
			return 6
		}
	} else {
		if s >= 1 && s <= 3 {
			return 7
		}
		if s >= 4 && s <= 6 {
			return 8
		}
		if s >= 7 && s <= 9 {
			return 9
		}
	}
	return 0
}

//String returns string representation of cell.
func (c *cell) String() string {
	var mark string
	if c.solutionCell {
		mark = "x"
	} else {
		mark = "o"
	}
	return fmt.Sprintf("%s=%d %s", c.Id, c.Value(), mark)
}

//IsEqual returns if cell are equal (same row and column) or not.
func (c *cell) IsEqual(b *cell) bool {
	return c.Id == b.Id
}
