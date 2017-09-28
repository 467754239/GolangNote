package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
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
	for url := range ch {
		printUrl(url)
	}
	/*
		for {
			url, ok := <-ch
			// ok 判断channel是否被关闭。
			// ok为false表示channel被关掉了。
			if !ok {
				break
			}
			printUrl(url)
		}
	*/
}

func main() {
	// https://golang.org/pkg/sync/#example_WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(3)

	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			work(ch)
		}()
	}

	for i := 0; i < 5; i++ {
		url := "https://zhuanlan.zhihu.com/auxten"
		ch <- url
	}
	fmt.Println("close ch")
	close(ch)

	// Wait for all HTTP fetches to complete.
	// wait all goroutine finish
	wg.Wait()

}
