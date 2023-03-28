package sec3

func IterativeBinarySearch(arr []int, key int) (index int) {
	l := 0
	r := len(arr) - 1
	index = -1
	for l <= r {
		m := (l + r) / 2
		if arr[m] == key {
			index = m
			break
		} else if arr[m] > key {
			r = m - 1
			continue
		}
		l = m + 1
	}
	return index
}

func RecursiveBinarySearch(arr []int, l, r, key int) int {
	if l > r {
		return -1
	}
	m := (l + r) / 2
	if arr[m] == key {
		return m
	}
	if arr[m] > key {
		return RecursiveBinarySearch(arr, l, m-1, key)
	}
	return RecursiveBinarySearch(arr, m+1, r, key)
}
