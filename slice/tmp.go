package main

import "fmt"

func main() {

	/*
		slice : 引用
	*/

	var s1 []int
	printSlice(s1)
	fmt.Println("s1 == nil", s1 == nil)

	s2 := []int{2, 3, 5, 7, 11}
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)

	// make(type, len, cap)
	s3 := make([]int, 5)
	printSlice(s3)

	s4 := make([]int, 5, 7)
	printSlice(s4)

	// define arr
	primes := [8]int{2, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println(primes)

	var s6 []int = primes[1:4]
	printSlice(s6)

	// 切片的字面量
	s7 := []int{1, 2, 5, 7, 11}
	fmt.Println(s7)
	s8 := []bool{true, false, true, false}
	fmt.Println(s8)

	// append
	var s9 []int
	for i := 0; i < 10; i++ {
		s9 = append(s9, i)
	}
	fmt.Printf("s9:%v\n", s9)

	var s10 []int = []int{10, 11, 12}
	s9 = append(s9, s10...)
	fmt.Printf("s9:%v\n", s9)

	s9 = append(s9, 15)
	s9 = append(s9, 16, 17, 18, 19)
	fmt.Printf("s9:%v\n", s9)

	// update
	s11 := []string{"a1", "a2"}
	fmt.Printf("s11:%v\n", s11)
	s11[0] = "aa1"
	fmt.Printf("s11:%v\n", s11)

	for i, v := range primes {
		fmt.Println(i, v)
	}

}

func printSlice(s []int) {
	fmt.Printf("mem_add=%p, len=%d cap=%d %v\n", s, len(s), cap(s), s)
}
