// 双向链表
package linkedlist

import (
	"fmt"
	"strings"
)

type List struct {
	first, last *element
	size        int
}

type element struct {
	prev, next *element
	value      interface{}
}

func (e *element) String() string {
	if e != nil {
		return fmt.Sprintf("%v", e.value)
	}
	return ""
}

type Iterator struct {
	list  *List
	elem  *element
	index int
}

func New(values ...interface{}) *List {
	l := new(List)
	l.Add(values...)
	return l
}

func (l *List) Add(values ...interface{}) {
	for _, v := range values {
		e := new(element)
		e.value = v
		if l.size == 0 {
			l.first = e
		} else {
			e.prev = l.last
			l.last.next = e
		}
		l.last = e
		l.size++
	}
}

func (l *List) Get(index int) (interface{}, bool) {
	if index < 0 || index >= l.size {
		return nil, false
	}

	e := l.getElem(index)
	return e.value, true
}

func (l *List) Remove(index int) {
	if index < 0 || index >= l.size {
		return
	}

	e := l.getElem(index)
	if index == 0 {
		l.first = e.next
	} else if index == l.size-1 {
		l.last = e.prev
	}
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
	l.size--
}

func (l *List) getElem(index int) *element {
	var e *element
	var i int
	if l.size-index > index {
		// 从前往后
		for e = l.first; i != index; e, i = e.next, i+1 {
		}
	} else {
		// 从后往前
		for e, i = l.last, l.size-1; i != index; e, i = e.prev, i-1 {
		}
	}
	return e
}

// 按照values顺序在头部添加一些元素
func (l *List) Prepend(values ...interface{}) {
	if len(values) > 0 {
		var last, newFirst *element
		for i, v := range values {
			e := new(element)
			e.value = v
			if i == 0 {
				newFirst = e
			} else {
				e.prev = last
				last.next = e
			}
			last = e
		}
		if l.size == 0 {
			l.last = last
		} else {
			l.first.prev = last
			last.next = l.first
		}
		l.first = newFirst
		l.size += len(values)
	}
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Map(f func(v interface{}) interface{}) {
	e := l.first
	for {
		e.value = f(e.value)
		e = e.next
		if e == nil {
			break
		}
	}
}

func (l *List) Set(index int, v interface{}) {
	if index < 0 || index >= l.size {
		return
	}

	e := l.getElem(index)
	e.value = v
}

func (l *List) Insert(index int, v interface{}) {
	if index < 0 || index >= l.size {
		return
	}

	insertedElem := &element{value: v}
	e := l.getElem(index)
	e.prev.next, insertedElem.prev = insertedElem, e.prev
	insertedElem.next, e.prev = e, insertedElem
	l.size++
}

func (l *List) String() string {
	valueStrings := make([]string, 0, l.size)
	for e := l.first; e != nil; e = e.next {
		valueStrings = append(valueStrings, fmt.Sprintf("%v", e.value))
	}
	return fmt.Sprintf("linkedlist[%s]", strings.Join(valueStrings, ", "))
}

func (l *List) Iterator() *Iterator {
	return &Iterator{list: l, elem: nil, index: -1}
}

// for i.Next() {
//     ...
// }
func (i *Iterator) Next() bool {
	i.index++
	if i.index < i.list.size {
		if i.elem == nil {
			i.elem = i.list.first
		} else {
			i.elem = i.elem.next
		}
		return true
	} else {
		i.elem = nil
		return false
	}
}

func (i *Iterator) Prev() bool {
	i.index--
	if i.index >= 0 {
		if i.elem == nil {
			i.elem = i.list.last
		} else {
			i.elem = i.elem.prev
		}
		return true
	} else {
		i.elem = nil
		return false
	}
}

func (i *Iterator) Begin() {
	i.index = -1
	i.elem = nil
}

func (i *Iterator) End() {
	i.index = i.list.size
	i.elem = nil
}

func (i *Iterator) First() bool {
	i.Begin()
	return i.Next()
}

func (i *Iterator) Last() bool {
	i.End()
	return i.Prev()
}

func (i *Iterator) Value() interface{} {
	return i.elem.value
}
