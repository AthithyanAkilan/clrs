package sec3

import "math"

func MergeSort(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		merge(arr, p, q, r)
	}

}

func merge(arr []int, p, q, r int) {

	left := make([]int, 0)
	right := make([]int, 0)
	for i := p; i <= q; i++ {
		left = append(left, arr[i])
	}
	left = append(left, math.MaxInt)
	for i := q + 1; i <= r; i++ {
		right = append(right, arr[i])
	}
	right = append(right, math.MaxInt)
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}
