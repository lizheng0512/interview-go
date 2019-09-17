package linkedlist

// 题目24: 反转单向链表, 返回反转后的头节点
func Reverse(head *element) *element {
	if head == nil {
		return nil
	}
	if head.next == nil {
		return head
	}
	var prev, cur *element
	cur = head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}
