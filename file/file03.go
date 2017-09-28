package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
		方式3.1 小文件一次性读取.
	*/
	buf, err := ioutil.ReadFile("/var/log/messages")
	//buf, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf))

}
