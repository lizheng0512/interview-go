package tree

// 求一颗二叉树的直径
//var max = 0
//func DiameterOfBinaryTree(root *IntBinaryTreeNode) int {
//	if root == nil {
//		return 0
//	}
//	height(root)
//	return max
//}
//
//func height(root *IntBinaryTreeNode) int {
//	if root == nil {
//		return 0
//	}
//	hl := height(root.Left)
//	hr := height(root.Right)
//	if hl + hr > max {
//		max = hl + hr
//	}
//	if hl > hr {
//		return hl + 1
//	} else {
//		return hr + 1
//	}
//}

func DiameterOfBinaryTree(root *IntBinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return height(root.Left) + height(root.Right)
}

func height(root *IntBinaryTreeNode) int {
	if root == nil {
		return 0
	}
	hl := height(root.Left)
	hr := height(root.Right)
	if hl > hr {
		return hl + 1
	} else {
		return hr + 1
	}
}
