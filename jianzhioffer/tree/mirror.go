package tree

// 题目27: 输入一个二叉树, 输出它的镜像
// 思路: 镜像就是从跟节点开始遍历, 前序遍历, 把每个非叶子节点的左右子节点交换位置
func MirrorRecursively(root *BinaryTreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	MirrorRecursively(root.Left)
	MirrorRecursively(root.Right)
}
