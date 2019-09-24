package linkedlist

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	n1 := &intElement{value: 1}
	n2 := &intElement{value: 2}
	n3 := &intElement{value: 3}
	n4 := &intElement{value: 4}
	n5 := &intElement{value: 5}
	n6 := &intElement{value: 6}
	n7 := &intElement{value: 7}
	n1.next = n3
	n3.next = n4
	n4.next = n5
	//n2.next = n4
	//n4.next = n6
	n6.next = n7
	mergedHead := Merge(n1, n2)
	for mergedHead != nil {
		fmt.Println(mergedHead.value)
		mergedHead = mergedHead.next
	}
}
