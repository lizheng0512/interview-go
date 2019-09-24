package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(BinarySearchRecursively(arr, 10))
	fmt.Println(BinarySearchIteratively(arr, 0))
	fmt.Println(BinarySearchIteratively(arr, 1))
	fmt.Println(BinarySearchIteratively(arr, 2))
	fmt.Println(BinarySearchIteratively(arr, 3))
	fmt.Println(BinarySearchIteratively(arr, 4))
	fmt.Println(BinarySearchIteratively(arr, 5))
	fmt.Println(BinarySearchIteratively(arr, 6))
	fmt.Println(BinarySearchIteratively(arr, 7))
	fmt.Println(BinarySearchIteratively(arr, 8))
	fmt.Println(BinarySearchIteratively(arr, 9))
	fmt.Println(BinarySearchIteratively(arr, 10))
	fmt.Println(BinarySearchIteratively(arr, 11))
}
