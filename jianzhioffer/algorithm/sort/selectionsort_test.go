package sort

import (
	"fmt"
	"testing"
)

func TestSimpleSort(t *testing.T) {
	arr := []int{3, 2, 5, 1, 4, 6, 13, 2, 5, 1, 23, 7, 14, 0, 11}
	SelectionSort(arr)
	fmt.Println(arr)
}
