package string

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	fmt.Println(Match("123", "123"))
	fmt.Println(Match("123", ".2*3"))
	fmt.Println(Match("123", "1*2*3"))
}
