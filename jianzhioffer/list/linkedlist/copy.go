package linkedlist

import (
	"fmt"
	"strings"
)

// 题目35: 复制一个复杂链表, 复杂链表的节点结构如下, 除了一个指向下一个节点的指针还有一个指向链表中任意节点的指针
type ComplexListNode struct {
	Value         interface{}
	Next, Sibling *ComplexListNode
}

func (cln *ComplexListNode) String() string {
	if cln == nil {
		return ""
	}
	strArr := make([]string, 0)
	for cln != nil {
		strArr = append(strArr, fmt.Sprintf("value: %v, next: %v, sibling: %v",
			cln.valueString(), cln.Next.valueString(), cln.Sibling.valueString()))
		cln = cln.Next
	}
	return strings.Join(strArr, "\n")
}

func (cln *ComplexListNode) valueString() string {
	if cln != nil {
		return fmt.Sprintf("%v", cln.Value)
	}
	return "nil"
}

// 复制复杂链表分为3步
// 第一步将所有节点拷贝一份, 拷贝的节点放到原始节点的next,
// 第二部遍历整个链表, 为拷贝节点添加sibling指针, 原始节点的sibling指向的节点的next就是原始节点的拷贝节点的sibling
// 第三步拆分两个链表, 奇数位置的节点是原始链表的节点, 偶数位置的节点是拷贝节点
func CopyComplexList(head *ComplexListNode) *ComplexListNode {
	if head == nil {
		return nil
	}
	CopyComplexListNode(head)
	createSiblingForCopiedNode(head)
	return splitList(head)
}

// 复制节点
func CopyComplexListNode(head *ComplexListNode) {
	for head != nil {
		copiedNode := &ComplexListNode{Value: head.Value, Next: head.Next}
		head.Next = copiedNode
		head = copiedNode.Next
	}
}

// 为复制出来的节点创建sibling指针
func createSiblingForCopiedNode(head *ComplexListNode) {
	for head != nil {
		copiedNode := head.Next
		if head.Sibling != nil {
			copiedNode.Sibling = head.Sibling.Next
		}
		head = copiedNode.Next
	}
}

// 拆分两个链表, 返回复制的链表的头节点
func splitList(head *ComplexListNode) (copiedHead *ComplexListNode) {
	var copiedNode *ComplexListNode
	if head != nil {
		copiedHead = head.Next
		copiedNode = copiedHead
		head.Next = copiedNode.Next
		head = head.Next
	} else {
		return
	}
	for head != nil {
		copiedNode.Next = head.Next
		copiedNode = copiedNode.Next
		head.Next = copiedNode.Next
		head = head.Next
	}
	return
}
