package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(200)
		fmt.Println(writer.Write([]byte("你好")))
	})
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("启动服务失败！", err)
	}
}
