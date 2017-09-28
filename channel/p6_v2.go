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
	for url := range ch {
		printUrl(url)
	}
	wg.Done()
}

func main() {
	// https://golang.org/pkg/sync/#example_WaitGroup
	wg := new(sync.WaitGroup)
	wg.Add(3)

	ch := make(chan string)

	for i := 0; i < 3; i++ {
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
