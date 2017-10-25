package main

import "fmt"

func main() {
	// make只能针对slice 不能对arr

	// 5代表初始长度，而不是初始元素。
	a := make([]int, 5)
	// a := []int{0, 0, 0, 0, 0}
	printSlice("a", a)

	// 0 代表初始长度
	// 5 代表初始容量
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
