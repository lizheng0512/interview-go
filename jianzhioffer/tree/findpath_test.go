package tree

import (
	"fmt"
	"testing"
)

func TestFindPath(t *testing.T) {
	//       0
	//     /   \
	//    4     2
	//   / \   / \
	//  3   4 5   6
	tree := &IntBinaryTreeNode{
		Value: 0,
		Left: &IntBinaryTreeNode{
			Value: 4,
			Left:  &IntBinaryTreeNode{Value: 3},
			Right: &IntBinaryTreeNode{Value: 4},
		},
		Right: &IntBinaryTreeNode{
			Value: 2,
			Left:  &IntBinaryTreeNode{Value: 5},
			Right: &IntBinaryTreeNode{Value: 6},
		},
	}
	FindPath(tree, 7)
	fmt.Println()
	FindPath(tree, 8)
	FindPath(tree, 0)
	FindPath(nil, 0)
}
