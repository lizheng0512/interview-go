package tree

type Float64BinaryTreeNode struct {
	Value       float64
	Left, Right *Float64BinaryTreeNode
}

// 题目26: 判断一个树tree1中是否包含另一个子树tree2
func IsSubTree(treeRoot1, treeRoot2 *Float64BinaryTreeNode) (result bool) {
	if treeRoot1 == nil || treeRoot2 == nil {
		return
	}
	if treeRoot1.Equal(treeRoot2) {
		result = treeRoot1.HasSubTree(treeRoot2)
	}
	if !result {
		result = IsSubTree(treeRoot1.Left, treeRoot2)
	}
	if !result {
		result = IsSubTree(treeRoot1.Right, treeRoot2)
	}
	return
}

// 判断当前子树和树node2的结构是否相等
func (node *Float64BinaryTreeNode) HasSubTree(node2 *Float64BinaryTreeNode) (result bool) {
	// 走到了叶子节点或者子树走到了叶子节点都返回true
	if (node == nil && node2 == nil) || node2 == nil {
		return true
	}
	if node == nil || !node.Equal(node2) {
		return
	}
	return node.Left.HasSubTree(node2.Left) && node.Right.HasSubTree(node2.Right)
}

// 判断两个节点的值是否相等
func (node *Float64BinaryTreeNode) Equal(node2 *Float64BinaryTreeNode) (result bool) {
	if node == nil || node2 == nil {
		return
	}
	// 判断两个浮点数差的绝对值小于0.00000001, 则认为相等
	if node.Value-node2.Value > -0.00000001 && node.Value-node2.Value < 0.00000001 {
		result = true
	}
	return
}
