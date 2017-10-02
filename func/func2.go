package main

import "fmt"

func printV1() {
	fmt.Println("print version1")
}

func printV2() {
	fmt.Println("print version2.")
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

/*
type fstruct struct {
	Func func()
}
*/
func main() {
	/*
		函数是一等公民

		函数一是一种类型
		函数变量可以赋值多次
		参数是空 返回值也是空
	*/
	var f func()
	f = printV1
	f()

	f = printV2
	f()

	/*
		var s []string
		s = []string{"w", "a"}
		fmt.Println(s)
	*/

	// 定义一个函数列表
	var flist [2]func()
	flist = [2]func(){printV1, printV2}
	fmt.Println(flist)

	// 定义一个函数切片
	var fslice []func()
	fslice = []func(){printV1, printV2}
	fmt.Println(fslice)

	// 定义一个函数map
	var fmap map[string]func()
	fmap = map[string]func(){
		"v1": printV1,
		"v2": printV2,
	}
	fmt.Println(fmap)

	var f1 func(x, y int) int
	f1 = add
	sum1 := f1(2, 3)
	fmt.Println("sum1:", sum1)
}
