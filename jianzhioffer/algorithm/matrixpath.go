package algorithm

func MatrixPathCount(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return MatrixPathCount(m, n-1) + MatrixPathCount(m-1, n)
}
