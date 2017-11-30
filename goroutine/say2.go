package main

import (
	"fmt"
	"time"
)

func say(s string, c int) {
	for i := 0; i < c; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}

}

func main() {
	go say(" world", 10)
	say("hello", 5)
}
