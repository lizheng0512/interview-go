package stack

// 题目30: 定义一个栈, 实现一个得到栈内最小值的方法, 并且要求栈的min, push, pop时间复杂度都是O(1)
type stackWithMin struct {
	dataStack *Stack
	minStack  *Stack
}

func NewStackWithMin() *stackWithMin {
	return &stackWithMin{New(), New()}
}

func (swm *stackWithMin) Min() (int, bool) {
	peek, ok := swm.minStack.Peek()
	if ok {
		return peek.(int), ok
	} else {
		return 0, ok
	}
}

func (swm *stackWithMin) Pop() {
	value, ok := swm.dataStack.Pop()
	if ok {
		intValue := value.(int)
		min, ok2 := swm.Min()
		if ok2 && intValue == min {
			swm.minStack.Pop()
		}
	}
}

func (swm *stackWithMin) Push(v int) {
	swm.dataStack.Push(v)
	min, ok := swm.Min()
	if ok {
		if min > v {
			swm.minStack.Push(v)
		}
	} else {
		swm.minStack.Push(v)
	}
}
