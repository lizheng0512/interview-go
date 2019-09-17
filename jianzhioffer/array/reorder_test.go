package array

import (
	"fmt"
	"testing"
)

func TestMakeOddBeforeEvenNumber(t *testing.T) {
	var arr []int
	arr = []int{1, 2, 3, 4, 5, 6, 7}
	f := func(i int) bool {
		return i&0x01 != 0
	}
	Reorder(arr, f)
	fmt.Println(arr)
	arr = []int{-1, -2, -3, -4, -5, -6, -7}
	Reorder(arr, f)
	fmt.Println(arr)
}
