package main

import "fmt"

/*
	https://golang.org/doc/effective_go.html#channels
*/
func main() {
	urls := []string{"http://www1.baidu.com", "http://www2.baidu.com", "http://www3.baidu.com"} // 1000个url

	// 以并发10个抓取url
	token := make(chan bool, 3)

	for _, url := range urls {
		token <- true
		go func(url string) {
			fmt.Println(url)
			<-token
		}(url)
	}
}
