package sort

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	ages := []int{22, 24, 23, 22, 21, 18, 19, 28, 33, 31, 24, 29, 35, 35, 41, 20}
	Case1(ages)
	fmt.Println(ages)
}
