package sec1

import (
	"reflect"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{100, 0, 20, 35, 12}
	InsertionSort(arr)
	res := []int{100, 35, 20, 12, 0}
	sort.Ints(res)
	if !reflect.DeepEqual(arr, res) {
		t.Errorf("%v is expected but we have %v", arr, res)
	}
}
