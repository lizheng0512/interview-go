package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	sum := md5.Sum([]byte("fasdjklfjasdkljgds"))
	fmt.Printf("%x\n", sum)
	unix := time.Now().Unix()
	fmt.Printf("%x\n", unix)
	fmt.Println(len(`{"x":1234,"y":11,"w":200,"h":200,"prob":0.9999}`))
	str := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	part := str[:1:1]
	fmt.Println(part, len(part), cap(part))
}
