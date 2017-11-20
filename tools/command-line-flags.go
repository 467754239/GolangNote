package main

import (
	"flag"
	"fmt"
)

var (
	wordPtr = flag.String("word", "foo", "a string")
	numbPtr = flag.Int("numb", 42, "an int")
	boolPtr = flag.Bool("fork", false, "a bool")
	svar    string
)

func main() {
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println(*wordPtr)
}
