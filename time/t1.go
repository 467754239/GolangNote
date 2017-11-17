package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())

	var n time.Duration
	n = 2 * time.Hour
	fmt.Println(int64(n))
	fmt.Println(n.String())

	n = 3*time.Hour + 30*time.Minute
	fmt.Println(n.String())
}
