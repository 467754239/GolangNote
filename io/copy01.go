package main

import (
	"log"
	"os"
)

func main() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(buf[:n])
	}
}
