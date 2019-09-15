package stack

import "interview-go/jianzhioffer/list/linkedlist"

type Stack struct {
	list *linkedlist.List
}

func New() *Stack {
	return &Stack{list: linkedlist.New()}
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	value, ok = s.list.Get(0)
	s.list.Remove(0)
	return
}

func (s *Stack) Push(value interface{}) {
	s.list.Prepend(value)
}

func (s *Stack) Peek() (interface{}, bool) {
	return s.list.Get(0)
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) Empty() bool {
	return s.list.Size() == 0
}

type Iterator struct {
	stack *Stack
	index int
}

func (s *Stack) Iterator() *Iterator {
	return &Iterator{stack: s, index: -1}
}

func (i *Iterator) Next() bool {
	i.index++
	return i.index >= 0 && i.index < i.stack.Size()
}

func (i *Iterator) Prev() bool {
	i.index--
	return i.index >= 0 && i.index < i.stack.Size()
}

func (i *Iterator) Value() interface{} {
	value, _ := i.stack.list.Get(i.index)
	return value
}

func (i *Iterator) Begin() {
	i.index = -1
}

func (i *Iterator) End() {
	i.index = i.stack.Size()
}
