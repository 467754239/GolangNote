package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// a 和 b 有一部分指向names的同一块内存地址
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

	c := a[1:2]
	fmt.Println(c)

	d := a[1:4]
	fmt.Println(d)

	var p [2]*string
	p[0] = &names[0]
	p[1] = &names[1]
	*p[0] = "AAA"
	fmt.Println(*p[0])

	/*
		var start *string
		var length int
		var cap int
	*/
}
