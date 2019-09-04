package main

import (
	"fmt"
	"strings"
)

func main() {
	// 根据空格切分
	fields := strings.Fields("a,b c dd")
	for e := range fields {
		fmt.Println(fields[e])
	}
}
