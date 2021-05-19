package utils

import (
	"reflect"
	"testing"
)

func TestSliceShouldBeReverted(t *testing.T) {
	origSlice := []string{"p", "r", "z", "a", "b"}
	expected := []string{"b", "a", "z", "r", "p"}
	ReverseSlice(origSlice)
	if !reflect.DeepEqual(origSlice, expected) {
		t.Errorf("ReverseSlice returns %v, but expected was: %v", origSlice, expected)
	}
}
