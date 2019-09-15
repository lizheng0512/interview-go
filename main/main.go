package main

import (
	"fmt"
	"math"
	"strconv"
)

type A struct {
	str string
}

// 导致无限递归，报协程栈溢出
func (a *A) String() string {
	return fmt.Sprintf("%s ", a)
}

func main() {
	a := "012"
	pa := &a
	b := []byte(*pa)
	pb := &b

	a += "3"
	*pa += "4"
	b[1] = '5'

	fmt.Println(*pa)
	fmt.Println(string(*pb))

	var m = map[rune]uint8{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12, 'M': 13, 'N': 14,
		'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26, '我': 1}
	fmt.Println(m)

	fmt.Println(-1 << 1)
	fmt.Println(-1 >> 100)
	fmt.Println(-4 >> 1)
	fmt.Println(-1 & 0x82)
	fmt.Println(int8(-0x80))
	fmt.Println((10 << 100) & (1<<63 - 1))
	fmt.Println(1<<63 - 1)
	fmt.Println(0x0f ^ 0x01)
	testDefer()
	fmt.Println(^uint(0))
	fmt.Println(^uint32(0) >> 63)

	fmt.Println(strconv.Atoi("9223372036854775807"))

	n := uint(18446744073709551600)
	maxVal := uint(1)<<uint(32) - 1
	fmt.Println(n+9 > maxVal)

	fmt.Println(-(-100))

	fmt.Println(math.Pow(1<<63, 1<<63))
	fmt.Println(uint(-(-10)))
}

func testDefer() {
	arr := []int{1, 2, 3, 4, 5, 0}
	defer printLast(arr)
	arr[len(arr)-1] = 6
}

func printLast(arr []int) {
	if len(arr) > 0 {
		fmt.Println(arr[len(arr)-1])
	}
}
