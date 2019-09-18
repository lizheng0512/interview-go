package tree

import (
	"fmt"
	"testing"
)

func TestVerifySequenceOfBST(t *testing.T) {
	fmt.Println(VerifySequenceOfBST([]int{5, 7, 6, 9, 11, 10, 8}))
	fmt.Println(VerifySequenceOfBST([]int{7, 4, 6, 5}))
	//fmt.Println(VerifySequenceOfBST(nil))
	fmt.Println(VerifySequenceOfBST([]int{1, 2, 3}))
}
