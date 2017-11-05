package main

import "fmt"

func main() {
	s := "golang你好"
	fmt.Println("len s: ", len(s))

	cnt := 0
	for _, v := range s {
		fmt.Println(v)
		cnt += 1
	}
	fmt.Printf("cnt1: %v\n", cnt)

	cnt = 0
	for _, v := range []byte(s) {
		fmt.Println(v)
		cnt += 1
	}
	fmt.Printf("cnt2: %v\n", cnt)
}
