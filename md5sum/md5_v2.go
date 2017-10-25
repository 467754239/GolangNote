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
		md5是一个16字节的数组

		linux
		$ md5sum filename
		16进制编码
		md5sym 结果是一个16字节的数组
	*/
	content, err := fileMd5sum("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
}
