package linkedlist

import "fmt"

// 题目：用O(1)的时间复杂度删除单向链表中的一个节点
// 思路：将node的下一个节点的值赋值给node，并把node的next指针指向下下个节点
func DeleteNode(list *List, node *element) {
	if list == nil || node == nil || list.size == 0 {
		return
	}
	// 如果node是尾节点，也就是没有next，需要从头遍历
	if node.next == nil {
		fmt.Println("node.next is nil.")
		e := list.first
		for e.next != node {
			e = e.next
		}
		if e.next == node {
			e.next = e.next.next
		}
	} else {
		node.value = node.next.value
		node.next = node.next.next
	}
}

// 题目：删除按值升序或者降序排列好的单向链表中重复的节点
func DeleteDuplicateNode(list *List) {
	if list == nil || list.size <= 1 {
		return
	}

	var pPreNode *element
	pNode := list.first
	for pNode != nil {
		if pNode.next != nil && pNode.value == pNode.next.value {
			// 当前重复的值
			value := pNode.value
			// pNode走过所有重复的节点，直到尾部或者不重复
			for pNode != nil && pNode.value == value {
				pNode = pNode.next
			}
			// pPreNode == nil说明第一个元素就是重复的
			if pPreNode == nil {
				list.first = pNode
			} else {
				pPreNode.next = pNode
			}
		} else {
			pPreNode = pNode
			pNode = pNode.next
		}
	}
}
