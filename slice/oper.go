package main

import (
	"fmt"
	"strconv"
)

func operationalHandler(x int, oper string, y int) (sum int) {
	switch oper {
	case "/":
		sum = x / y
	case "+":
		sum = x + y
	case "-":
		sum = x - y
	case "*":
		sum = x * y
	}
	return

}

func calc(args ...string) (sum int) {
	sum, _ = strconv.Atoi(args[0])

	for i := 1; i < len(args); i++ {
		if i+1 < len(args) {
			if ii, err := strconv.Atoi(args[i+1]); err == nil {
				sum = operationalHandler(sum, args[i], ii)
			}
		}
	}
	return
}

func main() {
	fmt.Println(calc("4", "+", "2", "/", "3", "+", "1", "*", "10"))
}
