package linkedlist

// 将链表[1,2,3,4,5]转换为[1,5,2,4,3]
func ReorderList(head *element) {
	if head == nil || head.next == nil {
		return
	}
	// 将一个链表从中间拆分成两个链表
	fast, slow := head.next, head
	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}
		slow = slow.next
	}
	head2 := slow.next
	slow.next = nil
	// 反转后面的链表
	var prev2, cur2 *element
	cur2 = head2
	for cur2 != nil {
		next := cur2.next
		cur2.next = prev2
		prev2 = cur2
		cur2 = next
	}
	// 此时prev2为反转后的头节点
	for head != nil && prev2 != nil {
		next := head.next
		head.next = prev2
		next2 := prev2.next
		prev2.next = next
		head = next
		prev2 = next2
	}
}
