package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	str := "北京"
	bytes := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		bytes = append(bytes, str[i])
	}
	fmt.Println("\nbytes长度为", len(bytes))
	fmt.Println(bytes)
	str2 := string(bytes)
	fmt.Println(str2)
	bytes[0] = 230
	fmt.Println(bytes)
	fmt.Println(str2)

	str3 := "北京"
	str4 := str3
	fmt.Println(unsafe.Pointer(&str3), unsafe.Pointer(&str4))
	fmt.Println(uintptr(unsafe.Pointer(&str3)), uintptr(unsafe.Pointer(&str4)))
	fmt.Println(unsafe.Sizeof("北"), unsafe.Sizeof(str3))
	jing := (*[]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&str3)) + unsafe.Sizeof("北")))
	fmt.Println(string(*jing))
	byteArr := make([]byte, 4, 4)
	fmt.Println(copy(byteArr, "123"), byteArr)
	fmt.Println(copy(byteArr[3:], "23"), byteArr)

	ss := &s{str: unsafe.Pointer(&str3), len: 12}
	ss2 := ss

	fmt.Println(uintptr(unsafe.Pointer(&ss)), uintptr(unsafe.Pointer(&ss2)))

	//var sb strings.Builder
	sb := new(strings.Builder)
	sb.WriteString("123")
	sb.WriteString("123")
	sb.WriteString("123")
	fmt.Println(sb.String())

	ra := strings.ReplaceAll("123456", "", "-")
	fmt.Println(ra)

	var sb2 strings.Builder
	fmt.Println(sb2.Cap())
	sb.Grow(21)
	fmt.Println(sb2.Cap())
	strings.ReplaceAll()
}

type s struct {
	str unsafe.Pointer
	len int
}

func gostringnocopy(b []byte) string {
	ss := s{str: unsafe.Pointer(&b), len: len(b)}
	s := *(*string)(unsafe.Pointer(&ss))
	return s
}
