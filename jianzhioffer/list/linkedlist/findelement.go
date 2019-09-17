package linkedlist

// 题目22: 查找单向链表中的倒数第k个节点, 最后一个节点为倒数第1个
func FindLastK(head *element, k int) *element {
	if head == nil || k <= 0 {
		return nil
	}

	front, behind := head, head

	// 让front先指向第倒数k+1
	for i := 0; i < k-1; i++ {
		// 考虑链表长度不足k的情况
		if front.next != nil {
			front = front.next
		} else {
			return nil
		}
	}

	// behind开始跟着一起往下遍历, 直到front到达链表的尾部
	for front.next != nil {
		front = front.next
		behind = behind.next
	}
	return behind
}

// 题目23: 链表中存在环, 找到环的入口节点
func FindEntranceOfRing(head *element) *element {
	if head == nil {
		return nil
	}
	// 定义两个一快一慢的指针, 当慢的指针与快的指针相遇的时候, 则在环中
	behind := head.next
	if behind == nil {
		// 长度为1
		return nil
	}
	front := behind.next
	var meetingNode *element
	for front != nil && behind != nil {
		if front == behind {
			meetingNode = front
			break
		}
		// 慢的指针一次走一步
		behind = behind.next
		// 快的指针一次走两步
		front = front.next
		if front != nil {
			front = front.next
		}
	}

	// 如果meetingNode为nil, 则说明链表中没有环, 否则一定会相遇
	if meetingNode == nil {
		return nil
	}

	// 从相遇的节点出发, 再次回到这个节点, 则可以得到环的长度n
	n := 1
	start := meetingNode
	for start.next != meetingNode {
		start = start.next
		n++
	}

	// 重新声明两个指针, 从头节点开始, 第一个指针先走环的长度n步, 第二个指针再开始走, 当它们相遇时则为环的入口节点
	front, behind = head, head
	for i := 0; i < n; i++ {
		front = front.next
	}

	for front != behind {
		front = front.next
		behind = behind.next
	}

	return front
}
