package tree

// 判断一颗树是不是BST
func IsBST(root *IntBinaryTreeNode) bool {
	return IsBSTCore(root, 1<<63-1, -1<<63)
}

func IsBSTCore(root *IntBinaryTreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Value > max {
		return false
	}
	if root.Value < min {
		return false
	}
	return IsBSTCore(root.Left, root.Value, min) && IsBSTCore(root.Right, max, root.Value)
}

func Size(root *IntBinaryTreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	count += Size(root.Left)
	count += Size(root.Right)
	return count
}

// 在二叉树中查找最大的BST，返回它的节点数
func FindMaxBST(root *IntBinaryTreeNode) int {
	if root == nil {
		return 0
	}
	if IsBST(root) {
		return Size(root)
	}
	leftBSTCount := FindMaxBST(root.Left)
	rightBSTCount := FindMaxBST(root.Right)
	if leftBSTCount > rightBSTCount {
		return leftBSTCount
	} else {
		return rightBSTCount
	}
}
