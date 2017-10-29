package main

import "fmt"

func main() {
	// 创建map
	m1 := make(map[string]string)
	m1["name"] = "zhengyscn"
	m1["address"] = "bj"
	fmt.Println(m1)

	// or

	// 创建并且初始化
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m2)
}
