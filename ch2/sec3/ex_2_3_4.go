package sec3

func RecursiveInsertionSort(arr []int, curr int) {
	if curr == 0 {
		return
	}
	RecursiveInsertionSort(arr, curr-1)
	key := arr[curr]
	j := 0
	for j = curr - 1; j >= 0 && arr[j] > key; j-- {
		arr[j+1] = arr[j]
	}
	arr[j+1] = key
}
