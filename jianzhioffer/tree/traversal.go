package tree

import (
	"fmt"
	"interview-go/jianzhioffer/stack"
)

type BinaryTreeNode struct {
	Value               interface{}
	Parent, Left, Right *BinaryTreeNode
}

func (n *BinaryTreeNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%v", n.Value)
}

func PreOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		fmt.Printf("%s ", root)
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	}
}

func InOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Printf("%s ", root)
		InOrderTraversal(root.Right)
	}
}

func PostOrderTraversal(root *BinaryTreeNode) {
	if root != nil {
		PostOrderTraversal(root.Left)
		PostOrderTraversal(root.Right)
		fmt.Printf("%s ", root)
	}
}

// 题目32: 将二叉树从上到下按行遍历, 每行从左往右, 打印一行之后换行
func WidthFirstTraversal(root *BinaryTreeNode) {
	if root == nil {
		fmt.Println()
		return
	}
	// 借助一个队列, 将待打印的节点放入队列, 每次打印一个节点的时候, 如果这个节点有子节点, 就将子节点放入队列的尾部
	queue := stack.NewQueue()
	queue.Append(root)
	// 下一行总数
	nextLevelCount := 0
	// 当前行待打印的总数
	toBePrintCount := 1
	for head, ok := queue.DeleteHead(); ok; head, ok = queue.DeleteHead() {
		headNode := head.(*BinaryTreeNode)
		if headNode.Left != nil {
			queue.Append(headNode.Left)
			nextLevelCount++
		}
		if headNode.Right != nil {
			queue.Append(headNode.Right)
			nextLevelCount++
		}
		// 打印一次
		fmt.Print(headNode.Value, " ")
		toBePrintCount--
		// 当前行待打印的打印完之后, 打印换行符, 并把下一行统计好的数量赋值给toBePrintCount, 同时nextLevelCount置为0
		if toBePrintCount == 0 {
			fmt.Println()
			toBePrintCount = nextLevelCount
			nextLevelCount = 0
		}
	}
}

// 题目: 按照之字型逐行打印二叉树, 第一行从左到右, 第二行从右到左...
func WidthFirstTraversal2(root *BinaryTreeNode) {
	if root == nil {
		fmt.Println()
		return
	}
	// 利用两个栈来实现, 一个存放奇数行的子节点(从左到右), 一个存放偶数行的子节点(从右到左)
	stacks := []*stack.Stack{stack.New(), stack.New()}
	// 用来区分奇偶行, current = 0 表示奇数行
	current := 0
	next := 1
	stacks[current].Push(root)

	for !stacks[0].Empty() || !stacks[1].Empty() {
		pop, _ := stacks[current].Pop()
		popNode := pop.(*BinaryTreeNode)
		fmt.Print(popNode.Value, " ")

		if current == 0 {
			if popNode.Left != nil {
				stacks[next].Push(popNode.Left)
			}
			if popNode.Right != nil {
				stacks[next].Push(popNode.Right)
			}
		} else {
			if popNode.Right != nil {
				stacks[next].Push(popNode.Right)
			}
			if popNode.Left != nil {
				stacks[next].Push(popNode.Left)
			}
		}

		// 打印完当前行之后
		if stacks[current].Empty() {
			fmt.Println()
			current = 1 - current
			next = 1 - next
		}
	}
}
