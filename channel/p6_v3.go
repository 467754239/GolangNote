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

func work(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range ch {
		printUrl(url)
	}
}

func main() {
	ch := make(chan string)

	// https://golang.org/pkg/sync/#example_WaitGroup
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(ch, wg)
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
