package stack

import (
	"fmt"
	"strings"
)

// 用两个栈实现基本队列功能
type queue struct {
	s1, s2 *Stack
}

func NewQueue() *queue {
	return &queue{New(), New()}
}

func (q *queue) Append(value interface{}) {
	q.s1.Push(value)
}

func (q *queue) DeleteHead() (interface{}, bool) {
	if q.s2.Empty() {
		for {
			v, ok := q.s1.Pop()
			if ok {
				q.s2.Push(v)
			} else {
				break
			}
		}
	}
	return q.s2.Pop()
}

func (q *queue) String() string {
	i1 := q.s1.Iterator()
	i2 := q.s2.Iterator()
	arr := make([]string, 0, q.s1.Size()+q.s2.Size())
	for i2.Next() {
		arr = append(arr, fmt.Sprintf("%v", i2.Value()))
	}
	i1.End()
	for i1.Prev() {
		arr = append(arr, fmt.Sprintf("%v", i1.Value()))
	}
	return strings.Join(arr, ", ")
}
