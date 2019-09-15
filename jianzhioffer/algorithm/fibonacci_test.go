package algorithm

import (
	"fmt"
	"testing"
)

func TestFibonacciRecursive(t *testing.T) {
	fmt.Println(FibonacciRecursive(100))
}

func TestFibonacci(t *testing.T) {
	fmt.Println(Fibonacci(30))
}
