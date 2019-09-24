package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//arr := []int{3,9,11,2,4,7,3,5,6}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(arr)
	fmt.Println(arr)
}
