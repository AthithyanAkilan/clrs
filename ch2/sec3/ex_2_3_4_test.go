package sec3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecursiveInsertionSort(t *testing.T) {
	type args struct {
		arr []int
		res []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: fmt.Sprintf("%v", ([]int{100, 20, 30})),
			args: args{arr: []int{100, 20, 30}, res: []int{20, 30, 100}}},
		{name: fmt.Sprintf("%v", ([]int{100, 20, 30})),
			args: args{arr: []int{-100, 20, -3}, res: []int{-100, -3, 20}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecursiveInsertionSort(tt.args.arr, len(tt.args.arr)-1)
			if !reflect.DeepEqual(tt.args.arr, tt.args.res) {
				t.Errorf("%v is expected but we have %v", tt.args.res, tt.args.arr)
			}
		})
	}
}
