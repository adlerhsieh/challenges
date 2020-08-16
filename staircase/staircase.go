package staircase

func St(n int, ways []int) int {
	if n == 0 || n == 1 {
		return 1
	}

	sum := 0
	for _, w := range ways {
		newN := n - w
		if newN >= 0 {
			sum += St(n-w, ways)
		}
	}

	return sum
}
