package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func worker(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(url, resp.Status)

}

func main() {
	wg := new(sync.WaitGroup)

	urls := []string{
		"http://www.baidu.com/",
		"http://mp3.baidu.com/",
		"http://baike.baidu.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(url)
		}()
	}

	wg.Wait()

}
