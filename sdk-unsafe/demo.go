package main

import (
	"fmt"
	"unsafe"
)

// 结构体和内存都是在内存中一段连续的内存，指向结构体或者数据的指针相当于指向了结构体的第一个属性或者数据的第一个元素
func main() {
	arr := [3]uint{1, 2, 3}
	// unsafe.Sizeof返回变量所占的字节数
	fmt.Println(unsafe.Sizeof(arr))
	fmt.Println(unsafe.Sizeof(uint8(1)))
	fmt.Println(unsafe.Sizeof(uint32(1)))
	// slice永远是24, 因为slice中包含数组的起始地址, 元素数, 和容量, 每个占8字节, 共计24字节
	fmt.Println(unsafe.Sizeof([]int{1, 2, 3, 4, 5, 6, 7}))
	// unsafe.Offsetof返回结构体中某个属性字节的偏移量
	st := s{}
	fmt.Println(unsafe.Offsetof(st.a), unsafe.Offsetof(st.b))
	// unsafe.Pointer是通用指针类型，可以实现指针类型的转换
	// 利用unsafe.pointer的方式实现修改数组中的元素
	// 取得数组开始的地址, 转换成uintptr是为了参与指针运算
	parr := uintptr(unsafe.Pointer(&arr))
	// 要修改数组中的第三个元素, 计算第三个元素指针
	p := (*uint)(unsafe.Pointer(parr + 2*unsafe.Sizeof(arr[0])))
	*p = 4
	fmt.Println(arr)
	// 修改结构体s的属性b的值
	ps := new(s)
	fmt.Println(unsafe.Offsetof(ps.a), unsafe.Offsetof(ps.b))
	//unsafe.Pointer(uintptr(unsafe.Pointer(ps)) + unsafe.Offsetof())

}

type s struct {
	a uint16
	b uint8
}
