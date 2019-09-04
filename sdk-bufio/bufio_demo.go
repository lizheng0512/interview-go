package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("关闭文件失败！")
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
