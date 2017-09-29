package main

import (
	"bufio"
	"fmt"
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
		方式4 按行读取，按分隔符读取.
	*/
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
