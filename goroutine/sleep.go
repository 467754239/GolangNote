package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{3, 2, 1, 7, 9}
	for _, n := range s {
		go func(n int) {
			time.Sleep(time.Second * time.Duration(n))
			fmt.Println(n)
		}(n)
	}
	time.Sleep(10 * time.Second)

}
