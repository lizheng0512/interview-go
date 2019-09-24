package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		fmt.Println(1)
		wg.Done()
	}()
	go func() {
		fmt.Println(2)
		wg.Done()
	}()
	go func() {
		fmt.Println(3)
		wg.Done()
	}()
	wg.Wait()
}
