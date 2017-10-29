package main

import "fmt"

func main() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m1)

	fmt.Println(m1["a"])
	m1["a"] = m1["b"] + 1
	fmt.Println(m1)

	// 如果c在m1中不存在，则获取到改类型的默认值。
	c1 := m1["c"]
	fmt.Println(c1)

	c2, ok := m1["c"]
	if ok {
		fmt.Println(c2)
	} else {
		fmt.Println("not found.")
	}

	if c, ok := m1["c"]; ok {
		fmt.Println(c)
	}

	// 删除元素
	delete(m1, "a")
	fmt.Println(m1)

	/*
		遍历的2中方式
		1. 遍历key和value的方式
		2. 遍历k的方式
	*/
	m1["name"] = 4
	for k, v := range m1 {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}

	for k := range m1 {
		fmt.Printf("k:%v\n", k)
	}

	// 空map不能做任何操作
	var m2 map[string]int
	fmt.Printf("m2 == nil -> %v\n", m2 == nil)
}
