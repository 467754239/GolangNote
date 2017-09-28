package main

import (
	"io"
	"os"
)

func main() {
	/*
		5. 操作类文件的神器
	*/
	io.Copy(os.Stdout, os.Stdin)
}
