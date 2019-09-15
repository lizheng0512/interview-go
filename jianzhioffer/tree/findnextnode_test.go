package tree

import (
	"fmt"
	"testing"
)

func TestFindNextNode(t *testing.T) {
	n0 := &binaryTreeNode{value: 0}
	n1 := &binaryTreeNode{value: 1}
	n2 := &binaryTreeNode{value: 2}
	n3 := &binaryTreeNode{value: 3}
	n4 := &binaryTreeNode{value: 4}
	n5 := &binaryTreeNode{value: 5}
	n6 := &binaryTreeNode{value: 6}
	n7 := &binaryTreeNode{value: 7}
	n8 := &binaryTreeNode{value: 8}
	n9 := &binaryTreeNode{value: 9}
	n10 := &binaryTreeNode{value: 10}
	n0.left, n0.right, n1.parent, n2.parent = n1, n2, n0, n0
	n1.left, n1.right, n3.parent, n4.parent = n3, n4, n1, n1
	n2.left, n2.right, n5.parent, n6.parent = n5, n6, n2, n2
	n3.left, n7.parent = n7, n3
	n4.left, n4.right, n8.parent, n9.parent = n8, n9, n4, n4
	n6.right, n10.parent = n10, n6
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
	fmt.Println(FindNextNodeByPostOrder(n3))
}
