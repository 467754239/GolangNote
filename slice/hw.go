package main

import (
	"fmt"
	"log"
)

// 反转字符串
func reverseString(s string) (string, error) {
	var ss string
	for i := len(s); i > 0; i-- {
		ss = fmt.Sprintf("%v%v", ss, string(s[i-1]))
	}
	return ss, nil
}

// 反转切片
func reverseSlice(s []int) ([]int, error) {
	var ss []int
	for i := len(s); i > 0; i-- {
		ss = append(ss, s[i-1])
	}
	return ss, nil

}

// 反转切片的2部分
func reverseSSlice(s []int, n int) ([]int, error) {

	s1 := s[n:]
	s2, _ := reverseSlice(s[:n])
	fmt.Println(s1)
	fmt.Println(s2)
	s2 = append(s2, s1...)
	return s2, nil
}

func main() {
	sli := []int{1, 2, 3, 8, 10}
	s1, err := reverseSlice(sli)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s1)

	str := "1x3u4aa566"
	s2, err := reverseString(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s2)

	var n int = 2
	s := []int{1, 3, 5, 7, 9}
	s3, _ := reverseSSlice(s, n)
	fmt.Println(s3)

}
