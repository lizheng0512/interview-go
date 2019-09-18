package tree

import "testing"

//       0
//     /   \
//    1     2
//   / \   / \
//  3   4 5   6
var tree = &BinaryTreeNode{
	Value: 0,
	Left: &BinaryTreeNode{
		Value: 1,
		Left:  &BinaryTreeNode{Value: 3},
		Right: &BinaryTreeNode{Value: 4},
	},
	Right: &BinaryTreeNode{
		Value: 2,
		Left:  &BinaryTreeNode{Value: 5},
		Right: &BinaryTreeNode{Value: 6},
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

func TestWidthFirstTraversal(t *testing.T) {
	WidthFirstTraversal(tree)
	WidthFirstTraversal(nil)
	WidthFirstTraversal(&BinaryTreeNode{Value: 1})
}

func TestWidthFirstTraversal2(t *testing.T) {
	WidthFirstTraversal2(tree)
}
