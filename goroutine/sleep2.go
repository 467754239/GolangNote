package main

import (
	"fmt"
	"time"
)

func main() {
	s := "hello\n"
	//fmt.Printf("%T\n", s[2])
	for i, v := range s {
		go func(i int, v rune) {
			//go func(v int32) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Print(string(v))
		}(i, v)
	}
	time.Sleep(time.Second * 1)
}
