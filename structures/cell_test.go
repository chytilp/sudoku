package structures

import (
	"fmt"
	"testing"
)

func TestCellCorrectCoordinates(t *testing.T) {
	var tests = []struct {
		id                 string
		expectedRow        uint8
		expectedColumn     uint8
		expectedTextColumn string
		expectedSquare     uint8
	}{
		{"a1", 1, 1, "a", 1},
		{"i9", 9, 9, "i", 9},
	}
	for _, test := range tests {
		c, _ := NewCell(test.id, 0)
		if got := c.Row(); got != test.expectedRow {
			t.Errorf("id=%s Cell.Row() = %d, expected %d", test.id, got, test.expectedRow)
		}
		if got := c.Column(); got != test.expectedColumn {
			t.Errorf("id=%s Cell.Column() = %d, expected %d", test.id, got, test.expectedColumn)
		}
		if got := c.TextColumn(); got != test.expectedTextColumn {
			t.Errorf("id=%s Cell.TextColumn() = %s, expected %s", test.id, got, test.expectedTextColumn)
		}
		if got := c.Square(); got != test.expectedSquare {
			t.Errorf("id=%s Cell.Square() = %d, expected %d", test.id, got, test.expectedSquare)
		}
	}
}

func TestCellCreateFromString(t *testing.T) {
	var tests = []struct {
		input                string
		expectedID           string
		expectedValue        uint8
		expectedSolutionCell bool
	}{
		{"a1=5 x", "a1", 5, true},
		{"d7=1 o", "d7", 1, false},
		{"i9=9 x", "i9", 9, true},
	}
	for _, test := range tests {
		c, _ := NewCellFromString(test.input)
		if c.Id != test.expectedID {
			t.Errorf("NewCellFromString creates cell, Id=%s, but expected was Id=%s", c.Id, test.expectedID)
		}
		if c.Value() != test.expectedValue {
			t.Errorf("NewCellFromString creates cell, Value=%d, but expected was Value=%d", c.Value(), test.expectedValue)
		}
		if c.SolutionCell() != test.expectedSolutionCell {
			t.Errorf("NewCellFromString creates cell, solutionCell=%t, but expected was solutionCell=%t", c.SolutionCell(), test.expectedSolutionCell)
		}
	}
}

func TestCellWrongCoordinates(t *testing.T) {
	var tests = []struct {
		id          string
		expectedErr string
	}{
		{"j7", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "j7")},
		{"a0", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "a0")},
		{"78", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "78")},
		{"aa", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "aa")},
		{"a", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "a")},
		{"a10", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "a10")},
		{"k12", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "k12")},
		{"  ", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "")},
		{"", fmt.Sprintf(ErrInvalidCellIDFormatMsg, "")},
	}
	for _, test := range tests {
		_, err := NewCell(test.id, 0)
		if err == nil {
			t.Errorf("NewCell for id: %s should return  error", test.id)
			return
		}
		if err.Error() != test.expectedErr {
			t.Errorf("Expected error is: %v, but was %v", test.expectedErr, err.Error())
		}
	}
}

func TestCellsAreEqual(t *testing.T) {
	c1, _ := NewCell("a1", 0)
	c2, _ := NewCell("a1", 1)
	c3, _ := NewCell("a2", 1)
	if !c1.IsEqual(c2) {
		t.Errorf("Cells c1: %s and c2: %s must be equal.", c1, c2)
	}
	if c2.IsEqual(c3) {
		t.Errorf("Cells c2: %s and c3: %s need not be equal.", c1, c2)
	}
}

func TestCellSetValue(t *testing.T) {
	var expected uint8
	c, _ := NewCell("a1", expected)
	if c.Value() != expected {
		t.Errorf("Cell c: %s expected value: %d, but is %d.", c, expected, c.Value())
	}
	expected = 5
	c.SetValue(expected)
	if c.Value() != expected {
		t.Errorf("Cell c: %s expected value: %d, but is %d.", c, expected, c.Value())
	}
}

func TestCellSetInvalidValue(t *testing.T) {
	c, _ := NewCell("a1", 0)
	var wrongValue uint8 = 10
	expectedErr := fmt.Sprintf(ErrInvalidValueMsg, wrongValue)
	err := c.SetValue(wrongValue)
	if err == nil {
		t.Errorf("Cell.SetValue(%d) should return  error", wrongValue)
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error is: %v, but was %v", expectedErr, err.Error())
	}
}

func TestCellStringRepresentation(t *testing.T) {
	c, _ := NewCell("a1", 0)
	repr := fmt.Sprintf("%s", c)
	expected := "a1=0 o"
	if repr != expected {
		t.Errorf("Representation of Cell should be %s, but was %s.", expected, repr)
	}
	c, _ = NewSolutionCell("b7", 5)
	repr = fmt.Sprintf("%s", c)
	expected = "b7=5 x"
	if repr != expected {
		t.Errorf("Representation of Cell should be %s, but was %s.", expected, repr)
	}
}
