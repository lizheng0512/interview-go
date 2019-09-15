package tree

// 题目：给定一颗二叉树和一个节点，找出中序遍历的下一个节点
// 思路：如果这个节点有右子树，则下一个节点为右子树中不包含左节点的节点
// 如果没有右子树，则判断它是否是它父节点的左节点，如果是，则下一个节点为它的父节点
// 如果是右节点且没有右子树，则它的下一个节点是向上遍历遇到的第一个为父节点的左节点的节点的父节点
func FindNextNodeByInOrder(node *binaryTreeNode) *binaryTreeNode {
	if node == nil {
		return nil
	}
	if node.right != nil {
		pr := node.right
		for pr.left != nil {
			pr = pr.left
		}
		return pr
	} else if p := node.parent; p != nil {
		if p.left == node {
			return p
		} else {
			// 往上找，直到有一个节点是它父节点的左节点
			for p.parent != nil && p.parent.left != p {
				p = p.parent
			}
			if p.parent == nil {
				return nil
			} else {
				return p.parent
			}
		}
	}
	return nil
}

func FindNextNodeByPreOrder(node *binaryTreeNode) *binaryTreeNode {
	if node == nil {
		return nil
	}
	if node.left != nil {
		return node.left
	} else if node.right != nil {
		return node.right
	} else if p := node.parent; p != nil {
		// node左右节点都为空时，往上找，直到找到一个节点为父节点的左节点，并且右节点不为空，这时返回右节点
		for p.right == nil || p.right == node {
			node = p
			p = node.parent
			if p == nil {
				return nil
			}
		}
		if p.right != nil {
			return p.right
		}
	}
	return nil
}

func FindNextNodeByPostOrder(node *binaryTreeNode) *binaryTreeNode {
	if node == nil {
		return nil
	}
	if p := node.parent; p != nil {
		pr := p.right
		// 只有当node为左节点，并且右节点不为空的时候，依次按照从左到右往下找，直到找到第一个没有子节点的
		if pr != node && pr != nil {
			for pr.left != nil || pr.right != nil {
				if pr.left != nil {
					pr = pr.left
				} else {
					pr = pr.right
				}
			}
			return pr
		}
		return p
	}
	return nil
}
