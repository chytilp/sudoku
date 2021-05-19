package utils

import (
	"reflect"
)

//ReverseSlice func reverts slice (inline), first item is original last item and so on
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
