package sec1

func AddBinaryArray(a, b []int) (c []int) {
	c = make([]int, len(a)+1)
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		c[i+1] = (a[i] + b[i] + carry) % 2
		carry = (a[i] + b[i] + carry) / 2
	}
	c[0] = carry
	return c
}
