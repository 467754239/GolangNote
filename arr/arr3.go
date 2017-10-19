package main

import "fmt"

func main() {
	// 数组变量赋值
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)

	var a2 [3]int
	a2 = a1
	fmt.Println(a2)

	// 数组比较
	/*
		相等的标准是什么？
		- 数组中元素的个数、类型、每个元素都要一样。
	*/
	fmt.Println(a1 == a2)
	fmt.Println(&a1 == &a2)
}
