package tree

import (
	"fmt"
	"testing"
)

func TestIsSubTree(t *testing.T) {
	//           2
	//          /
	//         3
	//        / \
	//       4   5
	//          / \
	//         6   8
	//        /   / \
	//       7   9  10
	//                \
	//                11
	tree1 := &Float64BinaryTreeNode{
		Value: 2,
		Left: &Float64BinaryTreeNode{
			Value: 3,
			Left: &Float64BinaryTreeNode{
				Value: 4,
			},
			Right: &Float64BinaryTreeNode{
				Value: 5,
				Left: &Float64BinaryTreeNode{
					Value: 6,
					Left: &Float64BinaryTreeNode{
						Value: 7,
					},
				},
				Right: &Float64BinaryTreeNode{
					Value: 8,
					Left: &Float64BinaryTreeNode{
						Value: 9,
					},
					Right: &Float64BinaryTreeNode{
						Value: 10,
						Right: &Float64BinaryTreeNode{
							Value: 11,
						},
					},
				},
			},
		},
		Right: nil,
	}
	tree2 := &Float64BinaryTreeNode{
		Value: 8,
		Left: &Float64BinaryTreeNode{
			Value: 9,
		},
		Right: &Float64BinaryTreeNode{
			Value: 10,
			//Right: &Float64BinaryTreeNode{
			//	Value: 11,
			//},
		},
	}
	fmt.Println(IsSubTree(tree1, tree2))
}
