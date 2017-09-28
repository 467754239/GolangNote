package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式3.2 小文件一次性读取.
	*/
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf))

}
