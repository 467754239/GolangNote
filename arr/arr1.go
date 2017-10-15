package main

import "fmt"

func main() {
	/*
		[3]int 3表示容量 代表只能存放三个元素 否则报错 索引越界
	*/
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])

	// 迭代的方式 而非全内存拷贝
	for i, v := range arr {
		fmt.Printf("index: %d, val: %d\n", i, v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	for i := range arr {
		fmt.Println("i >", i)
	}

	var arr1 [2]int
	arr1 = [2]int{1, 3}
	fmt.Printf("arr1: %v\n", arr1)

	var arr2 [2]int
	arr2[0] = 5
	arr2[1] = 3
	arr2[1] = 7
	//	arr2[3] = 7
	fmt.Printf("arr2 :%v\n", arr2)

	// 数组的初始化
	var q1 [3]int = [3]int{1, 2, 3}
	var q2 [3]int = [3]int{1, 2}
	fmt.Println("q1: ", q1)
	fmt.Println("q2: ", q2[2])

	// 初始化不指定元素的个数，让系统自动识别。
	q3 := [...]int{1, 2, 3, 4}
	fmt.Println(q3)

	// 比较怪异的用法，根据索引来进行初始化
	// 索引为4的初始化为2，索引为10的初始化为-1，其它的为int类型的默认值0
	// 初始化稀疏数组的方式 很难用到它
	q4 := [...]int{4: 2, 10: -1}
	fmt.Println(q4)
}
