package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
)

func fileMd5sum(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	md5sum := md5.Sum(content)
	return fmt.Sprintf("%x", md5sum), nil
}

func main() {
	/*
		const Size = 16
		16字节的数组  32个字符

		16进制编码
		单个16进制数取值范围 0 ~ -128
	*/

	/*
		data := []byte("hello")
		md5sum := md5.Sum(data)
		// %x 可以打印一个数的16进制
		fmt.Printf("%x\n", md5sum)
	*/

	content, err := fileMd5sum("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}
