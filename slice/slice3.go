package main

import "fmt"

func main() {
	/*
		切片始化
		分2步
		1. 声明一个匿名数组
		2. 对数组全部切片，然后复制给变量。
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)
	s = s[1:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}
