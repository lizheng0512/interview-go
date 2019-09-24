package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//arr := []int{3, 2, 5, 1, 4, 6, 13, 2, 5, 1, 23, 7, 14, 0, 11}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	arr := []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort(arr)
	fmt.Println(arr)
}

func TestQuickSort2(t *testing.T) {
	arr := []int{3, 2, 5, 1, 4, 6, 13, 2, 5, 1, 23, 7, 14, 0, 11}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//arr := []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	QuickSort2(arr)
	fmt.Println(arr)
}
