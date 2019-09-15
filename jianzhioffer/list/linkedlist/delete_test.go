package linkedlist

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	elem := list.getElem(3)
	DeleteNode(list, elem)
	fmt.Println(list)
	elem = list.getElem(3)
	DeleteNode(list, elem)
	fmt.Println(list)
	elem = list.getElem(0)
	DeleteNode(list, elem)
	fmt.Println(list)
}

func TestDeleteDuplicateNode(t *testing.T) {
	var list *List
	list = New(1, 2, 3, 3, 5)
	DeleteDuplicateNode(list)
	fmt.Println(list)
	list = New(1, 2, 3, 3)
	DeleteDuplicateNode(list)
	fmt.Println(list)
	list = New(0, 0, 1, 2, 3, 3, 5)
	DeleteDuplicateNode(list)
	fmt.Println(list)
	list = New(0, 0, 1, 1, 2, 2, 3, 3)
	DeleteDuplicateNode(list)
	fmt.Println(list)
}
