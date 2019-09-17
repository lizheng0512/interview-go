package tree

// 题目：根据二叉树的前序和中序遍历结果，重建二叉树，假设前序和中序遍历结果中都不包含重复的数字
// 思路：前序遍历结果中第一个数字肯定是根节点，再在中序遍历结果中找根节点的位置，得出根节点左侧的节点都是左子树的节点，
// 根节点右侧的节点都是右子树的节点，再根据左子树出现的节点个数来确定前序遍历结果中根节点之后的哪几个左子树节点，
// 剩余的则为右子树节点
func Construct(preOrder, inOrder []int) *BinaryTreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 || len(preOrder) != len(inOrder) {
		panic("请输入正确的前序遍历和中序遍历结果")
	}

	return constructRecursively(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

func constructRecursively(preOrder []int, ps, pe int, inOrder []int, is, ie int) *BinaryTreeNode {
	if ps > pe {
		return nil
	}

	// 获取根节点的值
	rootValue := preOrder[ps]

	// 记录根节点在中序遍历中出现的位置
	index := is
	for ; index <= ie && inOrder[index] != rootValue; index++ {
	}

	// 如果在中序遍历结果中没有找到rootValue
	if index > ie {
		panic("输入的前序遍历和中序遍历结果有误")
	}

	node := &BinaryTreeNode{Value: rootValue}

	// 将左子树的前序遍历和中序遍历的结果序列输入到下一次递归
	// 前序遍历结果序列起止为[ps + 1, ps + index - is]，index - is为左子树节点的数量
	// 中序遍历结果序列起止为[is, index - 1]，因为根节点的左侧都是左子树节点
	node.Left = constructRecursively(preOrder, ps+1, ps+index-is, inOrder, is, index-1)
	// 将右子树的前序遍历和中序遍历的结果序列输入到下一次递归
	// 前序遍历结果序列起止为[ps + index - is + 1, pe], 为排除根节点和左子树之后剩余的部分
	// 中序遍历结果序列起止为[index + 1, ie], 因为根节点的右侧都是右子树节点
	node.Right = constructRecursively(preOrder, ps+index-is+1, pe, inOrder, index+1, ie)

	return node
}
