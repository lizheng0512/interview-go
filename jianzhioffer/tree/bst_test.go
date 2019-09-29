package tree

import (
	"fmt"
	"testing"
)

func getIntTree() *IntBinaryTreeNode {
	tree := &IntBinaryTreeNode{
		Value: 5,
		Left: &IntBinaryTreeNode{
			Value: 3,
			Left: &IntBinaryTreeNode{
				Value: 2,
			},
			Right: &IntBinaryTreeNode{
				Value: 6,
			},
		},
		Right: &IntBinaryTreeNode{
			Value: 9,
			Left: &IntBinaryTreeNode{
				Value: 7,
				Left:  &IntBinaryTreeNode{Value: 6},
				Right: &IntBinaryTreeNode{Value: 8},
			},
			Right: &IntBinaryTreeNode{
				Value: 10,
			},
		},
	}
	return tree
}

func TestIsBST(t *testing.T) {
	tree := getIntTree()
	fmt.Println(IsBST(tree))
	fmt.Println(Size(tree))
	fmt.Println(FindMaxBST(tree))
}
