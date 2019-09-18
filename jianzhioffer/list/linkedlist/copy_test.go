package linkedlist

import (
	"fmt"
	"testing"
)

func TestCopyComplexList(t *testing.T) {
	n1 := &ComplexListNode{Value: 1}
	n2 := &ComplexListNode{Value: 2}
	n3 := &ComplexListNode{Value: 3}
	n4 := &ComplexListNode{Value: 4}
	n5 := &ComplexListNode{Value: 5}
	n6 := &ComplexListNode{Value: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n1.Sibling = n4
	n2.Sibling = n3
	n4.Sibling = n1
	n5.Sibling = n2
	n6.Sibling = n6
	fmt.Println(n1)
	fmt.Println()
	copyComplexList := CopyComplexList(n1)
	fmt.Println(n1)
	fmt.Println()
	fmt.Println(copyComplexList)
	fmt.Println(CopyComplexList(nil))
}
