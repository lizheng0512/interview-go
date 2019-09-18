package stack

import (
	"fmt"
	"testing"
)

func TestIsPopSequence(t *testing.T) {
	//push := []int{1, 2, 3, 4, 5}
	push := []int{1}
	//pop := []int{4, 5, 3, 2, 1}
	//pop := []int{5, 4, 3, 2, 1}
	//pop := []int{5, 4, 3, 1, 2}
	pop := []int{1}
	fmt.Println(IsPopSequence(push, pop))
	fmt.Println(IsPopSequence(nil, nil))
}
