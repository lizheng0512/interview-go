package stack

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Append(1)
	queue.Append(2)
	queue.Append(3)
	queue.Append(4)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
	queue.Append(5)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue)
}
