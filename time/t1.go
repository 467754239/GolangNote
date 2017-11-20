package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())

	// type Duration int64
	var n time.Duration
	n = 2 * time.Hour
	fmt.Println(int64(n))
	fmt.Println(n.String())

	n = 3*time.Hour + 30*time.Minute
	fmt.Println(n.String())

	fmt.Println("----------add----------------")
	t := time.Now()

	// 1小时之前的时间
	t1 := t.Add(-(1*time.Hour + 30*time.Minute))
	fmt.Printf("t1:%v\n", t1)

	// 求两个时间点的差
	fmt.Println(t1.Sub(t))
}
