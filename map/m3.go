package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m == nil)

	m = make(map[string]int)
	m["a"] = 1
	fmt.Println(m)

}
