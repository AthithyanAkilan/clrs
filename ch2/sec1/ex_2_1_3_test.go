package sec1

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		arr []int
		v   int
		i   int
		b   bool
	}{
		{[]int{1, 100, -1, 20, 32}, 100, 1, true},
		{[]int{1, 100, -1, 20, 32}, 21, -1, false},
		{[]int{1, 100, -1, 20, 32}, 32, 4, true}}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Array: %v , v : %v, index: %v, found:%v", tc.arr, tc.v, tc.i, tc.b), func(t *testing.T) {
			i, b := LinearSearch(tc.arr, tc.v)
			if tc.i != i || tc.b != b {
				t.Errorf("Got v=%v b=%v", i, b)
			}
		})
	}
}
