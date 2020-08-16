package fib

var resultMap = map[int]int{}

func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	nMinusOne := n - 1
	n1, ok := resultMap[nMinusOne]
	if !ok {
		n1 = Fib(nMinusOne)
		resultMap[nMinusOne] = n1
	}
	nMinusTwo := n - 2
	n2, ok := resultMap[nMinusTwo]
	if !ok {
		n2 = Fib(nMinusTwo)
		resultMap[nMinusTwo] = n2
	}

	return n1 + n2
}
