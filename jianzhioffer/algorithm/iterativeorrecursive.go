package algorithm

func AddFrom1ToNRecursive(n, r int) int {
	if n <= 0 {
		return r
	} else {
		return AddFrom1ToNRecursive(n-1, n+r)
	}
}

func AddFrom1ToNIterative(n int) (result int) {
	for i := 1; i <= n; i++ {
		result += i
	}
	return
}
