package sec3

import "testing"

func TestIterativeBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Pass1", args: args{arr: []int{20, 21, 23, 100}, key: 100}, want: 3},
		{name: "Pass2", args: args{arr: []int{-221, -10, 12, 123}, key: -221}, want: 0},
		{name: "Fail1", args: args{arr: []int{20, 21, 23, 100}, key: 110}, want: -1},
		{name: "Fail2", args: args{arr: []int{-221, -10, 12, 123}, key: 1}, want: -1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IterativeBinarySearch(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("IterativeBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Pass1", args: args{arr: []int{20, 21, 23, 100}, key: 100}, want: 3},
		{name: "Pass2", args: args{arr: []int{-221, -10, 12, 123}, key: -221}, want: 0},
		{name: "Fail1", args: args{arr: []int{20, 21, 23, 100}, key: 110}, want: -1},
		{name: "Fail2", args: args{arr: []int{-221, -10, 12, 123}, key: 1}, want: -1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecursiveBinarySearch(tt.args.arr, 0, len(tt.args.arr)-1, tt.args.key); got != tt.want {
				t.Errorf("RecursiveBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
