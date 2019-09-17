package linkedlist

type intElement struct {
	value int
	next  *intElement
}

// 题目25: 将两个递增排序链表合并为一个递增排序链表
func Merge(head1 *intElement, head2 *intElement) *intElement {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	var mergedHead *intElement
	if head1.value < head2.value {
		mergedHead = head1
		mergedHead.next = Merge(head1.next, head2)
	} else {
		mergedHead = head2
		mergedHead.next = Merge(head1, head2.next)
	}

	return mergedHead
}
