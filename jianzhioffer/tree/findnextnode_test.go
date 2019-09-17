package tree

import (
	"fmt"
	"testing"
)

func TestFindNextNode(t *testing.T) {
	n0 := &BinaryTreeNode{Value: 0}
	n1 := &BinaryTreeNode{Value: 1}
	n2 := &BinaryTreeNode{Value: 2}
	n3 := &BinaryTreeNode{Value: 3}
	n4 := &BinaryTreeNode{Value: 4}
	n5 := &BinaryTreeNode{Value: 5}
	n6 := &BinaryTreeNode{Value: 6}
	n7 := &BinaryTreeNode{Value: 7}
	n8 := &BinaryTreeNode{Value: 8}
	n9 := &BinaryTreeNode{Value: 9}
	n10 := &BinaryTreeNode{Value: 10}
	n0.Left, n0.Right, n1.Parent, n2.Parent = n1, n2, n0, n0
	n1.Left, n1.Right, n3.Parent, n4.Parent = n3, n4, n1, n1
	n2.Left, n2.Right, n5.Parent, n6.Parent = n5, n6, n2, n2
	n3.Left, n7.Parent = n7, n3
	n4.Left, n4.Right, n8.Parent, n9.Parent = n8, n9, n4, n4
	n6.Right, n10.Parent = n10, n6
	//                 0
	//              /     \
	//             1       2
	//            / \     / \
	//           3   4   5   6
	//          /   / \       \
	//         7   8   9      10
	// 前序遍历结果：0 1 3 7 4 8 9 2 5 6 10
	// 中序遍历结果：7 3 1 8 4 9 0 5 2 6 10
	// 后序遍历结果：7 3 8 9 4 1 5 10 6 2 0
	fmt.Println(FindNextNodeByInOrder(n9))
	fmt.Println(FindNextNodeByPreOrder(n7))
	fmt.Println(FindNextNodeByPostOrder(n4))
}
