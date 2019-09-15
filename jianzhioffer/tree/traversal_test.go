package tree

import "testing"

//       0
//     /   \
//    1     2
//   / \   / \
//  3   4 5   6
var tree = &binaryTreeNode{
	value: 0,
	left: &binaryTreeNode{
		value: 1,
		left:  &binaryTreeNode{value: 3},
		right: &binaryTreeNode{value: 4},
	},
	right: &binaryTreeNode{
		value: 2,
		left:  &binaryTreeNode{value: 5},
		right: &binaryTreeNode{value: 6},
	},
}

func TestPreOrderTraversal(t *testing.T) {
	PreOrderTraversal(tree)
}

func TestInOrderTraversal(t *testing.T) {
	InOrderTraversal(tree)
}

func TestPostOrderTraversal(t *testing.T) {
	PostOrderTraversal(tree)
}
