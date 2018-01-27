package main

import (
	"fmt"
)

func worker(ch chan string) {
	for url := range ch {
		fmt.Println(url)
	}

}

func main() {
	urls := []string{"http://www1.baidu.com", "http://www2.baidu.com", "http://www3.baidu.com"} // 1000个url

	// 以并发10个抓取url
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	for _, url := range urls {
		ch <- url
	}
}
