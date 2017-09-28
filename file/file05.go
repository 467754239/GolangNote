package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	/*
		方式4 按行读取，按分隔符读取.
	*/
	r := bufio.NewReader(fd)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

}
