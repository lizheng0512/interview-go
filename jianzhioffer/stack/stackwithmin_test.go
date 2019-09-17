package stack

import (
	"fmt"
	"testing"
)

func TestStackWithMin(t *testing.T) {
	stackWithMin := NewStackWithMin()
	stackWithMin.Push(2)
	fmt.Println(stackWithMin.Min())
	stackWithMin.Pop()
	fmt.Println(stackWithMin.Min())
	stackWithMin.Push(2)
	stackWithMin.Push(3)
	fmt.Println(stackWithMin.Min())
	stackWithMin.Push(1)
	fmt.Println(stackWithMin.Min())
	stackWithMin.Push(3)
	fmt.Println(stackWithMin.Min())
	stackWithMin.Pop()
	fmt.Println(stackWithMin.Min())
	stackWithMin.Pop()
	fmt.Println(stackWithMin.Min())
	stackWithMin.Pop()
	fmt.Println(stackWithMin.Min())
}
