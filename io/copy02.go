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

	buf := make([]byte, 1024)
	for {
		n, err := fd.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(buf[:n])
	}
}
