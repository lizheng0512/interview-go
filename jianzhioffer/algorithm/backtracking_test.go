package algorithm

import (
	"fmt"
	"testing"
)

func TestFindPathInMatrix(t *testing.T) {
	matrix := [][]string{
		{"a", "b", "c", "d"},
		{"e", "f", "g", "h"},
		{"i", "j", "k", "l"},
		{"m", "n", "o", "p"},
	}
	fmt.Println(FindPathInMatrix(matrix, "abcgfj1klponmie"))
}

func TestMovingCount(t *testing.T) {
	fmt.Println(MovingCount(4, 4, 2))
}
