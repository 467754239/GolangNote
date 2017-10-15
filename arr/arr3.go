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
		一样的标准是什么？
			1. 数组的元素个数一样
			2. 数组的类型一样
			3. 数组中每个元素一样
	*/
	fmt.Println(a1 == a2)
	fmt.Println(&a1 == &a2)
}
