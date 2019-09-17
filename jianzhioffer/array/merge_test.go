package array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	//arr1 := []int{1, 3, 4, 5, 7, 9}
	arr2 := []int{2}
	//mergedArr := Merge(arr1, arr2)
	mergedArr := Merge(nil, arr2)
	fmt.Println(mergedArr)
}
