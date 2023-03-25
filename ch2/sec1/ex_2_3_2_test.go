package sec1

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestMergeSortWithoutSentinels(t *testing.T) {
	testCases := []struct {
		arr []int
		res []int
	}{
		{[]int{100, 0, 20, 35, 1, 12}, []int{100, 35, 20, 12, 1, 0}},
		{[]int{100, 0, 120, -12, 35}, []int{100, 35, 120, -12, 0}},
		{[]int{-100, 0, -20, -35, 12}, []int{-100, -35, -20, 12, 0}},
		{[]int{100, 0, 20, 35, 12}, []int{100, 35, 20, 12, 0}}}
	for _, tc := range testCases {
		sort.Ints(tc.res)
		t.Run(fmt.Sprintf("arr: %v, res: %v", tc.arr, tc.res), func(t *testing.T) {
			MergeSortWithoutSentinels(tc.arr, 0, len(tc.arr)-1)
			if !reflect.DeepEqual(tc.arr, tc.res) {
				t.Errorf("%v is expected but we have %v", tc.res, tc.arr)
			}
		})
	}
}
