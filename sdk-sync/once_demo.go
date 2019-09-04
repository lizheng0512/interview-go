package main

import (
	"fmt"
	"sync"
)

// once保证只执行一次
func main() {
	one := new(sync.Once)
	for i := 0; i < 10; i++ {
		fmt.Println("for in.")
		if i == 1 {
			one.Do(func() {
				fmt.Println("你好。")
			})
		} else {
			one.Do(func() {
				fmt.Println("hello")
			})
		}
	}
}
