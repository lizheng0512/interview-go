package bitoperation

import (
	"fmt"
	"testing"
)

func TestCal26ToDecimal(t *testing.T) {
	fmt.Println(Cal26ToDecimal("AB"))
}

func TestCount1(t *testing.T) {
	fmt.Println(Count1(0))
	fmt.Println(Count1Optimum(-16))
}
