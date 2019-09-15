package tree

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {
	preOrder := []int{0, 1, 3, 4, 2, 5, 6}
	inOrder := []int{3, 1, 4, 0, 5, 2, 6}
	root := Construct(preOrder, inOrder)
	PreOrderTraversal(root)
	fmt.Println()
	InOrderTraversal(root)
	fmt.Println()
	PostOrderTraversal(root)
}
