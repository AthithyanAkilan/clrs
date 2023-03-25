package sec1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddBinaryArray(t *testing.T) {
	testCases := []struct {
		a []int
		b []int
		c []int
	}{
		{[]int{1, 0, 1}, []int{1, 1, 1}, []int{1, 1, 0, 0}},
		{[]int{1, 0, 0}, []int{0, 0, 1}, []int{0, 1, 0, 1}},
		{[]int{1, 0, 1}, []int{1, 0, 1}, []int{1, 0, 1, 0}}}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("A: %v , B : %v, C: %v", tc.a, tc.b, tc.c), func(t *testing.T) {
			res := AddBinaryArray(tc.a, tc.b)
			if !reflect.DeepEqual(res, tc.c) {
				t.Errorf("Got C=%v ", res)
			}
		})
	}
}
