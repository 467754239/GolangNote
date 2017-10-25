package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	/*
		data := []byte("hello")
		md5sum := md5.Sum(data)
		fmt.Printf("%x\n", md5sum)
	*/
	md5sum := md5.Sum([]byte("hello"))
	fmt.Printf("%x\n", md5sum)

	md5sum1 := md5.Sum([]byte("hello1"))
	fmt.Println(md5sum1)

	if md5sum == md5sum1 {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}

}
