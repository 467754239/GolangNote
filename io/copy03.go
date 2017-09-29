package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var fd *os.File
	var err error

	if len(os.Args) > 1 {
		fd, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
	} else {
		fd = os.Stdin
	}

	io.Copy(os.Stdout, fd)
}
