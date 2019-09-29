package linkedlist

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	e1 := &element{value: 1}
	e2 := &element{value: 2}
	e3 := &element{value: 3}
	e4 := &element{value: 4}
	e5 := &element{value: 5}
	e6 := &element{value: 6}
	e7 := &element{value: 7}
	e8 := &element{value: 8}
	//e9 := &element{value:9}
	e1.next = e2
	e2.next = e3
	e3.next = e4
	e4.next = e5
	e5.next = e6
	e6.next = e7
	e7.next = e8
	//e8.next = e9
	ReorderList(e1)
	for e1 != nil {
		fmt.Println(e1.value)
		e1 = e1.next
	}
}
