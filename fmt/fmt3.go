package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var n int

	var line string
	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		//fmt.Scanln(&line)
		line, _ = f.ReadString('\n')
		fmt.Sscanln(line, &s, &n)
		if s == "stop" {
			break
		}
		fmt.Println(s, n)
	}
}
