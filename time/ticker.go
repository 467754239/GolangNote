package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 5)
	for range timer.C {
		fmt.Println("ticker")
	}
}
