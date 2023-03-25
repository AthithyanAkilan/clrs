package sec2

import (
	"reflect"
	"sort"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{100, 0, 20, 35, 12}
	SelectionSort(arr)
	res := []int{100, 35, 20, 12, 0}
	sort.Ints(res)
	if !reflect.DeepEqual(arr, res) {
		t.Errorf("%v is expected but we have %v", res, arr)
	}
}
