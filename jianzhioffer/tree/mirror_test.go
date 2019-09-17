package tree

import (
	"testing"
)

func TestMirrorRecursively(t *testing.T) {
	tree := &BinaryTreeNode{
		Value: 8,
		Left: &BinaryTreeNode{
			Value: 9,
		},
		Right: &BinaryTreeNode{
			Value: 10,
			Right: &BinaryTreeNode{
				Value: 11,
			},
		},
	}
	PreOrderTraversal(tree)
	MirrorRecursively(tree)
	PreOrderTraversal(tree)
}
