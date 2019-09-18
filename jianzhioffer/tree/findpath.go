package tree

import (
	"fmt"
	"interview-go/jianzhioffer/list/linkedlist"
)

// 题目34: 在二叉树中打印出路径和为输入值的路径
func FindPath(root *IntBinaryTreeNode, expectedSum int) {
	if root == nil {
		return
	}
	findPathCore(root, expectedSum, 0, linkedlist.New())
}

func findPathCore(node *IntBinaryTreeNode, expectedSum, currentSum int, path *linkedlist.List) {
	// 增加当前路径和
	currentSum += node.Value
	// 在路径中添加节点值
	path.Add(node.Value)

	isLeaf := node.Left == nil && node.Right == nil

	// 满足条件, 打印路径
	if isLeaf && currentSum == expectedSum {
		iter := path.Iterator()
		for iter.Next() {
			fmt.Print(iter.Value(), " ")
		}
		fmt.Println()
	}

	if node.Left != nil {
		findPathCore(node.Left, expectedSum, currentSum, path)
	}
	if node.Right != nil {
		findPathCore(node.Right, expectedSum, currentSum, path)
	}

	// 递归返回时弹出当前节点
	if path.Size() > 0 {
		path.Remove(path.Size() - 1)
	}
}
