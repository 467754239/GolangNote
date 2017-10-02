package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func add(m, n int) int {
	return m + n
}

func sub(m, n int) int {
	return m - n
}

func main() {
	var soper []string
	var mapfunc map[string]func(int, int) int

	soper = []string{"+", "-", "*", "/"}
	mapfunc = map[string]func(int, int) int{
		"+": add,
		"-": sub,
	}

	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	oper := os.Args[2]

	f := mapfunc[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
