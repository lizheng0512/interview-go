package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if err := ioutil.WriteFile("./testfile", []byte("123"), 0666); err != nil {
		fmt.Println(err)
	}
}
