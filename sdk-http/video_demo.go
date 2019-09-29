package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in")
		video, err := os.Open("./毕业晚会唱歌.mov")
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "video/mp4")
		http.ServeContent(w, r, "123", time.Now(), video)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
