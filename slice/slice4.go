package main

import "fmt"

func printSlice(s []int) {
	// cap显示slice潜在的容量
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4] // 对slice进行扩容，但是不会引起内存分配，因此在cap范围内。
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)

}
