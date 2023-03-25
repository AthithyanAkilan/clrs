package sec3

func MergeSortWithoutSentinels(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		mergeWithoutSentinels(arr, p, q, r)
	}

}

func mergeWithoutSentinels(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, 0)
	right := make([]int, 0)
	for i := p; i <= q; i++ {
		left = append(left, arr[i])
	}
	for i := q + 1; i <= r; i++ {
		right = append(right, arr[i])
	}
	i := 0
	j := 0
	k := 0
	for k = p; i < n1 && j < n2; k++ {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	for ; i < n1; i, k = i+1, k+1 {
		arr[k] = left[i]
	}
	for ; j < n2; j, k = j+1, k+1 {
		arr[k] = right[j]
	}
}
