package main

import "fmt"

func main() {
	/*
		空切片等于nil， 但长度等于0的切片不一定等于nil
	*/
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}

	a := [...]int{1, 3, 5, 7, 9}
	s1 := a[:0]
	fmt.Println(s1 == nil)
}
