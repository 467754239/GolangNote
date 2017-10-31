package main

import "fmt"

func main() {
	var s string
	var n int

	for {
		fmt.Print("> ")
		fmt.Scanln(&s, &n)
		if s == "stop" {
			break
		}
		fmt.Println(s, n)
	}
}
