package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6, 7)
	reversedHead := Reverse(list.getElem(0))
	for reversedHead != nil {
		fmt.Println(reversedHead.value)
		reversedHead = reversedHead.next
	}
}
