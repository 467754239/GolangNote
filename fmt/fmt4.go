package main

import "fmt"

func main() {
	var arr = [3]int{1, 2, 3}
	fmt.Printf("%v, %p, %T\n", arr, &arr, arr)

	var s = []int{1, 2, 3}
	fmt.Printf("%v, %p, %T\n", s, &s, s)
	fmt.Printf("%v, %p, %T\n", s, s, s)
}
