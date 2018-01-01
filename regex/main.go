package main

import (
	"fmt"
	"regexp"
)

func main() {
	var reg = regexp.MustCompile("[a-z]+")

	ok := reg.MatchString("abc12345")
	fmt.Println(ok)

	fmt.Println(reg.FindString("abc12345"))
}
