package tree

// 题目：给定一颗二叉树和一个节点，找出中序遍历的下一个节点
// 思路：如果这个节点有右子树，则下一个节点为右子树中不包含左节点的节点
// 如果没有右子树，则判断它是否是它父节点的左节点，如果是，则下一个节点为它的父节点
// 如果是右节点且没有右子树，则它的下一个节点是向上遍历遇到的第一个为父节点的左节点的节点的父节点
func FindNextNodeByInOrder(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		pr := node.Right
		for pr.Left != nil {
			pr = pr.Left
		}
		return pr
	} else if p := node.Parent; p != nil {
		if p.Left == node {
			return p
		} else {
			// 往上找，直到有一个节点是它父节点的左节点
			for p.Parent != nil && p.Parent.Left != p {
				p = p.Parent
			}
			if p.Parent == nil {
				return nil
			} else {
				return p.Parent
			}
		}
	}
	return nil
}

func FindNextNodeByPreOrder(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return node.Left
	} else if node.Right != nil {
		return node.Right
	} else if p := node.Parent; p != nil {
		// node左右节点都为空时，往上找，直到找到一个节点为父节点的左节点，并且右节点不为空，这时返回右节点
		for p.Right == nil || p.Right == node {
			node = p
			p = node.Parent
			if p == nil {
				return nil
			}
		}
		if p.Right != nil {
			return p.Right
		}
	}
	return nil
}

func FindNextNodeByPostOrder(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if p := node.Parent; p != nil {
		pr := p.Right
		// 只有当node为左节点，并且右节点不为空的时候，依次按照从左到右往下找，直到找到第一个没有子节点的
		if pr != node && pr != nil {
			for pr.Left != nil || pr.Right != nil {
				if pr.Left != nil {
					pr = pr.Left
				} else {
					pr = pr.Right
				}
			}
			return pr
		}
		// 否则的话直接返回父节点
		return p
	}
	return nil
}
