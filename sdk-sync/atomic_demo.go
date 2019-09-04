package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// atomic包提供了一系列原子操作，cas、swap、load、store等操作
	i := uint32(1)
	fmt.Println(atomic.CompareAndSwapUint32(&i, 1, 2), i)
}
