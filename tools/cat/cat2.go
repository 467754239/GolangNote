package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	buf, err := printFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
}
