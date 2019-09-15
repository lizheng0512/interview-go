package tree

import "fmt"

type binaryTreeNode struct {
	value               interface{}
	parent, left, right *binaryTreeNode
}

func (n *binaryTreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v", n.value)
}

func PreOrderTraversal(root *binaryTreeNode) {
	if root != nil {
		fmt.Printf("%s ", root)
		PreOrderTraversal(root.left)
		PreOrderTraversal(root.right)
	}
}

func InOrderTraversal(root *binaryTreeNode) {
	if root != nil {
		InOrderTraversal(root.left)
		fmt.Printf("%s ", root)
		InOrderTraversal(root.right)
	}
}

func PostOrderTraversal(root *binaryTreeNode) {
	if root != nil {
		PostOrderTraversal(root.left)
		PostOrderTraversal(root.right)
		fmt.Printf("%s ", root)
	}
}
