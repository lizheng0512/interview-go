package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	//arr := []int{3, 2, 5, 1, 4, 6, 13, 2, 5, 1, 23, 7, 14, 0, 11, 19}
	arr := []int{3, 2, 5, 1, 4, 6, 13, 2, 5, 1, 23, 7, 14, 0, 11}
	MergeSort(arr)
	fmt.Println(arr)
}
