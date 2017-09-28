package main

import (
	"bufio"
	"fmt"
	"io"
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
		方式2. 加上buffer的读取，很高效。
	*/
	r := bufio.NewReader(fd)

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		//fmt.Println(string(buf[:n]))
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Print(string(chunks))
}
