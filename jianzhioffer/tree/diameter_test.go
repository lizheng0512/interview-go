package tree

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	intTree := getIntTree()
	fmt.Println(DiameterOfBinaryTree(intTree))
	fmt.Println(DiameterOfBinaryTree(&IntBinaryTreeNode{Value: 1}))
}
