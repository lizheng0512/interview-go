package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := New()
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	iterator := s.Iterator()
	for iterator.Next() {
		fmt.Println(iterator.Value())
	}
	fmt.Println(s.Pop())
	iterator.Prev()
	for iterator.Prev() {
		fmt.Println(iterator.Value())
	}
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
