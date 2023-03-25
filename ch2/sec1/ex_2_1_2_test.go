package sec1

import (
	"reflect"
	"testing"
)

func TestInsertionSortNonDecreasingOrder(t *testing.T) {
	arr := []int{100, 0, 20, 35, 12}
	InsertionSortNonDecreasingOrder(arr)
	res := []int{100, 35, 20, 12, 0}
	if !reflect.DeepEqual(arr, res) {
		t.Errorf("%v is expected but we have %v", arr, res)
	}
}
