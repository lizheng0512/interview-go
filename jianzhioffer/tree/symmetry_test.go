package tree

import (
	"fmt"
	"testing"
)

func TestIsSymmetrical(t *testing.T) {
	tree := &IntBinaryTreeNode{
		Value: 8,
		Left: &IntBinaryTreeNode{
			Value: 9,
		},
		Right: &IntBinaryTreeNode{
			Value: 9,
			//Right: &IntBinaryTreeNode{
			//	Value: 11,
			//},
		},
	}
	fmt.Println(IsSymmetrical(tree))
}
