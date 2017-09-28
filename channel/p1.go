package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func printUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url, resp.Status)
}

func work(ch chan string) {
	url := <-ch
	printUrl(url)
}

func main() {
	ch := make(chan string)
	go work(ch)

	url := "https://zhuanlan.zhihu.com/auxten"
	ch <- url

	close(ch)
	time.Sleep(time.Second * 1)
}
