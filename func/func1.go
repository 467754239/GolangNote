package main

import "fmt"

// 语法格式
func add(x int, y int) int {
	return x + y
}

// 参数类型省略
func f(i, j, k int, s, t string) error {
	return nil
}

func f2(i int, j int, k int, s string, t string) error {
	return nil
}

/*
	多返回值
	是go错误处理的基础，几乎在所有方法调用都使用了多返回值。
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
	命名返回值
*/
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}

func main() {
	fmt.Println(add(3, 6))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//fmt.Println(split(17))
	x1, y1 := split(17)
	fmt.Printf("x1:%d, y1:%d\n", x1, y1)
}
