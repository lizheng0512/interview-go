package tree

type IntBinaryTreeNode struct {
	Value       int
	Left, Right *IntBinaryTreeNode
}

// 题目28: 判断一颗树是不是对称二叉树
// 原二叉树和它的镜像是一致的则为对称二叉树
// 判断一个二叉树是不是对称二叉树, 可以判断前序遍历和前序对称遍历(根->右->左)序列是否相等
func IsSymmetrical(root *IntBinaryTreeNode) (result bool) {
	return isSymmetrical(root, root)
}

func isSymmetrical(root, root2 *IntBinaryTreeNode) bool {
	if root == nil && root2 == nil {
		return true
	}
	if root == nil || root2 == nil {
		return false
	}
	if root.Value != root2.Value {
		return false
	}
	return isSymmetrical(root.Left, root2.Right) && isSymmetrical(root.Right, root2.Left)
}
