package algorithm

// 求斐波那契的第n项

// 效率极低的写法，有大量重复计算
func FibonacciRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// 从前往后计算，将之前的结果保留再变量中
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	num1 := 0
	num2 := 1
	result := 0
	for i := 1; i < n; i++ {
		result = num1 + num2
		num1 = num2
		num2 = result
	}
	return result
}
