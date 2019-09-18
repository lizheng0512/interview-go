package tree

// 题目33: 验证一个后序遍历序列是否是二叉搜索树的
// 思路: 因为是后序遍历, 所以序列的最后一个节点一定是树的根节点, 根据二叉搜索树的定义, 根节点左侧的左子树上左右的节点都小于根节点,
// 右子树上的所有节点都大于根节点, 可以区分出所有的左子树节点和右子树节点
func VerifySequenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		panic("请输入至少包含一个元素的后序遍历序列")
	}
	return verifySequenceOfBSTCore(sequence)
}

func verifySequenceOfBSTCore(sequence []int) bool {
	if len(sequence) <= 0 {
		return true
	}
	root := sequence[len(sequence)-1]
	splitIndex := 0
	for sequence[splitIndex] < root {
		splitIndex++
	}
	// 判断后面的节点是否都大于root
	for i := splitIndex; i < len(sequence)-1; i++ {
		if sequence[i] < root {
			return false
		}
	}
	// 同时验证左右子树
	return verifySequenceOfBSTCore(sequence[:splitIndex]) &&
		verifySequenceOfBSTCore(sequence[splitIndex:len(sequence)-1])
}
