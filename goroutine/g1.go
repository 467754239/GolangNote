package main

import (
	"fmt"
	"time"
)

// 斐波那契数列
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		fmt.Println("delay start.")
		time.Sleep(delay)
		fmt.Println("dealy finish.")
	}
}

func main() {
	go spinner(time.Second * 10)
	fib(45)
}
