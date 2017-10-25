package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3)
	printSlice(s)

	// +
	s1 := []int{2, 3, 4}
	s = append(s, s1...) // s1打散
	fmt.Println(s)

}

func printSlice(s []int) {
	// cap显示slice潜在的容量
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
