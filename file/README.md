## 操作文件的多种方式

- 原生读取，也称为裸读取，很少使用。
```
package main

import (
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
		方式1. 裸读取，很少使用。
	*/
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fd.Read(buf)
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
```
