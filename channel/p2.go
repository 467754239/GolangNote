package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	1. 给定一个URL，打印URL和URL的status.
*/
func printUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)
}

func work(ch chan string) {
	for {
		url, ok := <-ch
		// ok 判断channel是否被关闭。
		// ok为false表示channel被关掉了。
		if !ok {
			break
		}
		printUrl(url)
	}
}

func main() {
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go work(ch)
	}

	for i := 0; i < 5; i++ {
		url := "https://zhuanlan.zhihu.com/auxten"
		ch <- url
	}
	fmt.Println("close ch")
	close(ch)

	time.Sleep(3 * time.Second)
	fmt.Println("Over")

}
