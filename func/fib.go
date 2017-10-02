package main

import "fmt"

/*
	递归的方式实现斐波那契数列
	函数的一种调用范式

	唯一不好的就是性能比较差,但代码可读性比较好.

	爬虫,自已网页抓取自已网页中的链接.
*/
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func a(n int) int {
	fmt.Println("enter n: ", n)
	if n <= 1 {
		return 2
	}
	m := 2*a(n-1) + n - 1
	//fmt.Println("return from a(n-1):", n)
	fmt.Printf("return from a(n-1):%d, m:%d\n", n, m)
	return m
	/*
		2 * a9 + 9
		2 * (2 * 8a + 8) + 9

		n == 1的时候终止
	*/
}

/*
	通项公式: a(n) = 2 * a(n-1) + n -1
	第一项: a(1) = 2
	求: 第10项 a(10)
	a(10) = 2 * a(10-1) + 10 - 1
*/

func main() {
	fmt.Println(fib(15))

	fmt.Println(a(10))
}
