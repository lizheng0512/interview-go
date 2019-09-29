package algorithm

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	length := int(1e8)
	fmt.Println(length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	fmt.Println(TopK(arr, 1000))
}
