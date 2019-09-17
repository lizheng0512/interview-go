package linkedlist

import (
	"fmt"
	"testing"
)

func TestFindLastK(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(FindLastK(list.getElem(0), 8))
}

func TestFindEntranceOfRing(t *testing.T) {
	list := New(1, 2, 3, 4, 5, 6, 7)
	list.getElem(6).next = list.getElem(0)
	fmt.Println(FindEntranceOfRing(list.getElem(0)))
}
