package tree

import "fmt"

type BinaryTreeNode struct {
	Value               interface{}
	Parent, Left, Right *BinaryTreeNode
}

func (n *BinaryTreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v", n.Value)
}

func PreOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		fmt.Printf("%s ", root)
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	}
}

func InOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Printf("%s ", root)
		InOrderTraversal(root.Right)
	}
}

func PostOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		PostOrderTraversal(root.Left)
		PostOrderTraversal(root.Right)
		fmt.Printf("%s ", root)
	}
}
