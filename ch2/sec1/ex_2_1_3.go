package sec1

func LinearSearch(arr []int, s int) (p int, ok bool) {
	for i, v := range arr {
		if v == s {
			return i, true
		}
	}
	return -1, false
}
