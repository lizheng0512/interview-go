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

func TestSum16(t *testing.T) {
	decimal := convertHex("A8")
	fmt.Println(decimal)
	hex := toHex(decimal)
	fmt.Println(hex)
	fmt.Println(Sum16("a8", "3"))
}
